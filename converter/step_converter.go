package converter

import (
	"fmt"

	bitriseModels "github.com/bitrise-io/bitrise/models"
	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// BitriseVerifiedStepLibGitURI ...
	BitriseVerifiedStepLibGitURI = "https://github.com/bitrise-io/bitrise-steplib.git"
)

//----------------------
// Common methods

func getInputByKey(inputs []envmanModels.EnvironmentItemModel, key string) (envmanModels.EnvironmentItemModel, error) {
	for _, input := range inputs {
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

func convertStepsInputs(originalInputs, diffInputs []envmanModels.EnvironmentItemModel, conversionMap map[string]string) ([]envmanModels.EnvironmentItemModel, error) {
	mergedStepInputs := []envmanModels.EnvironmentItemModel{}
	for _, specInput := range originalInputs {
		specKey, _, err := specInput.GetKeyValuePair()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		workflowInputKey, found := conversionMap[specKey]
		if found == false {
			mergedStepInputs = append(mergedStepInputs, specInput)
			continue
		}

		workflowInput, err := getInputByKey(diffInputs, workflowInputKey)
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		_, workflowValue, err := workflowInput.GetKeyValuePair()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		if workflowValue == "" {
			continue
		}

		workflowOptions, err := workflowInput.GetOptions()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		workflowOptions.Title = nil
		workflowOptions.Description = nil
		workflowOptions.Summary = nil
		workflowOptions.ValueOptions = []string{}
		workflowOptions.IsRequired = nil
		workflowOptions.IsDontChangeValue = nil
		// workflowOptions.IsExpand should be keep

		mergedInput := envmanModels.EnvironmentItemModel{
			specKey:                 workflowValue,
			envmanModels.OptionsKey: workflowOptions,
		}

		mergedStepInputs = append(mergedStepInputs, mergedInput)
	}
	return mergedStepInputs, nil
}

func convertStep(convertedWorkflowStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) (stepmanModels.StepModel, error) {
	// The new StepLib version of step
	specStep, err := GetStepFromNewSteplib(newStepID, BitriseVerifiedStepLibGitURI)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	if convertedWorkflowStep.Title != nil && *convertedWorkflowStep.Title != "" {
		specStep.Title = pointers.NewStringPtr(*convertedWorkflowStep.Title)
	}
	if convertedWorkflowStep.IsAlwaysRun != nil {
		specStep.IsAlwaysRun = pointers.NewBoolPtr(*convertedWorkflowStep.IsAlwaysRun)
	}

	// Merge new StepLib version inputs, with old workflow defined
	mergedInputs, err := convertStepsInputs(specStep.Inputs, convertedWorkflowStep.Inputs, inputConversionMap)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	specStep.Inputs = mergedInputs

	return specStep, nil
}

func convertStepAndCreateStepListItem(convertedWorkflowStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) ([]bitriseModels.StepListItemModel, error) {
	newStep, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := BitriseVerifiedStepLibGitURI + "::" + newStepID

	stepListItem := bitriseModels.StepListItemModel{
		stepIDDataString: newStep,
	}

	return []bitriseModels.StepListItemModel{stepListItem}, nil
}

//----------------------
// slack

func convertSlack(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewSlackStepID
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

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}

//----------------------
// hipchat

func convertHipchat(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewHipchatStepID
	inputConversionMap := map[string]string{
		"auth_token":         "HIPCHAT_TOKEN",
		"room_id":            "HIPCHAT_ROOMID",
		"from_name":          "HIPCHAT_FROMNAME",
		"from_name_on_error": "HIPCHAT_ERROR_FROMNAME",
		"message":            "HIPCHAT_MESSAGE",
		"message_on_error":   "HIPCHAT_ERROR_MESSAGE",
		"color":              "HIPCHAT_MESSAGE_COLOR",
		"color_on_error":     "HIPCHAT_ERROR_MESSAGE_COLOR",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}

//----------------------
// script

func converScript(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewScriptStepID
	inputConversionMap := map[string]string{
		"content":          "GENERIC_SCRIPT_RUNNER_CONTENT",
		"runner_bin":       "GENERIC_SCRIPT_RUNNER_BIN",
		"working_dir":      "GENERIC_SCRIPT_RUNNER_WORKING_DIR",
		"script_file_path": "GENERIC_SCRIPT_RUNNER_SCRIPT_TMP_PATH",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}

//----------------------
// xcode-archive

func converXcodeArchive(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	// Cerificate step separated in new StepLib
	// Step (https://github.com/bitrise-io/steps-certificate-and-profile-installer.git)
	// need to insert befor Xcode-Archive
	certificateStepGitURI := "https://github.com/bitrise-io/steps-certificate-and-profile-installer.git"
	certificateStepTitle := "steps-certificate-and-profile-installer"

	certificateStep, err := GetStepFromGit(certificateStepGitURI)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}
	certificateStep.RunIf = pointers.NewStringPtr(".IsCI")
	certificateStep.Title = pointers.NewStringPtr(certificateStepTitle)

	stepIDDataString := "git::" + certificateStepGitURI + "@master"
	stepListItems := []bitriseModels.StepListItemModel{
		bitriseModels.StepListItemModel{
			stepIDDataString: certificateStep,
		},
	}

	// Convert Xcode-Archive step
	newStepID := NewXcodeArchiveStepID
	inputConversionMap := map[string]string{
		"project_path": "XCODE_BUILDER_PROJECT_PATH",
		"scheme":       "XCODE_BUILDER_SCHEME",
		// "project_path": "",
		// "output_dir": "",
	}

	newStep, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString = BitriseVerifiedStepLibGitURI + "::" + newStepID
	stepListItems = append(stepListItems, bitriseModels.StepListItemModel{stepIDDataString: newStep})

	return stepListItems, nil
}
