package runtime

import (
	"github.com/gravestench/runtime/pkg"
)

/*
	these are just some exports to:
	- prevent you from having to know to import from pkg
	- make the interfaces less wordy in your code
*/

type (
	Runtime = pkg.RuntimeInterface
	R       = Runtime // for even more brevity

	Service = pkg.RuntimeServiceInterface
	S       = Service
)

var New = pkg.New
var _ = New
