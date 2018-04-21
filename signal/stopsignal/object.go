package stopsignal

import (
	"github.com/the-anna-project/the-anna-project/signal"
)

type Object struct {
}

func (o *Object) AppendData(v interface{}) {
}

func (o *Object) Copy() signal.Interface {
	return nil
}

func (o *Object) Data() interface{} {
	return nil
}
