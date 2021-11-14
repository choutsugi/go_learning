package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Z = logrus.New()

func Init() {
	Z.Out = os.Stdout
	Z.SetLevel(logrus.DebugLevel)
}
