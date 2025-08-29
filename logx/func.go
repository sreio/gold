package logx

func Info(args ...any) {
	L.Info(args...)
}
func InfoF(format string, args ...any) {
	L.Infof(format, args...)
}
func Warn(args ...any) {
	L.Warn(args...)
}
func WarnF(format string, args ...any) {
	L.Warnf(format, args...)
}
func Error(args ...any) {
	L.Error(args...)
}
func ErrorF(format string, args ...any) {
	L.Errorf(format, args...)
}
func Fatal(args ...any) {
	L.Fatal(args...)
}
func FatalF(format string, args ...any) {
	L.Fatalf(format, args...)
}
func Debug(args ...any) {
	L.Debug(args...)
}
func DebugF(format string, args ...any) {
	L.Debugf(format, args...)
}
func Trace(args ...any) {
	L.Trace(args...)
}
func TraceF(format string, args ...any) {
	L.Tracef(format, args...)
}
func Panic(args ...any) {
	L.Panic(args...)
}
func PanicF(format string, args ...any) {
	L.Panicf(format, args...)
}
