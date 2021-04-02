package goargs

import "strconv"

// Arg is the structure for a single argument.
type Arg string

// String returns a string type argument.
func (a *Arg) String() string {
	return string(*a)
}

// Int converts and returns an int64 type argument.
func (a *Arg) Int(defaultValue int64) int64 {
	value, err := strconv.ParseInt(a.String(), 0, 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// Float converts and returns an int64 type argument.
func (a *Arg) Float(defaultValue float64) float64 {
	value, err := strconv.ParseFloat(a.String(), 64)
	if err != nil {
		return defaultValue
	}
	return value
}

// Bool converts and returns an bool type argument.
func (a *Arg) Bool(defaultValue bool) bool {
	value, err := strconv.ParseBool(a.String())
	if err != nil {
		return defaultValue
	}
	return value
}