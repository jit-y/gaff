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
			Type:       "action",
			LabelNames: []string{"actor", "name"},
		},
	},
}

type Strategy struct {
	actions []*Action
}

type Action struct {
	Name      string
	ActorName string
	Body      *hcl.Body
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
		case "action":
			action, actionDiags := decodeAction(block)

			strategy.actions = append(strategy.actions, action)
			diags = append(diags, actionDiags...)
		default:

		}
	}

	return strategy, diags
}

func decodeAction(block *hcl.Block) (*Action, hcl.Diagnostics) {
	a := &Action{
		Name:      block.Labels[0],
		ActorName: block.Labels[1],
		Body:      &block.Body,
	}

	diags := hcl.Diagnostics{}

	return a, diags
}
