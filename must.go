package getenv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// MustBool returns value of environment variable name as a bool.
// It panics if environment variable is not set or failed to parse as bool.
func MustBool(name string) bool {
	value, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("required env %s is not set", name))
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Sprintf("env %s=%q is not a valid bool", name, value))
	}
	return b
}

// MustDur returns value of environment variable name as a [time.Duration].
// It panics if environment variable is not set or failed to parse as a [time.Duration].
func MustDur(name string) time.Duration {
	value, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("required env %s is not set", name))
	}
	dur, err := time.ParseDuration(value)
	if err != nil {
		panic(fmt.Sprintf("env %s=%q is not a valid duration", name, value))
	}
	return dur
}

// MustFloat returns value of environment variable name as a float64.
// It panics if environment variable is not set or failed to parse as a float64.
func MustFloat(name string) float64 {
	value, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("required env %s is not set", name))
	}
	v, err := strconv.ParseFloat(value, parseBits)
	if err != nil {
		panic(fmt.Sprintf("env %s=%q is not a valid float", name, value))
	}
	return v
}

// MustInt returns value of environment variable name as an int.
// It panics if environment variable is not set or failed to parse as an int.
func MustInt(name string) int {
	value, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("required env %s is not set", name))
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("env %s=%q is not a valid integer", name, value))
	}
	return i
}

// MustStr returns value of environment variable name as a string.
// It panics if environment variable is not set.
func MustStr(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("required env %s is not set", name))
	}
	return value
}
