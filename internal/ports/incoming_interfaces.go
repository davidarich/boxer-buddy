package ports

type Interop interface {
	StartIO()
}

type InteropWriter interface {
	Write(message []byte)
}

type MessageRouter interface {
	Route(msg []byte, response chan []byte) (err error)
}

type ViewManager interface {
	OpenMain(title string)
}

type WindowManager interface {
	FocusWindowByPid(pid int)
}
