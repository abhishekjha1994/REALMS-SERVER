package errutil

import (
	r "github.com/revel/revel"
)

// CheckErr checks and handles errors
func CheckErr(err error, msg string) {
	if err != nil {
		r.ERROR.Println(err, msg)
	}
}
