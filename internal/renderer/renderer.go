package renderer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// Render executes the named template from tmplFS with data, writing the result to outPath.
// Parent directories of outPath are created as needed.
func Render(tmplFS fs.FS, tmplName string, data any, outPath string) error {
	tmplContent, err := fs.ReadFile(tmplFS, tmplName)
	if err != nil {
		return fmt.Errorf("renderer: read template %q: %w", tmplName, err)
	}

	tmpl, err := template.New(tmplName).Funcs(sprig.TxtFuncMap()).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("renderer: parse template %q: %w", tmplName, err)
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return fmt.Errorf("renderer: create output dir: %w", err)
	}

	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("renderer: create output file %q: %w", outPath, err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("renderer: execute template %q: %w", tmplName, err)
	}
	return nil
}
