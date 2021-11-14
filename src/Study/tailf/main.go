package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	filename := `./xx.log`
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}

	//打开文件读取数据
	file, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tail %s failed, err:%v\n\n", filename, err)
		return
	}

	var (
		msg *tail.Line
		ok  bool
	)

	for {
		msg, ok = <-file.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", file.Filename)
			break
		}
		fmt.Printf("msg:%s", msg.Text)
	}
}
