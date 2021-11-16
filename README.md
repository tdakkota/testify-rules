# testify-rules

[Ruleguard](https://github.com/quasilyte/go-ruleguard) ruleset bundle for [testify](https://github.com/stretchr/testify).

## Install

See [ruleguard documentation](https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#ruleguard-bundles).
```go
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	testifyrules "github.com/tdakkota/testify-rules"
)

func init() {
	// Imported rules will have a "" prefix.
	dsl.ImportRules("", testifyrules.Bundle)
}
```

