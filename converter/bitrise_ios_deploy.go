package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: bitrise-ios-deploy
// new name: bitrise-ios-deploy

/*
old version source: https://github.com/bitrise-io/steps-xcode-builder.git

inputs:
  - STEPLIB_FORMATTED_OUTPUT_FILE_PATH
  - STEP_BITRISE_IOS_DEPLOY_BUILD_URL
  - STEP_BITRISE_IOS_DEPLOY_API_TOKEN
  - STEP_BITRISE_IOS_DEPLOY_IPA_PATH
  - STEP_BITRISE_IOS_DEPLOY_NOTIFY_USER_GROUPS
  - STEP_BITRISE_IOS_DEPLOY_NOTIFY_EMAILS
  - STEP_BITRISE_IOS_DEPLOY_ENABLE_PUBLIC_PAGE
*/

/*
new version source: https://github.com/bitrise-io/steps-bitrise-ios-deploy.git

inputs:
inputs:
- build_url
- build_api_token
- ipa_path
- notify_user_groups
- notify_email_list
- is_enable_public_page
*/

func convertBitriseIosDeploy(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	stepListItems, err := certificateStep()
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	// Convert Bitrise iOS App Deploy step
	newStepID := NewXcodeArchiveStepID
	inputConversionMap := map[string]string{
		"build_url":             "STEP_BITRISE_IOS_DEPLOY_BUILD_URL",
		"build_api_token":       "STEP_BITRISE_IOS_DEPLOY_API_TOKEN",
		"ipa_path":              "STEP_BITRISE_IOS_DEPLOY_IPA_PATH",
		"notify_user_groups":    "STEP_BITRISE_IOS_DEPLOY_NOTIFY_USER_GROUPS",
		"notify_email_list":     "STEP_BITRISE_IOS_DEPLOY_NOTIFY_EMAILS",
		"is_enable_public_page": "STEP_BITRISE_IOS_DEPLOY_ENABLE_PUBLIC_PAGE",
	}

	newStep, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := BitriseVerifiedStepLibGitURI + "::" + newStepID
	stepListItems = append(stepListItems, bitriseModels.StepListItemModel{stepIDDataString: newStep})

	return stepListItems, nil
}
