package models

import (
	"testing"

	"gopkg.in/yaml.v2"
)

// ----------------------------
// --- InputModel -> EnvironmentItemModel

func TestConvertInputModel(t *testing.T) {
	valueOptions := []string{"option1", "option2"}

	input := InputModel{
		MappedTo:          "key",
		Title:             "title",
		Description:       "description",
		Value:             "value",
		ValueOptions:      valueOptions,
		IsRequired:        false,
		IsExpand:          false,
		IsDontChangeValue: false,
	}

	newInput := input.convert()

	key, value, err := newInput.GetKeyValuePair()
	if err != nil {
		t.Fatal("Falied to get input key, value:", newInput)
	}
	if key != "key" {
		t.Fatal("Falied to convert input key")
	}
	if value != "value" {
		t.Fatal("Falied to convert input value")
	}

	options, err := newInput.GetOptions()
	if err != nil {
		t.Fatal("Falied to get input options:", newInput)
	}
	if options.Title == nil || *options.Title != "title" {
		t.Fatal("Falied to convert options.Title")
	}
	if options.Description == nil || *options.Description != "description" {
		t.Fatal("Falied to convert options.Title")
	}
	if options.Summary == nil || *options.Summary != "" {
		t.Fatal("Falied to convert options.Summary")
	}
	if options.ValueOptions[0] != "option1" || options.ValueOptions[1] != "option2" {
		t.Fatal("Falied to convert options.ValueOptions")
	}
	if options.IsRequired == nil || *options.IsRequired != false {
		t.Fatal("Falied to convert options.IsRequired")
	}
	if options.IsExpand == nil || *options.IsExpand != false {
		t.Fatal("Falied to convert options.IsExpand")
	}
	if options.IsDontChangeValue == nil || *options.IsDontChangeValue != false {
		t.Fatal("Falied to convert options.IsDontChangeValue")
	}
}

func TestConvertStepModel(t *testing.T) {
	stringSlice := []string{"item1", "item2"}

	step := StepModel{
		ID:            "id",
		SteplibSource: "steplibSource",
		VersionTag:    "0.0.1",
		Name:          "name",
		Description:   "description",
		Website:       "http://website.com",
		ForkURL:       "http://fork.com",
		Source: map[string]string{
			"git": "http://git.com",
		},
		HostOsTags:          stringSlice,
		ProjectTypeTags:     stringSlice,
		TypeTags:            stringSlice,
		IsRequiresAdminUser: false,
	}

	newStep := step.convert()
	if newStep.Title == nil || *newStep.Title != "name" {
		t.Fatal("Failed to convert newStep.Title")
	}
	if newStep.Description == nil || *newStep.Description != "description" {
		t.Fatal("Failed to convert newStep.Description")
	}
	if newStep.Website == nil || *newStep.Website != "http://website.com" {
		t.Fatal("Failed to convert newStep.Website")
	}
	if newStep.Source.Git != "http://git.com" {
		t.Fatal("Failed to convert newStep.Source.Git")
	}
	if newStep.HostOsTags[0] != "item1" || newStep.HostOsTags[1] != "item2" {
		t.Fatal("Failed to convert newStep.HostOsTags")
	}
	if newStep.ProjectTypeTags[0] != "item1" || newStep.ProjectTypeTags[1] != "item2" {
		t.Fatal("Failed to convert newStep.ProjectTypeTags")
	}
	if newStep.TypeTags[0] != "item1" || newStep.TypeTags[1] != "item2" {
		t.Fatal("Failed to convert newStep.TypeTags")
	}
	if newStep.IsRequiresAdminUser == nil || *newStep.IsRequiresAdminUser != false {
		t.Fatal("Failed to convert newStep.IsRequiresAdminUser")
	}
}

func workflowModelFromYAMLBytes(bytes []byte) (workflow WorkflowModel, err error) {
	if err = yaml.Unmarshal(bytes, &workflow); err != nil {
		return
	}
	return
}

func TestConvertWorkflow(t *testing.T) {
	workflowString := `
format_version: 0.9.0
environments:
- title: TEST_ENV
  mapped_to: TEST_ENV
  is_expand: false
  value: converter
steps:
- id: bash-script-runner
  steplib_source: https://github.com/bitrise-io/bitrise-step-collection
  version_tag: 2.0.0
  source:
    git: https://github.com/bitrise-io/steps-bash-script.git

`
	workflow, err := workflowModelFromYAMLBytes([]byte(workflowString))
	if err != nil {
		t.Fatal("Failed to parse workflow:", err)
	}

	newWorkflow := workflow.convert()
	if len(newWorkflow.Steps) != 1 {
		t.Fatal("Failed to parse workflow:", err)
	}
	for _, stepListItem := range newWorkflow.Steps {
		for key := range stepListItem {
			if key != "_::https://github.com/bitrise-io/steps-bash-script.git@2.0.0" {
				t.Fatal("Failed to convert workflow: bad steplist item key:", key)
			}
		}
	}
}
