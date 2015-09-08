package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldParsePushNotificationStepID ...
	OldParsePushNotificationStepID = "parse-push-notification"
	// NewParsePushNotificationStepID ...
	NewParsePushNotificationStepID = "push-notification-with-parse"
)

//----------------------
// old name: parse-push-notification
// new name: push-notification-with-parse

/*
old version source: https://github.com/bitrise-io/steps-push-notification-with-parse.git

inputs:
- PARSE_PUSH_APP_ID
- PARSE_PUSH_REST_KEY
- PARSE_PUSH_MESSAGE
*/

/*
new version source: https://github.com/bitrise-io/steps-push-notification-with-parse.git

inputs:
- app_id
- rest_key
- message
*/

// ConvertParsePushNotification ...
func ConvertParsePushNotification(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewParsePushNotificationStepID
	inputConversionMap := map[string]string{
		"app_id":   "PARSE_PUSH_APP_ID",
		"rest_key": "PARSE_PUSH_REST_KEY",
		"message":  "PARSE_PUSH_MESSAGE",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
