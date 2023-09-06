// nolint: dupword
package kvalid_test

import (
	"fmt"

	"github.com/xuender/kvalid"
)

func ExampleIsURL() {
	fmt.Println(kvalid.IsURL("http//:example.com"))
	fmt.Println(kvalid.IsURL("http://www.example.com"))
	fmt.Println(kvalid.IsURL("ftp://aa:pp@www.example.com"))
	fmt.Println(kvalid.IsURL("http://localhost:4200"))
	fmt.Println(kvalid.IsURL("http://127.0.0.1:4200"))

	// Output:
	// false
	// true
	// true
	// true
	// true
}
