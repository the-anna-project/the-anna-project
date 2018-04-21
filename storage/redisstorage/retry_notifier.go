package redisstorage

import (
	"fmt"
	"time"
)

func (o *Object) retryNotifier(err error, d time.Duration) {
	o.logger.Log("level", "warning", "message", "storage lookup failed", "stack", fmt.Sprintf("%#v", err))
}
