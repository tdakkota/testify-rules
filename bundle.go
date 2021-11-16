package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// Bundle holds the rules package metadata.
//
// In order to be importable from other gorules package,
// a package must define a Bundle variable.
var Bundle = dsl.Bundle{}
