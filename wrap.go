package log

func (p *Wrap) Info(args ...interface{}) {
	p.logger.Info(args...)
}

func (p *Wrap) Warn(args ...interface{}) {
	p.logger.Warn(args...)
}

func (p *Wrap) Error(args ...interface{}) {
	p.logger.Error(args...)
}

func (p *Wrap) Panic(args ...interface{}) {
	p.logger.Error(args...)
}

func (p *Wrap) Infof(template string, args ...interface{}) {
	p.logger.Infof(template, args...)
}

func (p *Wrap) Warnf(template string, args ...interface{}) {
	p.logger.Warnf(template, args...)
}

func (p *Wrap) Errorf(template string, args ...interface{}) {
	p.logger.Errorf(template, args...)
}

func (p *Wrap) Panicf(template string, args ...interface{}) {
	p.logger.Panicf(template, args...)
}

func (p *Wrap) With(args ...interface{}) *Wrap {
	return &Wrap{logger: p.logger.With(args...)}
}
