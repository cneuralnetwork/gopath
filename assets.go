package gopath

import "embed"

//go:embed content web/dist
var Assets embed.FS

func FS() embed.FS { return Assets }
