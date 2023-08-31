package kvalid

type jsonStruct struct {
	Rule string `json:"rule"`
	Min  int64  `json:"min,omitempty"`
	Msg  string `json:"msg,omitempty"`
}
