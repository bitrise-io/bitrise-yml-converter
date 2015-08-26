package models

import (
	bitriseModels "github.com/bitrise-io/bitrise/models/models_1_0_0"
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

func (input InputModel) convert() envmanModels.EnvironmentItemModel {
	environment := envmanModels.EnvironmentItemModel{
		input.MappedTo:          input.Value,
		envmanModels.OptionsKey: input.getOptions(),
	}
	environment.FillMissingDefaults()
	return environment
}

// ----------------------------
// --- OutputModel -> EnvironmentItemModel

func (output OutputModel) getOptions() envmanModels.EnvironmentItemOptionsModel {
	return envmanModels.EnvironmentItemOptionsModel{
		Title:       pointers.NewStringPtr(output.Title),
		Description: pointers.NewStringPtr(output.Description),
	}
}

func (output OutputModel) convert() envmanModels.EnvironmentItemModel {
	environment := envmanModels.EnvironmentItemModel{
		output.MappedTo:         "",
		envmanModels.OptionsKey: output.getOptions(),
	}
	environment.FillMissingDefaults()
	return environment
}

// ----------------------------
// --- old StepModel -> new StepModel

func (oldStep StepModel) getSource() stepmanModels.StepSourceModel {
	return stepmanModels.StepSourceModel{
		Git: oldStep.Source["git"],
	}
}

func (oldStep StepModel) getStepLibIDVersionData() (string, string, string) {
	return oldStep.SteplibSource, oldStep.ID, oldStep.VersionTag
}

func (oldStep StepModel) getInputEnvironments() []envmanModels.EnvironmentItemModel {
	inputs := []envmanModels.EnvironmentItemModel{}
	for _, oldInput := range oldStep.Inputs {
		newInput := oldInput.convert()
		inputs = append(inputs, newInput)
	}
	return inputs
}

func (oldStep StepModel) getOutputEnvironments() []envmanModels.EnvironmentItemModel {
	outputs := []envmanModels.EnvironmentItemModel{}
	for _, oldOutput := range oldStep.Outputs {
		newOutput := oldOutput.convert()
		outputs = append(outputs, newOutput)
	}
	return outputs
}

func (oldStep StepModel) convert() stepmanModels.StepModel {
	newStep := stepmanModels.StepModel{
		Title:               pointers.NewStringPtr(oldStep.Name),
		Description:         pointers.NewStringPtr(oldStep.Description),
		Website:             pointers.NewStringPtr(oldStep.Website),
		Source:              oldStep.getSource(),
		HostOsTags:          oldStep.HostOsTags,
		ProjectTypeTags:     oldStep.ProjectTypeTags,
		TypeTags:            oldStep.TypeTags,
		IsRequiresAdminUser: pointers.NewBoolPtr(oldStep.IsRequiresAdminUser),
		Inputs:              oldStep.getInputEnvironments(),
		Outputs:             oldStep.getOutputEnvironments(),
	}

	return newStep
}

// ----------------------------
// --- old WorkflowModel -> new StepModel

func (oldWorkflow WorkflowModel) getEnvironments() []envmanModels.EnvironmentItemModel {
	environments := []envmanModels.EnvironmentItemModel{}
	for _, oldEnv := range oldWorkflow.Environments {
		newEnv := oldEnv.convert()
		environments = append(environments, newEnv)
	}
	return environments
}

// Convert ...
func (oldWorkflow WorkflowModel) Convert() bitriseModels.WorkflowModel {
	newWorkflow := bitriseModels.WorkflowModel{
		Environments: oldWorkflow.getEnvironments(),
	}

	stepList := []bitriseModels.StepListItemModel{}
	for _, oldStep := range oldWorkflow.Steps {
		newStep := oldStep.convert()

		_, _, version := oldStep.getStepLibIDVersionData()

		stepIDDataString := "_::" + newStep.Source.Git + "@" + version
		stepListItem := bitriseModels.StepListItemModel{
			stepIDDataString: newStep,
		}
		stepList = append(stepList, stepListItem)
	}
	newWorkflow.Steps = stepList

	return newWorkflow
}

// GetDefaultSteplibSource ...
func GetDefaultSteplibSource(workflow WorkflowModel) string {
	defaultSource := ""
	for _, step := range workflow.Steps {
		if defaultSource == "" {
			defaultSource = step.SteplibSource
		} else if defaultSource != step.SteplibSource {
			return ""
		}
	}
	return defaultSource
}

// ConvertToBitriseDataModel ...
func (oldWorkflow WorkflowModel) ConvertToBitriseDataModel() bitriseModels.BitriseDataModel {
	workflow := oldWorkflow.Convert()

	bitriseData := bitriseModels.BitriseDataModel{
		FormatVersion: "0.9.8",
		Workflows: map[string]bitriseModels.WorkflowModel{
			"target": workflow,
		},
	}

	defaultStepLibSource := GetDefaultSteplibSource(oldWorkflow)
	if defaultStepLibSource != "" {
		bitriseData.DefaultStepLibSource = defaultStepLibSource
	}

	return bitriseData
}
