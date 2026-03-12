package zap

type Logger struct{}

func L() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, fields ...interface{})  {}
func (l *Logger) Error(msg string, fields ...interface{}) {}
func (l *Logger) Warn(msg string, fields ...interface{})  {}
func (l *Logger) Debug(msg string, fields ...interface{}) {}