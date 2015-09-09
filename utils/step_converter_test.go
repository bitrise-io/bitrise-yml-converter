package utils

import (
	"testing"

	"gopkg.in/yaml.v2"

	envmanModels "github.com/bitrise-io/envman/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

func TestConvertStepsInputs(t *testing.T) {
	originalInputsStr := `
envs:
- KEY1: value1
  opts:
    title: title
    description: description
    summary: summary
    value_options:
    - option1
    is_required: false
    is_expand: false
    is_dont_change_value: false
- KEY3: value3
- KEY4: value4
`

	diffInputsStr := `
envs:
- KEY2:
  opts:
    title: title2
    description: description2
    summary: summary2
    value_options:
    - option3
    is_required: true
    is_expand: true
    is_dont_change_value: true
`

	conversionMap := map[string]string{
		"KEY1": "KEY2",
	}

	originalEnvs := envmanModels.EnvsYMLModel{}
	if err := yaml.Unmarshal([]byte(originalInputsStr), &originalEnvs); err != nil {
		t.Fatal("Failed to yaml.Unmarshal, err:", err)
	}

	diffEnvs := envmanModels.EnvsYMLModel{}
	if err := yaml.Unmarshal([]byte(diffInputsStr), &diffEnvs); err != nil {
		t.Fatal("Failed to yaml.Unmarshal, err:", err)
	}

	convertedEnvs, err := convertStepsInputs(originalEnvs.Envs, diffEnvs.Envs, conversionMap)
	if err != nil {
		t.Fatal("Failed to convert envs, err:", err)
	}

	// Test length
	if len(convertedEnvs) != 3 {
		t.Fatal("Failed to convert envs")
	}

	convertedEnv := convertedEnvs[0]

	// Test key-value conversion
	key, value, err := convertedEnv.GetKeyValuePair()
	if err != nil {
		t.Fatal("Failed to get env key-value pair, err:", err)
	}

	if key != "KEY1" {
		t.Fatal("Failed to convert env key")
	}
	if value != "value1" {
		t.Fatal("Failed to convert env value")
	}

	// Test options conversion
	convertedOptions, err := convertedEnv.GetOptions()
	if err != nil {
		t.Fatal("Failed to get env options, err:", err)
	}

	if *convertedOptions.Title != "title" {
		t.Fatal("Failed to convert Title")
	}
	if *convertedOptions.Description != "description" {
		t.Fatal("Failed to convert Description")
	}
	if *convertedOptions.Summary != "summary" {
		t.Fatal("Failed to convert Summary")
	}
	if convertedOptions.ValueOptions[0] != "option1" {
		t.Fatal("Failed to convert ValueOptions")
	}
	if *convertedOptions.IsRequired != false {
		t.Fatal("Failed to convert IsRequired")
	}
	if *convertedOptions.IsExpand != true {
		t.Fatal("Failed to convert IsExpand")
	}
	if *convertedOptions.IsDontChangeValue != false {
		t.Fatal("Failed to convert IsDontChangeValue")
	}
}

func TestConvertStep(t *testing.T) {
	diffStepStr := `
title: title
is_always_run: true
`

	diffStep := stepmanModels.StepModel{}
	if err := yaml.Unmarshal([]byte(diffStepStr), &diffStep); err != nil {
		t.Fatal("Failed to yaml.Unmarshal, err:", err)
	}

	convertedStep, _, err := ConvertStep(diffStep, "script", map[string]string{})
	if err != nil {
		t.Fatal("Failed to convert step, err:", err)
	}

	if *convertedStep.Title != "title" {
		t.Fatal("Failed to convert Title")
	}
	if *convertedStep.IsAlwaysRun != true {
		t.Fatal("Failed to convert IsAlwaysRun")
	}
}
