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
	fmt.Println(kvalid.IsURL("http://www.example.com:8080"))

	// Output:
	// false
	// true
	// true
	// true
}
