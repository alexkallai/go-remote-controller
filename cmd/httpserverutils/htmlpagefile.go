package httpserverutils

import "embed"

// Embed necessary files into built exe
//go:embed index.html
var f embed.FS

var htmlFileTemplate, _ = f.ReadFile("index.html")
