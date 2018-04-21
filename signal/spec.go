package signal

type Interface interface {
	AppendData(interface{})
	Copy() Interface
	Data() interface{}
}
