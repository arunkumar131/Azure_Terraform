package e2e

import (
	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/packer"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"

	helper "github.com/Azure/terraform-module-test-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

var speicalTests = map[string]func(*testing.T){
	"quickstart/201-vmss-packer-jumpbox": test201VmssPackerJumpbox,
}

func Test_Quickstarts(t *testing.T) {
	msiId := os.Getenv("MSI_ID")
	if msiId != "" {
		_ = os.Setenv("TF_VAR_msi_id", msiId)
	}
	input := os.Getenv("CHANGED_FOLDERS")
	folders := strings.Split(input, ",")
	if input == "" {
		var err error
		folders, err = allExamples()
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	for _, f := range folders {
		f = strings.TrimSpace(f)
		if filepath.Dir(f) != "quickstart" {
			continue
		}
		rootPath := filepath.Join("..", "..")
		path := filepath.Join(rootPath, f)
		if !files.IsExistingDir(path) {
			continue
		}
		test, ok := speicalTests[f]
		if !ok {
			test = func(t *testing.T) {
				helper.RunE2ETest(t, rootPath, f, terraform.Options{
					Upgrade: true,
				}, nil)
			}
		}

		t.Run(f, test)
	}
}

func allExamples() ([]string, error) {
	examples, err := os.ReadDir("../../quickstart")
	if err != nil {
		return nil, err
	}
	var r []string
	for _, f := range examples {
		if !f.IsDir() {
			continue
		}
		r = append(r, filepath.Join("quickstart", f.Name()))
	}
	return r, nil
}

func test201VmssPackerJumpbox(t *testing.T) {
	examplePath := filepath.Join("..", "..", "quickstart", "201-vmss-packer-jumpbox")
	harnessPath := filepath.Join(examplePath, "test_harness")
	harnessOptions := &terraform.Options{
		TerraformDir: harnessPath,
	}
	defer terraform.Destroy(t, harnessOptions)
	terraform.InitAndApply(t, harnessOptions)
	harnessOutput := terraform.OutputAll(t, harnessOptions)
	imageResourceGroupName := harnessOutput["resource_group_name"].(string)
	pkrCfg := filepath.Join("..", "..", "quickstart", "201-vmss-packer-jumpbox", "ubuntu.pkr.hcl")
	packerVars := map[string]string{
		"image_resource_group_name": imageResourceGroupName,
	}
	if identityId := os.Getenv("MSI_ID"); identityId != "" {
		packerVars["client_id"] = identityId
	}
	if subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID"); subscriptionId != "" {
		packerVars["subscription_id"] = subscriptionId
	}
	_, err := packer.BuildArtifactE(t, &packer.Options{
		Template:   pkrCfg,
		Vars:       packerVars,
		VarFiles:   nil,
		WorkingDir: examplePath,
	})
	require.NoError(t, err)
	helper.RunE2ETest(t, examplePath, "", terraform.Options{
		Upgrade: true,
	}, nil)
}
