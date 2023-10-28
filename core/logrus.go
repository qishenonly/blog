package core

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/qishenonly/blog/global"

	"github.com/sirupsen/logrus"
)

// 颜色常量
const (
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

type LogFormatter struct{}

// Format 实现了 logrus.Formatter 接口
func (lf *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 设置颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	// 设置输出格式
	var buf *bytes.Buffer
	if entry.Buffer != nil {
		buf = entry.Buffer
	} else {
		buf = &bytes.Buffer{}
	}

	log := global.Config.Logger

	// 自定义日期格式
	customTime := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcValue := entry.Caller.Function
		fileValue := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		// 自定义输出格式
		fmt.Fprintf(buf, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s <%s>\n",
			log.Prefix, customTime, levelColor, entry.Level.String(), funcValue, fileValue, entry.Message)
	} else {
		fmt.Fprintf(buf, "%s[%s] \x1b[%dm[%s]\x1b[0m <%s>\n",
			log.Prefix, customTime, levelColor, entry.Level.String(), entry.Message)
	}

	return buf.Bytes(), nil
}

// InitLogger 初始化日志
func InitLogger() *logrus.Logger {
	lg := logrus.New()
	// 设置日志输出到标准输出
	lg.SetOutput(os.Stdout)
	// 设置是否输出文件名和行号
	lg.SetReportCaller(global.Config.Logger.ShowLine)
	// 设置日志格式
	lg.SetFormatter(&LogFormatter{})
	// 设置日志级别
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		lg.SetLevel(logrus.DebugLevel)
	} else {
		lg.SetLevel(level)
	}
	//InitDefaultLogger()
	return lg
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(level)
	}
}
