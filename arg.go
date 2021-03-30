package goargs

import "strconv"

type Arg string

func (a *Arg) String() string {
	return string(*a)
}

func (a *Arg) Int(defaultValue int64) int64 {
	value, err := strconv.ParseInt(a.String(), 0, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func (a *Arg) Float(defaultValue float64) float64 {
	value, err := strconv.ParseFloat(a.String(), 64)
	if err != nil {
		return defaultValue
	}
	return value
}

func (a *Arg) Bool(defaultValue bool) bool {
	value, err := strconv.ParseBool(a.String())
	if err != nil {
		return defaultValue
	}
	return value
}