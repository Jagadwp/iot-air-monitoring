package web

import "embed"

//go:embed *.html template/*.html
var Files embed.FS
