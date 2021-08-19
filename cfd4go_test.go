package cfd4go

import (
	"testing"
)

func TestPanic(t *testing.T) {
	err := Expect(func(assert Asserter) {
		assert(false, "always success !!!!")
	})

	defer Ensure(func(assert Asserter) {
		assert(err == nil, "ensure !!!!")
	})

}
