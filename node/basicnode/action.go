package basicnode

import "github.com/the-anna-project/the-anna-project/action"

func (o *Object) Action() action.Interface {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.action
}
