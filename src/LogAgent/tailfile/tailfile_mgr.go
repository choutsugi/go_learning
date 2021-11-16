package tailfile

import (
	"LogAgent/common"
	"LogAgent/logger"
)

type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask
	collectEntryList []common.CollectEntry
	confChan         chan []common.CollectEntry
}

var (
	mgr *tailTaskMgr
)

func Init(allConf []common.CollectEntry) (err error) {

	mgr = &tailTaskMgr{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: allConf,
		confChan:         make(chan []common.CollectEntry),
	}

	for _, conf := range allConf {
		// 创建任务
		task := newTailTask(conf)
		err = task.init()
		if err != nil {
			logger.Z.Errorf("tailfile: create task:%s failed, err:%v", conf.Path, err)
			continue
		}
		// 记录已启动的任务
		mgr.tailTaskMap[task.path] = task
		logger.Z.Infof("tailfile: task:%s is ready to start.", task.path)
		// 执行任务
		go task.run()
	}

	go mgr.watch()

	return
}

func (t *tailTaskMgr) watch() {
	for {
		// 阻塞等待配置更新
		newConf := <-mgr.confChan
		logger.Z.Infof("tailfile: configuration has been updated from etcd: %v.", newConf)
		// 处理新配置
		for _, conf := range newConf {
			// 若任务已存在：查看下一个配置
			if t.isExist(conf) {
				continue
			}
			// 若任务不存在：新建任务
			task := newTailTask(conf)
			err := task.init()
			if err != nil {
				logger.Z.Errorf("tailfile: create task:%s failed, err:%v", conf.Path, err)
				continue
			}
			mgr.tailTaskMap[task.path] = task
			logger.Z.Infof("tailfile: task:%s is ready to start.", task.path)
			go task.run()
		}

		// 检查已有任务是否被取消
		for key, task := range t.tailTaskMap {
			var found bool
			for _, conf := range newConf {
				if key == conf.Path {
					found = true
					break
				}
			}
			if !found {
				logger.Z.Infof("tailfile: task:%s is ready to stop.", key)
				task.cancel()
				delete(t.tailTaskMap, task.path)
			}
		}
	}
}

func (t *tailTaskMgr) isExist(conf common.CollectEntry) (ok bool) {
	_, ok = t.tailTaskMap[conf.Path]
	return
}

// UpdateConf 供etcd模块更新配置
func UpdateConf(newConf []common.CollectEntry) {
	mgr.confChan <- newConf
}
