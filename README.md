# kvalid

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

kvalid is a lightweight validation library that can export rules as JSON so browsers can apply the same rules.

Support Go 1.18 and later.

Learn from [github.com/AgentCosmic/xvalid](https://github.com/AgentCosmic/xvalid) .

## Use

Define rules:

```go
import (
  "fmt"
  "net/http"

  "github.com/xuender/kvalid"
  "github.com/xuender/kvalid/json"
)

type Book struct {
  Title  string `json:"title"`
  Author string `json:"author,omitempty"`
  Amount float64
}

// nolint: gomnd
func (p *Book) Validation(method string) *kvalid.Rules[*Book] {
  switch method {
  case http.MethodPut:
    return kvalid.New(p).
      Field(&p.Amount,
        kvalid.Required().SetMessage("amount required"),
        kvalid.MinNum(10.3).Optional().SetMessage("amount min 10.3"),
        kvalid.MaxNum(2000.0).SetMessage("amount max 2000"),
      )
  case http.MethodPost:
    return kvalid.New(p).
      Field(&p.Title,
        kvalid.Required().SetMessage("title required"),
        kvalid.MaxStr(200).SetMessage("title max 200"),
      ).
      Field(&p.Author,
        kvalid.Required().SetMessage("author required"),
        kvalid.MaxStr(100).SetMessage("author max 100"),
      )
  default:
    panic("illegal method:" + method)
  }
}

func (p *Book) Validate(method string) error {
  return p.Validation(method).Validate(p)
}
```

Validate objects:

```go
book := &Book{}
fmt.Println(book.Validate(http.MethodPost))
```

Export rules as JSON:

```go
data, _ := json.Marshal(book.Validation(http.MethodPut))
fmt.Println(string(data))
```

Bind object:

```go
source := &Book{Title: "Hello World", Author: "ender"}
target := &Book{}
rules := source.Validation(http.MethodPost)

rules.Bind(source, target)
```

## Validators

* Email
* MaxNum, MinNum
  * int, int8, int16, int32, int64
  * uint, uint8, uint16, uint32, uint64
  * float32, float64
  * byte, rune
* MaxStr, MinStr
* MaxNullInt, MinNullInt
* Pattern
* Required
* FieldFunc, StructFunc

## License

© ender, 2023~time.Now

[MIT LICENSE](https://github.com/xuender/kvalid/blob/master/LICENSE)

[action-url]: https://github.com/xuender/kvalid/actions
[action-svg]: https://github.com/xuender/kvalid/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/github.com/xuender/kvalid
[goreport-svg]: https://goreportcard.com/badge/github.com/xuender/kvalid

[godoc-url]: https://godoc.org/github.com/xuender/kvalid
[godoc-svg]: https://godoc.org/github.com/xuender/kvalid?status.svg

[license-url]: https://github.com/xuender/kvalid/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
