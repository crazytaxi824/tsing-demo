package global

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// 设置logger
func SetLogger() error {
	// 设置级别
	level := strings.ToLower(LocalConfig.Logger.Level)
	switch level {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "empty":
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.Disabled)
		return nil
	}

	// 设置时间格式
	if LocalConfig.Logger.TimeFormat == "timestamp" {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	} else {
		zerolog.TimeFieldFormat = timeFormater(LocalConfig.Logger.TimeFormat)
	}

	// 设置日志输出方式
	var output io.Writer
	var logFile *os.File
	var err error
	// 设置日志文件
	if LocalConfig.Logger.FilePath != "" {
		// 输出到文件
		if LocalConfig.Logger.FileMode == 0 {
			LocalConfig.Logger.FileMode = os.FileMode(0600)
		}
		logFile, err = os.OpenFile(LocalConfig.Logger.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, LocalConfig.Logger.FileMode)
		if nil != err {
			return err
		}
	}
	switch LocalConfig.Logger.Encode {
	// console编码
	case "console":
		if logFile != nil {
			output = zerolog.ConsoleWriter{
				Out:        logFile,
				NoColor:    true,
				TimeFormat: zerolog.TimeFieldFormat,
			}
		} else {
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: zerolog.TimeFieldFormat,
			}
		}
	// json编码
	case "json":
		if logFile != nil {
			output = logFile
		} else {
			output = os.Stdout
		}
	default:
		return errors.New("从配置文件的logger.encode中获得了未知的参数，目前只支持json|console")
	}

	log.Logger = log.Output(output)

	return nil
}

func timeFormater(str string) string {
	str = strings.Replace(str, "y", "2006", -1)
	str = strings.Replace(str, "m", "01", -1)
	str = strings.Replace(str, "d", "02", -1)
	str = strings.Replace(str, "h", "15", -1)
	str = strings.Replace(str, "i", "04", -1)
	str = strings.Replace(str, "s", "05", -1)
	return str
}
