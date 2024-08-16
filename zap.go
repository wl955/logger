package log

func Info(args ...interface{}) {
	logger.With("service", opt.service).Info(args...)
}

func Warn(args ...interface{}) {
	logger.With("service", opt.service).Warn(args...)
}

func Error(args ...interface{}) {
	logger.With("service", opt.service).Error(args...)
}

func Panic(args ...interface{}) {
	logger.With("service", opt.service).Panic(args...)
}

func Infof(template string, args ...interface{}) {
	logger.With("service", opt.service).Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.With("service", opt.service).Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.With("service", opt.service).Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logger.With("service", opt.service).Panicf(template, args...)
}

func With(args ...interface{}) *Wrap {
	return &Wrap{logger: logger.With("service", opt.service).With(args...)}
}
