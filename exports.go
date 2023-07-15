package runtime

import (
	"runtime/pkg"
)

type (
	Runtime = pkg.Runtime
	Service = pkg.RuntimeServiceInterface
)

var New = pkg.New
var _ = New
