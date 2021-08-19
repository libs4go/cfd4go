package cfd4go

import "github.com/libs4go/errors"

// ScopeOfAPIError .
const errVendor = "cfd4go"

// errors
var (
	ErrExpect = errors.New("Precondition check error", errors.WithVendor(errVendor), errors.WithCode(-1))
	ErrEnsure = errors.New("Aftercondition check error", errors.WithVendor(errVendor), errors.WithCode(-2))
	ErrAssert = errors.New("Condition check error", errors.WithVendor(errVendor), errors.WithCode(-3))
)

type Asserter func(bool, string)

type ConditionChecker func(assert Asserter)

func expectAssert(condition bool, msg string) {
	if !condition {
		panic(errors.Wrap(ErrExpect, msg))
	}
}

func ensureAssert(condition bool, msg string) {
	if !condition {
		panic(errors.Wrap(ErrEnsure, msg))
	}
}

// Expect
func Expect(checker ConditionChecker) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	checker(expectAssert)

	return
}

func Assert(condition bool, msg string) {
	if !condition {
		panic(errors.Wrap(ErrAssert, msg))
	}
}

func Ensure(checker ConditionChecker) {
	checker(ensureAssert)
}
