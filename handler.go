package job

type Handler interface {
	Confirm(func(Job))
	Process(func(Job))
	Submit(Job)
}