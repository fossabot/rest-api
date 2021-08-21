package logformatter

import (
	"bytes"
	"io"

	"github.com/sirupsen/logrus"
)

var (
	_ logrus.Formatter = &Formatter{}
)

type Formatter struct {
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := f.writeLevel(buffer, entry); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (f *Formatter) writeLevel(buffer io.Writer, entry *logrus.Entry) error {
	return nil
}
