package netopia

import "log"

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type DefaultLogger struct{}

func (d *DefaultLogger) Debugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}

func (d *DefaultLogger) Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func (d *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}
