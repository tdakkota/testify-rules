package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

//doc:summary Checks that expected and actual arguments are ordered correctly.
//doc:before  require.Equal(t, value, 0)
//doc:after   require.Equal(t, 0, value)
//doc:tags    diagnostic
func expectedIsFirstArg(m dsl.Matcher) {
	m.Import("github.com/stretchr/testify/assert")
	m.Import("github.com/stretchr/testify/require")

	m.Match(`$pkg.$method($_, $actual, $expected)`).
		Where(
			m["method"].Text.Matches(
				"^(?:Not)?(Equal|EqualValues|Same|Exactly|InDelta|InEpsilon|WithinDuration)(?:f)?$",
			) &&
				m["expected"].Const &&
				m["pkg"].Object.Is("PkgName") &&
				m["pkg"].Text.Matches("require|assert"),
		).
		Suggest("$pkg.$method($expected, $actual)").
		Report("$expected passed as actual argument, probably you should swap arguments")

	m.Match(`$t.$method($actual, $expected)`).
		Where(
			m["method"].Text.Matches(
				"^(?:Not)?(Equal|EqualValues|Same|Exactly|InDelta|InEpsilon|WithinDuration)(?:f)?$",
			) &&
				m["expected"].Const &&
				m["t"].Type.Is("*assert.Assertions") ||
				m["t"].Type.Is("*require.Assertions"),
		).
		Suggest("$t.$method($expected, $actual)").
		Report("$expected passed as actual argument, probably you should swap arguments")
}
