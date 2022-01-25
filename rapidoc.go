package rapidoc

import (
	"embed"
	_ "embed"
)


//go:embed rapidoc schemas
var Docs embed.FS
