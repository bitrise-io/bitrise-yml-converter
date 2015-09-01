package converter

import (
	"fmt"

	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// BitriseVerifiedStepLibGitURI ...
	BitriseVerifiedStepLibGitURI = "https://github.com/bitrise-io/bitrise-steplib.git"
)

func commonMergeStep(newWorkflowStep stepmanModels.StepModel, newSpecStep stepmanModels.StepModel) stepmanModels.StepModel {
	if newWorkflowStep.Title != nil && *newWorkflowStep.Title != "" {
		newSpecStep.Title = pointers.NewStringPtr(*newWorkflowStep.Title)
	}
	return newSpecStep
}

func getInputByKey(newStep stepmanModels.StepModel, key string) (envmanModels.EnvironmentItemModel, error) {
	for _, input := range newStep.Inputs {
		aKey, _, err := input.GetKeyValuePair()
		if err != nil {
			return envmanModels.EnvironmentItemModel{}, err
		}
		if aKey == key {
			return input, nil
		}
	}
	return envmanModels.EnvironmentItemModel{}, fmt.Errorf("No Environmnet found for key (%s)", key)
}

func convertSlack(convertedWorkflowStep stepmanModels.StepModel) (stepmanModels.StepModel, error) {
	stepID := "slack"
	specStep, err := GetStepFromNewSteplib(stepID, BitriseVerifiedStepLibGitURI)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	mergedStep := commonMergeStep(convertedWorkflowStep, specStep)

	inputConversionMap := map[string]string{
		"webhook_url":            "SLACK_WEBHOOK_URL",
		"channel":                "SLACK_CHANNEL",
		"from_username":          "SLACK_FROM_NAME",
		"from_username_on_error": "SLACK_ERROR_FROM_NAME",
		"message":                "SLACK_MESSAGE_TEXT",
		"message_on_error":       "SLACK_ERROR_MESSAGE_TEXT",
		"emoji":                  "SLACK_ICON_EMOJI",
		"emoji_on_error":         "SLACK_ERROR_ICON_EMOJI",
		"icon_url":               "SLACK_ICON_URL",
		"icon_url_on_error":      "SLACK_ERROR_ICON_URL",
	}

	mergedStepInputs := []envmanModels.EnvironmentItemModel{}
	for _, specInput := range mergedStep.Inputs {
		specKey, _, err := specInput.GetKeyValuePair()
		if err != nil {
			return stepmanModels.StepModel{}, err
		}

		workflowInputKey, found := inputConversionMap[specKey]
		if found == false {
			mergedStepInputs = append(mergedStepInputs, specInput)
			continue
		}

		workflowInput, err := getInputByKey(convertedWorkflowStep, workflowInputKey)
		if err != nil {
			return stepmanModels.StepModel{}, err
		}

		_, workflowValue, err := workflowInput.GetKeyValuePair()
		if err != nil {
			return stepmanModels.StepModel{}, err
		}

		workflowOptions, err := workflowInput.GetOptions()
		if err != nil {
			return stepmanModels.StepModel{}, err
		}

		mergedInput := envmanModels.EnvironmentItemModel{
			specKey:                 workflowValue,
			envmanModels.OptionsKey: workflowOptions,
		}

		mergedStepInputs = append(mergedStepInputs, mergedInput)
	}
	mergedStep.Inputs = mergedStepInputs

	return mergedStep, nil
}
