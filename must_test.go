package getenv_test

import (
	"testing"
	"time"

	"github.com/powerman/check"

	"github.com/powerman/getenv"
)

func TestMustBool(t *testing.T) {
	t.Run("unset", func(tt *testing.T) {
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustBool("MUST_BOOL_UNSET") }, "is not set")
	})
	t.Run("invalid", func(tt *testing.T) {
		tt.Setenv("MUST_BOOL_INVALID", "bad")
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustBool("MUST_BOOL_INVALID") }, "bad")
	})
	t.Run("true", func(tt *testing.T) {
		tt.Setenv("MUST_BOOL_TRUE", "true")
		t := check.T(tt)
		t.Equal(getenv.MustBool("MUST_BOOL_TRUE"), true)
	})
	t.Run("false", func(tt *testing.T) {
		tt.Setenv("MUST_BOOL_FALSE", "false")
		t := check.T(tt)
		t.Equal(getenv.MustBool("MUST_BOOL_FALSE"), false)
	})
}

func TestMustDur(t *testing.T) {
	t.Run("unset", func(tt *testing.T) {
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustDur("MUST_DUR_UNSET") }, "is not set")
	})
	t.Run("invalid", func(tt *testing.T) {
		tt.Setenv("MUST_DUR_INVALID", "notdur")
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustDur("MUST_DUR_INVALID") }, "notdur")
	})
	t.Run("valid", func(tt *testing.T) {
		tt.Setenv("MUST_DUR_VALID", "5s")
		t := check.T(tt)
		t.Equal(getenv.MustDur("MUST_DUR_VALID"), 5*time.Second)
	})
}

func TestMustFloat(t *testing.T) {
	t.Run("unset", func(tt *testing.T) {
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustFloat("MUST_FLOAT_UNSET") }, "is not set")
	})
	t.Run("invalid", func(tt *testing.T) {
		tt.Setenv("MUST_FLOAT_INVALID", "notfloat")
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustFloat("MUST_FLOAT_INVALID") }, "notfloat")
	})
	t.Run("valid", func(tt *testing.T) {
		tt.Setenv("MUST_FLOAT_VALID", "3.14")
		t := check.T(tt)
		t.Equal(getenv.MustFloat("MUST_FLOAT_VALID"), 3.14)
	})
}

func TestMustInt(t *testing.T) {
	t.Run("unset", func(tt *testing.T) {
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustInt("MUST_INT_UNSET") }, "is not set")
	})
	t.Run("invalid", func(tt *testing.T) {
		tt.Setenv("MUST_INT_INVALID", "notint")
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustInt("MUST_INT_INVALID") }, "notint")
	})
	t.Run("valid", func(tt *testing.T) {
		tt.Setenv("MUST_INT_VALID", "7")
		t := check.T(tt)
		t.Equal(getenv.MustInt("MUST_INT_VALID"), 7)
	})
}

func TestMustStr(t *testing.T) {
	t.Run("unset", func(tt *testing.T) {
		t := check.T(tt)
		t.PanicMatch(func() { getenv.MustStr("MUST_STR_UNSET") }, "is not set")
	})
	t.Run("empty", func(tt *testing.T) {
		tt.Setenv("MUST_STR_EMPTY", "")
		t := check.T(tt)
		t.Equal(getenv.MustStr("MUST_STR_EMPTY"), "")
	})
	t.Run("valid", func(tt *testing.T) {
		tt.Setenv("MUST_STR_VALID", "hello")
		t := check.T(tt)
		t.Equal(getenv.MustStr("MUST_STR_VALID"), "hello")
	})
}
