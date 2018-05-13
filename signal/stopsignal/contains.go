package stopsignal

import (
	"github.com/the-anna-project/the-anna-project/signal"
)

func Contains(sigs []signal.Interface) bool {
	for _, s := range sigs {
		_, ok := s.(*Object)
		if ok {
			return true
		}
	}

	return false
}
