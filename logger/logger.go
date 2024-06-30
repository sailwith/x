package logger

import "go.uber.org/zap"

type Logger struct {
	instance *zap.SugaredLogger
}

func New() (*Logger, error) {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &Logger{
		instance: logger.Sugar(),
	}, nil
}

func (l *Logger) Debug(args ...any) {
	l.instance.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...any) {
	l.instance.Debugf(template, args...)
}

func (l *Logger) Debugln(args ...any) {
	l.instance.Debugln(args...)
}

func (l *Logger) Info(args ...any) {
	l.instance.Info(args...)
}

func (l *Logger) Infof(template string, args ...any) {
	l.instance.Infof(template, args...)
}

func (l *Logger) Infoln(args ...any) {
	l.instance.Infoln(args...)
}

func (l *Logger) Warn(args ...any) {
	l.instance.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...any) {
	l.instance.Warnf(template, args...)
}

func (l *Logger) Warnln(args ...any) {
	l.instance.Warnln(args...)
}

func (l *Logger) Error(args ...any) {
	l.instance.Error(args...)
}

func (l *Logger) Errorf(template string, args ...any) {
	l.instance.Errorf(template, args...)
}

func (l *Logger) Errorln(args ...any) {
	l.instance.Errorln(args...)
}

func (l *Logger) Fatal(args ...any) {
	l.instance.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...any) {
	l.instance.Fatalf(template, args...)
}

func (l *Logger) Fatalln(args ...any) {
	l.instance.Fatalln(args...)
}

func (l *Logger) Panic(args ...any) {
	l.instance.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...any) {
	l.instance.Panicf(template, args...)
}

func (l *Logger) Panicln(args ...any) {
	l.instance.Panicln(args...)
}
