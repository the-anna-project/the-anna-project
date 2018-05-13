package basicport

func (o *Object) NodeID() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.nodeID
}
