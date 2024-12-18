package ex

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/zrcoder/amisgo/comp"
)

//go:embed *
var FS embed.FS

const defaultCodeKey = "Hello, Go+"

func Get() (options []any, defaultCode string, err error) {
	var es []fs.DirEntry
	es, err = FS.ReadDir(".")
	if err != nil {
		return
	}
	for _, f := range es {
		if !strings.HasSuffix(f.Name(), ".gop") {
			continue
		}
		data, err := FS.ReadFile(f.Name())
		if err != nil {
			return nil, "", err
		}
		key := strings.TrimSuffix(filepath.Base(f.Name()), filepath.Ext(f.Name()))
		val := string(data)
		if key == defaultCodeKey {
			defaultCode = val
		}
		options = append(options, comp.Option().Label(key).Value(val))
	}
	return
}
