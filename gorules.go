package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// Bundle holds the rules package metadata.
//
// In order to be importable from other gorules package,
// a package must define a Bundle variable.
var Bundle = dsl.Bundle{}

//doc:summary Prefer require.Len instead of comparing length.
//doc:before  require.Equal(t, 1, len(a))
//doc:after   require.Len(t, a, 1)
//doc:tags    diagnostic
func preferLen(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match("$pkg.Equal($t, $len, len($a))").
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Len($t, $a, $len)")

	m.Match("$t.Equal($len, len($a))").
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Len($a, $len)")
}

//doc:summary Prefer require.Empty instead of comparing length to zero.
//doc:before  require.Len(t, 0, a)
//doc:after   require.Empty(t, a)
//doc:tags    diagnostic
func preferEmpty(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match("$pkg.Len($t, 0, $a)").
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Empty($t, $a)")

	m.Match("$t.Len($a, 0)").
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

	m.Match("$pkg.Equal($t, nil, $a)").
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Nil($t, $a)")

	m.Match("$t.Equal(nil, $a)").
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Nil($a)")
}

//doc:summary Prefer require.NotNil instead of comparing to nil.
//doc:before  require.NotEqual(t, nil, a)
//doc:after   require.Nil(t, a)
//doc:tags    diagnostic
func preferNotNil(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match("$pkg.NotEqual($t, nil, $a)").
		Where(
			m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.NotNil($t, $a)")

	m.Match("$t.NotEqual(nil, $a)").
		Where(
			m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.NotNil($a)")
}

//doc:summary Prefer require.Error/NoError instead of comparing errors.
//doc:before  require.Nil(t, err)
//doc:after   require.NoError(t, err)
//doc:tags    diagnostic
func preferError(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match("$pkg.Nil($t, $a)").
		Where(
			m["a"].Type.Is("error") &&
				m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.NoError($t, $a)")

	m.Match("$t.Nil($a)").
		Where(
			m["a"].Type.Is("error") &&
				m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.NoError($a)")

	m.Match("$pkg.NotNil($t, $a)").
		Where(
			m["a"].Type.Is("error") &&
				m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).Suggest("$pkg.Error($t, $a)")

	m.Match("$t.NotNil($a)").
		Where(
			m["a"].Type.Is("error") &&
				m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).Suggest("$t.Error($a)")
}
