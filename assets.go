package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsecc6a17c77d7edf46353ad00b6e9157909587934 = "Date: {[{.now | formatAsDate}]}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": {"templates"}, "/templates": {"raw.tmpl"}}, map[string]*assets.File{
	"/": {
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1651414785, 1651414785980000000),
		Data:     nil,
	}, "/templates": {
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1651324357, 1651324357677918900),
		Data:     nil,
	}, "/templates/raw.tmpl": {
		Path:     "/templates/raw.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1651324357, 1651324357677918900),
		Data:     []byte(_Assetsecc6a17c77d7edf46353ad00b6e9157909587934),
	}}, "")
