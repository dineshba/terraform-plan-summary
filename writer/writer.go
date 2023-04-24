package writer

import (
	"github.com/dineshba/tf-summarize/terraform_state"
	"io"
)

type Writer interface {
	Write(writer io.Writer) error
}

func CreateWriter(tree, separateTree, drawable, mdEnabled, json bool, terraformState terraform_state.TerraformState) Writer {

	if tree {
		return NewTreeWriter(terraformState.ResourceChanges, drawable)
	}
	if separateTree {
		return NewSeparateTree(terraformState.AllResourceChanges(), drawable)
	}
	if json {
		return NewJsonWriter(terraformState.ResourceChanges)
	}

	return NewTableWriter(terraformState.AllResourceChanges(), terraformState.AllOutputChanges(), mdEnabled)
}
