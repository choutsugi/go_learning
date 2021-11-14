package tailfile

import (
	"LogAgent/logger"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func TailLines() chan *tail.Line {
	return tailObj.Lines
}

func TailFilename() string {
	return tailObj.Filename
}

func Init(filename string) (err error) {
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

	tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		logger.Z.Error("tailfile: create tailObj for path %s failed", filename)
		return err
	}

	return nil
}
