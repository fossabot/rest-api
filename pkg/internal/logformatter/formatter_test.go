package logformatter

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter_Format(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		log := logrus.New()
		log.Formatter = &Formatter{}

		log.WithField("test", t.Name()).Info("test message")
	})
}
