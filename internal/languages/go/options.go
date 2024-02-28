package _go

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/go-openapi/swag"

	"github.com/khizunov/go-openapi-generator/internal/templateutils"
)

// LanguageOpts to describe a language to the code generator
type LanguageOpts struct {
	ReservedWords        []string
	BaseImportFunc       func(string) string               `json:"-"`
	ImportsFunc          func(map[string]string) string    `json:"-"`
	ArrayInitializerFunc func(interface{}) (string, error) `json:"-"`
	reservedWordsSet     map[string]struct{}
	initialized          bool
	formatFunc           func(string, []byte) ([]byte, error)
	fileNameFunc         func(string) string // language specific source file naming rules
	dirNameFunc          func(string) string // language specific directory naming rules
}

// Init the language option
func (l *LanguageOpts) Init() {
	if l.initialized {
		return
	}
	l.initialized = true
	l.reservedWordsSet = make(map[string]struct{})
	for _, rw := range l.ReservedWords {
		l.reservedWordsSet[rw] = struct{}{}
	}
}

// MangleName makes sure a reserved word gets a safe name
func (l *LanguageOpts) MangleName(name, suffix string) string {
	if _, ok := l.reservedWordsSet[swag.ToFileName(name)]; !ok {
		return name
	}
	return strings.Join([]string{name, suffix}, "_")
}

// MangleVarName makes sure a reserved word gets a safe name
func (l *LanguageOpts) MangleVarName(name string) string {
	nm := swag.ToVarName(name)
	if _, ok := l.reservedWordsSet[nm]; !ok {
		return nm
	}
	return nm + "Var"
}

// MangleFileName makes sure a file name gets a safe name
func (l *LanguageOpts) MangleFileName(name string) string {
	if l.fileNameFunc != nil {
		return l.fileNameFunc(name)
	}
	return swag.ToFileName(name)
}

// ManglePackageName makes sure a package gets a safe name.
// In case of a file system path (e.g. name contains "/" or "\" on Windows), this return only the last element.
func (l *LanguageOpts) ManglePackageName(name, suffix string) string {
	if name == "" {
		return suffix
	}
	if l.dirNameFunc != nil {
		name = l.dirNameFunc(name)
	}
	pth := filepath.ToSlash(filepath.Clean(name))
	pkg := importAlias(pth)
	return l.MangleName(swag.ToFileName(templateutils.PrefixForName(pkg)+pkg), suffix)
}

// ManglePackagePath makes sure a full package path gets a safe name.
// Only the last part of the path is altered.
func (l *LanguageOpts) ManglePackagePath(name string, suffix string) string {
	if name == "" {
		return suffix
	}
	target := filepath.ToSlash(filepath.Clean(name)) // preserve path
	parts := strings.Split(target, "/")
	parts[len(parts)-1] = l.ManglePackageName(parts[len(parts)-1], suffix)
	return strings.Join(parts, "/")
}

// FormatContent formats a file with a language specific formatter
func (l *LanguageOpts) FormatContent(name string, content []byte) ([]byte, error) {
	if l.formatFunc != nil {
		return l.formatFunc(name, content)
	}
	return content, nil
}

// Imports generate the code to import some external packages, possibly aliased
func (l *LanguageOpts) Imports(imports map[string]string) string {
	if l.ImportsFunc != nil {
		return l.ImportsFunc(imports)
	}
	return ""
}

// ArrayInitializer builds a litteral array
func (l *LanguageOpts) ArrayInitializer(data interface{}) (string, error) {
	if l.ArrayInitializerFunc != nil {
		return l.ArrayInitializerFunc(data)
	}
	return "", nil
}

// baseImport figures out the base path to generate import statements
func (l *LanguageOpts) baseImport(tgt string) string {
	if l.BaseImportFunc != nil {
		return l.BaseImportFunc(tgt)
	}
	return ""
}

func importAlias(pkg string) string {
	_, k := path.Split(pkg)
	return k
}
