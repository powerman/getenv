package getenv_test

import (
	"os"
	"testing"

	"github.com/powerman/check"

	"github.com/powerman/getenv"
)

func TestScrub(tt *testing.T) {
	const prefix = "GETENV_SCRUB_TEST_"
	t := check.T(tt)

	tt.Setenv(prefix+"KEEP", "")         // Empty value — still removed.
	tt.Setenv(prefix+"SECRET", "s3cret") // Typical credential.
	tt.Setenv("GETENV_SCRUB_OTHER", "survive")
	tt.Setenv("SOME_VAR", "also-survive")

	getenv.Scrub(prefix)

	_, ok := os.LookupEnv(prefix + "KEEP")
	t.False(ok) // Empty var under prefix is removed.
	_, ok = os.LookupEnv(prefix + "SECRET")
	t.False(ok) // Prefixed vars are removed.

	v, ok := os.LookupEnv("GETENV_SCRUB_OTHER")
	t.True(ok)
	t.Equal(v, "survive")

	v, ok = os.LookupEnv("SOME_VAR")
	t.True(ok)
	t.Equal(v, "also-survive")
}

func TestScrubEmptyPrefix(tt *testing.T) {
	t := check.T(tt)
	t.PanicMatch(func() { getenv.Scrub("") }, "empty prefix")
}

func TestScrubNoMatch(tt *testing.T) {
	const prefix = "GETENV_SCRUB_NOMATCH_"
	t := check.T(tt)

	tt.Setenv("GETENV_UNRELATED", "keep")

	getenv.Scrub(prefix)

	v, ok := os.LookupEnv("GETENV_UNRELATED")
	t.True(ok)
	t.Equal(v, "keep")
}
