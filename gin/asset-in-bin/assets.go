package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets07ade61300579b200cc19f178b6b0caea1051d7a = "<!doctype html>\n<body>\n<p>Can you see this? → {{.Bar}}</p>\n</body>"
var _Assets1da9cabbab97644f0a2ff16a82ed89c8b1d2d446 = "<!doctype html>\n<body>\n<p>Can you see this? → {{.Foo}}</p>\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template"}, "/template": []string{"bar.html", "foo.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1687840777, 1687840777088176055),
		Data:     nil,
	}, "/template": &assets.File{
		Path:     "/template",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1687840718, 1687840718999425721),
		Data:     nil,
	}, "/template/bar.html": &assets.File{
		Path:     "/template/bar.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1687840718, 1687840718999320319),
		Data:     []byte(_Assets07ade61300579b200cc19f178b6b0caea1051d7a),
	}, "/template/foo.html": &assets.File{
		Path:     "/template/foo.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1687840702, 1687840702184597533),
		Data:     []byte(_Assets1da9cabbab97644f0a2ff16a82ed89c8b1d2d446),
	}}, "")
