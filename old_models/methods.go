package models

import (
	"fmt"

	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

// ----------------------------
// --- InputModel -> EnvironmentItemModel

func (input InputModel) getOptions() envmanModels.EnvironmentItemOptionsModel {
	return envmanModels.EnvironmentItemOptionsModel{
		Title:             pointers.NewStringPtr(input.Title),
		Description:       pointers.NewStringPtr(input.Description),
		ValueOptions:      input.ValueOptions,
		IsRequired:        pointers.NewBoolPtr(input.IsRequired),
		IsExpand:          pointers.NewBoolPtr(input.IsExpand),
		IsDontChangeValue: pointers.NewBoolPtr(input.IsDontChangeValue),
	}
}

func (input InputModel) convert() (envmanModels.EnvironmentItemModel, error) {
	environment := envmanModels.EnvironmentItemModel{
		input.MappedTo:          input.Value,
		envmanModels.OptionsKey: input.getOptions(),
	}
	if err := environment.FillMissingDefaults(); err != nil {
		return envmanModels.EnvironmentItemModel{}, err
	}
	return environment, nil
}

// ----------------------------
// --- OutputModel -> EnvironmentItemModel

func (output OutputModel) getOptions() envmanModels.EnvironmentItemOptionsModel {
	return envmanModels.EnvironmentItemOptionsModel{
		Title:       pointers.NewStringPtr(output.Title),
		Description: pointers.NewStringPtr(output.Description),
	}
}

func (output OutputModel) convert() (envmanModels.EnvironmentItemModel, error) {
	environment := envmanModels.EnvironmentItemModel{
		output.MappedTo:         "",
		envmanModels.OptionsKey: output.getOptions(),
	}

	if err := environment.FillMissingDefaults(); err != nil {
		return envmanModels.EnvironmentItemModel{}, err
	}

	return environment, nil
}

// ----------------------------
// --- old StepModel -> new StepModel

func (oldStep StepModel) getSource() stepmanModels.StepSourceModel {
	return stepmanModels.StepSourceModel{
		Git: oldStep.Source["git"],
	}
}

// GetStepLibIDVersionData ...
func (oldStep StepModel) GetStepLibIDVersionData() (string, string, string) {
	return oldStep.SteplibSource, oldStep.ID, oldStep.VersionTag
}

func (oldStep StepModel) getInputEnvironments() ([]envmanModels.EnvironmentItemModel, error) {
	inputs := []envmanModels.EnvironmentItemModel{}
	for _, oldInput := range oldStep.Inputs {
		newInput, err := oldInput.convert()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		inputs = append(inputs, newInput)
	}
	return inputs, nil
}

func (oldStep StepModel) getOutputEnvironments() ([]envmanModels.EnvironmentItemModel, error) {
	outputs := []envmanModels.EnvironmentItemModel{}
	for _, oldOutput := range oldStep.Outputs {
		newOutput, err := oldOutput.convert()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		outputs = append(outputs, newOutput)
	}
	return outputs, nil
}

// Convert ...
func (oldStep StepModel) Convert() (stepmanModels.StepModel, error) {
	inputs, err := oldStep.getInputEnvironments()
	if err != nil {
		return stepmanModels.StepModel{}, err
	}

	outputs, err := oldStep.getOutputEnvironments()
	if err != nil {
		return stepmanModels.StepModel{}, err
	}

	newStep := stepmanModels.StepModel{
		Title:               pointers.NewStringPtr(oldStep.Name),
		Description:         pointers.NewStringPtr(oldStep.Description),
		Website:             pointers.NewStringPtr(oldStep.Website),
		Source:              oldStep.getSource(),
		HostOsTags:          oldStep.HostOsTags,
		ProjectTypeTags:     oldStep.ProjectTypeTags,
		TypeTags:            oldStep.TypeTags,
		IsRequiresAdminUser: pointers.NewBoolPtr(oldStep.IsRequiresAdminUser),
		Inputs:              inputs,
		Outputs:             outputs,
	}

	return newStep, nil
}

// GetInputByKey ...
func (oldStep StepModel) GetInputByKey(key string) (InputModel, error) {
	for _, input := range oldStep.Inputs {
		if input.MappedTo == key {
			return input, nil
		}
	}
	return InputModel{}, fmt.Errorf("No input found for key (%s)", key)
}

// ----------------------------
// --- old WorkflowModel -> new StepModel

// GetEnvironments ...
func (oldWorkflow WorkflowModel) GetEnvironments() ([]envmanModels.EnvironmentItemModel, error) {
	environments := []envmanModels.EnvironmentItemModel{}
	for _, oldEnv := range oldWorkflow.Environments {
		newEnv, err := oldEnv.convert()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		environments = append(environments, newEnv)
	}
	return environments, nil
}
