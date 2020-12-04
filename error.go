package camp

type ErrMsg struct {
	Prefix string
	Msg    string
}

//通过ErrInfo构建
func NewError(msg string, track ...string) error {
	var t string
	if len(track) > 0 {
		t = track[0]
	}
	return &ErrMsg{
		Prefix: t,
		Msg:    msg,
	}
}

//通过Prefix构建
func PrefixError(trick string, err error) error {
	if err == nil {
		return nil
	}
	return &ErrMsg{
		Prefix: trick,
		Msg:    err.Error(),
	}
}

func (e *ErrMsg) Error() string {
	return e.error()
}

func (e *ErrMsg) error() string {
	var msg string
	if e.Prefix != "" {
		msg = "[TrackKey:" + e.Prefix + "]"
	}
	return msg + e.Msg
}

func NewErrorTrack(track string) ErrMsg {
	return ErrMsg{Prefix: track}
}

func (e ErrMsg) AddMsg(msg string) error {
	return &ErrMsg{Prefix: e.Prefix, Msg: msg}
}
