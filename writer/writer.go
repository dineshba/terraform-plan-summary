package writer

import (
	"io"

	"github.com/dineshba/tf-summarize/terraformstate"
	tfjson "github.com/hashicorp/terraform-json"
)

type Writer interface {
	Write(writer io.Writer) error
}

func CreateWriter(tree, separateTree, drawable, mdEnabled, json bool, plan tfjson.Plan) Writer {

	if tree {
		return NewTreeWriter(plan.ResourceChanges, drawable)
	}
	if separateTree {
		return NewSeparateTree(terraformstate.GetAllResourceChanges(plan), drawable)
	}
	if json {
		return NewJSONWriter(plan.ResourceChanges)
	}

	return NewTableWriter(terraformstate.GetAllResourceChanges(plan), terraformstate.GetAllOutputChanges(plan), mdEnabled)
}
