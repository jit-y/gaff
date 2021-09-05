package gaff

import (
	"io/fs"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Parser struct {
	fs fs.FS
}

var strategySchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "tactic",
			LabelNames: []string{"provider", "name"},
		},
	},
}

type Strategy struct {
	tactics []*Tactic
}

type Tactic struct {
	Name         string
	ProviderName string
	Config       *hcl.Body
}

func NewParser(f fs.FS) *Parser {
	return &Parser{
		fs: f,
	}
}

func (p *Parser) LoadHCLFile(filepath string) (*Strategy, hcl.Diagnostics) {
	src, err := fs.ReadFile(p.fs, filepath)
	if err != nil {
		return nil, hcl.Diagnostics{}
	}

	hp := hclparse.NewParser()
	file, diags := hp.ParseHCL(src, filepath)
	if diags.HasErrors() {
		return nil, diags
	}

	b := file.Body

	strategy := &Strategy{}
	content, contentDiags := b.Content(strategySchema)
	diags = append(diags, contentDiags...)

	for _, block := range content.Blocks {
		switch block.Type {
		case "tactic":
			tactic, tacticDiags := decodeTactic(block)

			strategy.tactics = append(strategy.tactics, tactic)
			diags = append(diags, tacticDiags...)
		default:

		}
	}

	return strategy, diags
}

func decodeTactic(block *hcl.Block) (*Tactic, hcl.Diagnostics) {
	t := &Tactic{
		Name:         block.Labels[0],
		ProviderName: block.Labels[1],
		Config:       &block.Body,
	}

	diags := hcl.Diagnostics{}

	return t, diags
}
