// Package util provides common static behavior.
package util

// Must panics on errors.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
