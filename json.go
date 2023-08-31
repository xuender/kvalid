package kvalid

type jsonStruct[N Number] struct {
	Rule    string `json:"rule"`
	Min     N      `json:"min,omitempty"`
	Max     N      `json:"max,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Msg     string `json:"msg,omitempty"`
}
