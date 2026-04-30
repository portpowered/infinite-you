package ui

import (
	"embed"
	"io/fs"
)

const (
	// BasePath is the HTTP route prefix used when the embedded dashboard shell
	// is served by the agent-factory API server.
	BasePath = "/dashboard/ui"
)

var (
	//go:embed dist dist/*
	embeddedDist embed.FS
)

// DistFS returns the embedded production dashboard assets rooted at dist/.
func DistFS() (fs.FS, error) {
	return fs.Sub(embeddedDist, "dist")
}
