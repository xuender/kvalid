package kvalid_test

import (
	"fmt"

	"github.com/xuender/kvalid"
)

func ExampleIsEmail() {
	fmt.Println(kvalid.IsEmail("test@example.com"))
	fmt.Println(kvalid.IsEmail("test"))

	// Output:
	// true
	// false
}
