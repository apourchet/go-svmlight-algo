package svmalgo

import (
	. "github.com/apourchet/go-svmlight"
)

type DTAttribute int

type DTInstance SVMInstance

type CompFunction func(instance DTInstance, attribute DTAttribute) bool

type InfoGainFunc func(instances []DTInstance, attributes []DTAttribute) DTAttribute
