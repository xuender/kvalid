package kvalid_test

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
func (p *Book) Validation(method string) *kvalid.Rules {
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

// nolint: lll
func ExampleRules_Validate() {
	book := &Book{}
	fmt.Println(book.Validate(http.MethodPost))

	book.Title = "Hello World"
	fmt.Println(book.Validate(http.MethodPost))

	book.Author = "ender"
	fmt.Println(book.Validate(http.MethodPost))
	fmt.Println(book.Validate(http.MethodPut))

	data, _ := json.Marshal(book.Validation(http.MethodPut))
	fmt.Println(string(data))

	// Output:
	// 	title required. author required.
	// author required.
	// <nil>
	// amount required.
	// {"Amount":[{"rule":"required","msg":"amount required"},{"rule":"minNum","min":10.3,"msg":"amount min 10.3"},{"rule":"maxNum","max":2000,"msg":"amount max 2000"}]}
}
