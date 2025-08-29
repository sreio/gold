package logx

import (
	"github.com/natefinch/lumberjack"
	"github.com/sreio/gold/config"
	"io"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// L 全局 logger
// 对外暴露一组便捷函数（也可以直接用 logger.L 取 *logrus.Logger）
//
//	logx.L.Info("Info")
//	logx.L.Warn("Warn")
//	logx.L.Error("Error")
//	logx.L.Fatal("Fatal")
//	logx.L.Debug("Debug")
//	logx.L.Trace("Trace")
//	logx.L.Panic("Panic")
//
// ....
var L = logrus.New()

// Init 初始化全局 logger
func Init(opt config.Logx) *logrus.Logger {
	// Level
	level := parseLevel(opt.Level)
	L.SetLevel(level)

	// Formatter
	timeFormat := opt.TimeFormat
	if timeFormat == "" {
		timeFormat = time.DateTime // 例如 2025-08-29 10:30:45
	}
	if opt.JSON {
		L.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	} else {
		L.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeFormat,
			PadLevelText:    true,
		})
	}

	L.SetReportCaller(opt.WithCaller)

	// Output 组合
	var writers []io.Writer

	if opt.Console {
		// 控制台：用 Lerr 更符合日志语义
		writers = append(writers, os.Stderr)
	}

	if opt.File.Enable {
		// 确保目录存在
		_ = os.MkdirAll(dirOf(opt.File.Path), 0o755)

		lj := &lumberjack.Logger{
			Filename:   opt.File.Path,
			MaxSize:    nonZero(opt.File.MaxSize, 100),  // 默认 100MB
			MaxBackups: nonZero(opt.File.MaxBackups, 7), // 默认保留 7 个
			MaxAge:     nonZero(opt.File.MaxAge, 7),     // 默认 7 天
			Compress:   opt.File.Compress,
		}

		// 如果需要“仅文件特定级别阈值”
		if lvlStr := strings.TrimSpace(opt.File.MinLevelForFile); lvlStr != "" && parseLevel(lvlStr) != level {
			// 用 Hook 把达到阈值的记录写入文件
			L.AddHook(newLevelSplitHook(parseLevel(lvlStr), lj))
		} else {
			// 否则文件与控制台共享同一 writer
			writers = append(writers, lj)
		}
	}

	if len(writers) == 0 {
		// 默认至少输出到控制台
		writers = append(writers, os.Stderr)
	}

	L.SetOutput(io.MultiWriter(writers...))
	return L
}

// --------- 工具 & Hook ----------

func parseLevel(s string) logrus.Level {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info", "":
		return logrus.InfoLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func dirOf(path string) string {
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "."
	}
	return path[:i]
}

func nonZero(v, def int) int {
	if v == 0 {
		return def
	}
	return v
}

// levelSplitHook: 把达到 minLevel 的日志单独写入 writer（常用于“文件仅收敛更高等级”）
type levelSplitHook struct {
	minLevel logrus.Level
	w        io.Writer
}

func newLevelSplitHook(min logrus.Level, w io.Writer) *levelSplitHook {
	return &levelSplitHook{minLevel: min, w: w}
}

func (h *levelSplitHook) Levels() []logrus.Level {
	// 所有级别都接收，Fire 决定是否写入
	return logrus.AllLevels
}

func (h *levelSplitHook) Fire(e *logrus.Entry) error {
	if e.Level <= h.minLevel { // 注意：级别值越小越详细，Trace < Debug < Info < Warn < Error...
		// 使用当前 formatter 格式化后写入指定 writer
		b, err := e.Logger.Formatter.Format(e)
		if err != nil {
			return err
		}
		_, _ = h.w.Write(b)
	}
	return nil
}
