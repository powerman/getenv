package getenv

import (
	"os"
	"strings"
)

// Scrub removes every environment variable with the given prefix from the process environment.
// It panics if prefix is empty — passing an empty prefix is always a bug:
// it would either silently skip scrubbing (leaking secrets in the process environment)
// or delete every environment variable, which is better expressed as an explicit
// [os.Clearenv] call.
//
// Typical use: after all configuration has been resolved from environment variables,
// call Scrub to prevent any subsequent code from reading configuration (especially credentials)
// from the environment.
func Scrub(prefix string) {
	if prefix == "" {
		panic("Scrub: empty prefix")
	}
	for _, kv := range os.Environ() {
		if name, _, ok := strings.Cut(kv, "="); ok && strings.HasPrefix(name, prefix) {
			_ = os.Unsetenv(name)
		}
	}
}
