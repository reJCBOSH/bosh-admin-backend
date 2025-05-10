package exception

type Exception struct {
	msg string
}

func (e *Exception) Error() string {
	return e.msg
}

func NewException(msg string) *Exception {
	return &Exception{
		msg: msg,
	}
}
