// Code generated by genprefer, DO NOT EDIT.

package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary Prefer require.Len instead of comparing length.
//doc:before  require.Equal(t, length, len(a))
//doc:after   require.Len(t, a, length)
//doc:tags    diagnostic
func preferLen(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.Equal($t, $length, len($a))`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Len($t, $a, $length)")

	m.Match(`$t.Equal($length, len($a))`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Len($a, $length)")

	m.Match(`$pkg.Equal($t, len($a), $length)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Len($t, $a, $length)")

	m.Match(`$t.Equal(len($a), $length)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Len($a, $length)")
}

//doc:summary Prefer require.Empty instead of comparing length.
//doc:before  require.Equal(t, 0, len(a))
//doc:after   require.Empty(t, a)
//doc:tags    diagnostic
func preferEmpty(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.Equal($t, 0, len($a))`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Empty($t, $a)")

	m.Match(`$t.Equal(0, len($a))`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Empty($a)")

	m.Match(`$pkg.Equal($t, len($a), 0)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Empty($t, $a)")

	m.Match(`$t.Equal(len($a), 0)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Empty($a)")
}

//doc:summary Prefer require.Nil instead of comparing to nil.
//doc:before  require.Equal(t, nil, a)
//doc:after   require.Nil(t, a)
//doc:tags    diagnostic
func preferNil(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.Equal($t, nil, $a)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Nil($t, $a)")

	m.Match(`$t.Equal(nil, $a)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Nil($a)")

	m.Match(`$pkg.Same($t, nil, $a)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Nil($t, $a)")

	m.Match(`$t.Same(nil, $a)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Nil($a)")

	m.Match(`$pkg.Equal($t, $a, nil)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Nil($t, $a)")

	m.Match(`$t.Equal($a, nil)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Nil($a)")

	m.Match(`$pkg.Same($t, $a, nil)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Nil($t, $a)")

	m.Match(`$t.Same($a, nil)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Nil($a)")
}

//doc:summary Prefer require.NotNil instead of comparing to nil.
//doc:before  require.NotEqual(t, nil, a)
//doc:after   require.NotNil(t, a)
//doc:tags    diagnostic
func preferNotNil(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.NotEqual($t, nil, $a)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.NotNil($t, $a)")

	m.Match(`$t.NotEqual(nil, $a)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.NotNil($a)")

	m.Match(`$pkg.NotEqual($t, $a, nil)`).
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.NotNil($t, $a)")

	m.Match(`$t.NotEqual($a, nil)`).
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.NotNil($a)")
}

//doc:summary Prefer require.NoError instead of comparing to nil.
//doc:before  require.NotNil(t, a)
//doc:after   require.Error(t, a)
//doc:tags    diagnostic
func preferError(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.NotNil($t, $a)`).
		Where(
			(m["a"].Type.Is("error")) &&
				m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Error($t, $a)")

	m.Match(`$t.NotNil($a)`).
		Where(
			(m["a"].Type.Is("error")) &&
				m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Error($a)")
}

//doc:summary Prefer require.NoError instead of comparing to nil.
//doc:before  require.Nil(t, a)
//doc:after   require.NoError(t, a)
//doc:tags    diagnostic
func preferNoError(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.Nil($t, $a)`).
		Where(
			(m["a"].Type.Is("error")) &&
				m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.NoError($t, $a)")

	m.Match(`$t.Nil($a)`).
		Where(
			(m["a"].Type.Is("error")) &&
				m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.NoError($a)")
}
