package server

import (
	"fmt"

	"github.com/thockin-test/util/foobar"
)

func Server() {
	fmt.Printf("This is %s\n", foobar.FooBar())
}
