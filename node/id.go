package node

func (o *Object) ID() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.id
}
