// Package z exposes a curated set of utility functions.
//
// Its primary goals are two-fold:
//
//  1. Make the best commonly used utilities discoverable
//  2. Reduce hand strain; "I can finally stop writing `func must(...)` everywhere!"
package z

import (
	"errors"
)

// Must panics IFF any error in errs is non-nil.
func Must(errs ...error) {
	if err := errors.Join(errs...); err != nil {
		panic(err)
	}
}

// MustBe panics IFF err is non-nil, otherwise it returns t.
func MustBe[T any](t T, err error) T {
	Must(err)
	return t
}

// P returns a pointer to t.
func P[T any](t T) *T {
	return &t
}
