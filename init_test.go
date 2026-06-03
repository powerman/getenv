package getenv_test

import (
	"os"
	"testing"

	"github.com/powerman/check"
)

func TestMain(m *testing.M) {
	// These env vars are used by all subtests in getenv_test.go.
	// Set them here (not with t.Setenv in each subtest) because
	// the top-level Test uses t.Parallel(), and t.Setenv would panic
	// if called after t.Parallel().
	os.Unsetenv("UNSET")
	os.Setenv("EMPTY", "")
	os.Setenv("BOOL", "false")
	os.Setenv("DUR", "1m")
	os.Setenv("FLOAT", "1.23")
	os.Setenv("INT", "42")
	os.Setenv("STR", "text")
	check.TestMain(m)
}
