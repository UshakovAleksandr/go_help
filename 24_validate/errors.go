package main

import "errors"

var (
	errFieldRequire = errors.New("field is require")
	errNotStruct    = errors.New("input message is not struct")
	errInvalidLen   = errors.New("invalid len")
)
