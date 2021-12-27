package proto

import (
	"embed"
)

//go:embed openapiv2/*
var OpenApiV2Fs embed.FS
