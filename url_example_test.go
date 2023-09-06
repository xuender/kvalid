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

	// Output:
	// false
	// true
	// true
}
