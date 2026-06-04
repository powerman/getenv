package getenv_test

import (
	"testing"
	"time"

	"github.com/powerman/check"

	"github.com/powerman/getenv"
)

func Test(t *testing.T) {
	t.Setenv("UNSET", "")
	t.Setenv("EMPTY", "")
	t.Setenv("BOOL", "false")
	t.Setenv("DUR", "1m")
	t.Setenv("FLOAT", "1.23")
	t.Setenv("INT", "42")
	t.Setenv("STR", "text")

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
	t.Run("LastErrs", func(tt *testing.T) {
		t := check.T(tt)
		// No error initially.
		t.Nil(getenv.LastErrs())
		// After one error LastErrs returns it and clears both accumulators.
		getenv.Bool("STR", true)
		t.Match(getenv.LastErrs(), "parse")
		t.Nil(getenv.LastErr())
		t.Nil(getenv.LastErrs())
		// After multiple errors LastErrs returns all of them joined.
		getenv.Bool("STR", true)
		getenv.Int("STR", 5)
		t.Match(getenv.LastErrs(), "parse")
		// LastErr is also cleared after calling LastErrs.
		t.Nil(getenv.LastErr())
		t.Nil(getenv.LastErrs())
		// LastErr also clears LastErrs accumulator.
		getenv.Bool("STR", true)
		getenv.Int("STR", 5)
		t.Match(getenv.LastErr(), "parse")
		t.Nil(getenv.LastErrs())
	})
}
