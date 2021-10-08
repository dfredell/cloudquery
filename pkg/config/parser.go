package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type SourceType string

const (
	SourceJSON = "json"
	SourceHCL  = "hcl"
)

// EnvVarPrefix is a prefix for environment variable names to be exported for HCL substitution.
const EnvVarPrefix = "CQ_VAR_"

// Parser is the main interface to read configuration files and other related
// files from disk.
//
// It retains a cache of all files that are loaded so that they can be used
// to create source code snippets in diagnostics, etc.
type Parser struct {
	fs         afero.Afero
	p          *hclparse.Parser
	HCLContext hcl.EvalContext
}

type Option func(*Parser)

func WithFS(fs afero.Fs) Option {
	return func(p *Parser) {
		p.fs = afero.Afero{Fs: fs}
	}
}

// WithEnvironmentVariables fills hcl.Context with values of environment variables given in vars.
// Only variables that start with given prefix are considered. Prefix is removed from the name and
// the name is lower cased then.
func WithEnvironmentVariables(prefix string, vars []string) Option {
	return func(p *Parser) {
		for _, e := range vars {
			pair := strings.SplitN(e, "=", 2)
			if strings.HasPrefix(pair[0], prefix) {
				name := strings.ToLower(pair[0][len(prefix):])
				p.HCLContext.Variables[name] = cty.StringVal(pair[1])
			}
		}
	}
}

// NewParser creates and returns a new Parser.
func NewParser(options ...Option) *Parser {
	p := Parser{
		fs: afero.Afero{Fs: afero.OsFs{}},
		p:  hclparse.NewParser(),
		HCLContext: hcl.EvalContext{
			Variables: make(map[string]cty.Value),
			Functions: make(map[string]function.Function),
		},
	}

	for _, opt := range options {
		opt(&p)
	}
	return &p
}

// LoadHCLFile is a low-level method that reads the file at the given path,
// parses it, and returns the hcl.Body representing its root. In many cases
// it is better to use one of the other Load*File methods on this type,
// which additionally decode the root body in some way and return a higher-level
// construct.
//
// If the file cannot be read at all -- e.g. because it does not exist -- then
// this method will return a nil body and error diagnostics. In this case
// callers may wish to ignore the provided error diagnostics and produce
// a more context-sensitive error instead.
//
// The file will be parsed using the HCL native syntax unless the filename
// ends with ".json", in which case the HCL JSON syntax will be used.
func (p *Parser) LoadHCLFile(path string) (hcl.Body, hcl.Diagnostics) {
	src, err := p.fs.ReadFile(path)

	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	return p.loadFromSource(path, src, SourceType(filepath.Ext(path)))
}

func (p *Parser) loadFromSource(name string, data []byte, ext SourceType) (hcl.Body, hcl.Diagnostics) {
	var file *hcl.File
	var diags hcl.Diagnostics
	switch ext {
	case SourceJSON:
		file, diags = p.p.ParseJSON(data, name)
	default:
		file, diags = p.p.ParseHCL(data, name)
	}
	// If the returned file or body is nil, then we'll return a non-nil empty
	// body so we'll meet our contract that nil means an error reading the file.
	if file == nil || file.Body == nil {
		return hcl.EmptyBody(), diags
	}

	return file.Body, diags
}
