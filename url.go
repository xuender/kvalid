package kvalid

import (
	"encoding/json"
	"regexp"
)

const (
	_schema = `((ftp|tcp|udp|wss?|https?):\/\/)`
	_user   = `(\S+(:\S*)?@)?`
	_domain = `(\w+([-_.]?\w)*\w\.\w+)`
	_port   = `(:(\d{1,5}))?`
	_path   = `((\/|\?|#)\S*)*`
	_url    = `^` + _schema + _user + _domain + _port + _path + `$`
)

var _urlRegex = regexp.MustCompile(_url)

// URLValidator field must be a valid URL.
type URLValidator struct {
	PatternValidator
}

// URL field must be a valid URL.
func URL() *URLValidator {
	return &URLValidator{
		PatternValidator{
			re:      _urlRegex,
			message: "Please use a valid URL",
		},
	}
}

// MarshalJSON for this validator.
func (p *URLValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStruct[int]{
		Rule: "url",
		Msg:  p.message,
	})
}

// IsURL returns true if the string is an URL.
func IsURL(url string) bool {
	return _urlRegex.MatchString(url)
}
