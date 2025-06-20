// embed.go
//
// Author: 0x716
// Created: 2025-06-18
//
// License: MIT License
//
// Description:
// This file handles embedding the static 'dist' directory into the binary using Go's embed package.
// The embed pattern and usage references the implementation style of PocketBase (https://github.com/pocketbase/pocketbase).
//
// Note:
// The embedded files are served without the "dist" prefix by using fs.Sub.
//
// ------------------------------------------------------------

package web

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var distDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS, _ = fs.Sub(distDir, "dist")
