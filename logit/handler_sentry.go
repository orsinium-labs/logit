package logit

import (
	"fmt"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

type SentryHandler struct {
	BaseHandler
	DSN     string
	Timeout string
}

func NewSentryHandler() SentryHandler {
	return SentryHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config SentryHandler) Parse() (*Handler, error) {
	allLevels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
	hook, err := logrus_sentry.NewSentryHook(config.DSN, allLevels)
	if err != nil {
		return nil, err
	}

	hook.Timeout, err = time.ParseDuration(config.Timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot parse timeout: %v", err)
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}
