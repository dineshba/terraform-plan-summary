package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	tfjson "github.com/hashicorp/terraform-json"
)

type BinaryParser struct {
	fileName string
}

func (j BinaryParser) Parse() (tfjson.Plan, error) {
	tfbinary := "terraform"
	if tfoverride, ok := os.LookupEnv("TF_BINARY"); ok {
		tfbinary = tfoverride
	}
	cmd := exec.Command(tfbinary, "show", "-json", j.fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return tfjson.Plan{}, fmt.Errorf(
			"error when running 'terraform show -json %s': \n%s\n\n%s",
			j.fileName, output, "Make sure you are running in terraform directory and terraform init is done")
	}
	plan := tfjson.Plan{}
	err = json.Unmarshal(output, &plan)
	if err != nil {
		return tfjson.Plan{}, fmt.Errorf("error when parsing input: %s", err.Error())
	}
	return plan, nil
}

func NewBinaryParser(fileName string) Parser {
	return BinaryParser{
		fileName: fileName,
	}
}
