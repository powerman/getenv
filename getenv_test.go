package getenv_test

import (
	"os"
	"testing"
	"time"

	"github.com/powerman/check"
	"github.com/powerman/getenv"
)

func Test(t *testing.T) {
	os.Unsetenv("UNSET")
	os.Setenv("EMPTY", "")
	os.Setenv("BOOL", "false")
	os.Setenv("DUR", "1m")
	os.Setenv("FLOAT", "1.23")
	os.Setenv("INT", "42")
	os.Setenv("STR", "text")

	t.Run("Bool", func(tt *testing.T) {
		t := check.T(tt)
		t.Equal(getenv.Bool("UNSET", true), true)
		t.Equal(getenv.Bool("EMPTY", true), true)
		t.Equal(getenv.Bool("BOOL", true), false)
		t.Nil(getenv.LastErr())
		t.Equal(getenv.Bool("STR", true), true)
		t.Match(getenv.LastErr(), "parse")
	})
	t.Run("Dur", func(tt *testing.T) {
		t := check.T(tt)
		t.Equal(getenv.Dur("UNSET", 3*time.Second), 3*time.Second)
		t.Equal(getenv.Dur("EMPTY", 3*time.Second), 3*time.Second)
		t.Equal(getenv.Dur("DUR", 3*time.Second), 60*time.Second)
		t.Nil(getenv.LastErr())
		t.Equal(getenv.Dur("STR", 3*time.Second), 3*time.Second)
		t.Match(getenv.LastErr(), "parse")
	})
	t.Run("Float", func(tt *testing.T) {
		t := check.T(tt)
		t.Equal(getenv.Float("UNSET", 0.5), 0.5)
		t.Equal(getenv.Float("EMPTY", 0.5), 0.5)
		t.Equal(getenv.Float("FLOAT", 0.5), 1.23)
		t.Nil(getenv.LastErr())
		t.Equal(getenv.Float("STR", 0.5), 0.5)
		t.Match(getenv.LastErr(), "parse")
	})
	t.Run("Int", func(tt *testing.T) {
		t := check.T(tt)
		t.Equal(getenv.Int("UNSET", 5), 5)
		t.Equal(getenv.Int("EMPTY", 5), 5)
		t.Equal(getenv.Int("INT", 5), 42)
		t.Nil(getenv.LastErr())
		t.Equal(getenv.Int("STR", 5), 5)
		t.Match(getenv.LastErr(), "parse")
	})
	t.Run("Str", func(tt *testing.T) {
		t := check.T(tt)
		t.Equal(getenv.Str("UNSET", "def"), "def")
		t.Equal(getenv.Str("EMPTY", "def"), "def")
		t.Equal(getenv.Str("STR", "def"), "text")
		t.Equal(getenv.Str("INT", "def"), "42")
		t.Nil(getenv.LastErr())
	})
}
