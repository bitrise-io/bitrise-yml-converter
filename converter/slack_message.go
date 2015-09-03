package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: slack-message
// new name: slack

/*
old version source: https://github.com/bitrise-io/steps-slack-message.git

inputs:
  - SLACK_WEBHOOK_URL
  - SLACK_CHANNEL
  - SLACK_FROM_NAME
  - SLACK_ERROR_FROM_NAME
  - SLACK_MESSAGE_TEXT
  - SLACK_ERROR_MESSAGE_TEXT
  - SLACK_ICON_EMOJI
  - SLACK_ERROR_ICON_EMOJI
  - SLACK_ICON_URL
  - SLACK_ERROR_ICON_URL
*/

/*
new version source: https://github.com/bitrise-io/steps-slack-message.git

inputs:
- content
inputs:
- webhook_url
- channel
- from_username
- from_username_on_error
- message
- message_on_error
- emoji
- emoji_on_error
- icon_url
- icon_url_on_error
*/

func convertSlackMessage(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
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
