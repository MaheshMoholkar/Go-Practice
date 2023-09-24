package util

import "strings"

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + "_" + s
	}
}

func TransformString(s string, fn TransformFunc) string {
	return fn(s)
}
