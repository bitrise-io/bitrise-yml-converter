package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldHipchatStepID ...
	OldHipchatStepID = "hipchat"
	// NewHipchatStepID ...
	NewHipchatStepID = "hipchat"
)

//----------------------
// old name: hipchat
// new name: hipchat

/*
old version source: https://github.com/bitrise-io/steps-hipchat.git

inputs:
  - HIPCHAT_TOKEN
  - HIPCHAT_ROOMID
  - HIPCHAT_FROMNAME
  - HIPCHAT_ERROR_FROMNAME
  - HIPCHAT_MESSAGE
  - HIPCHAT_ERROR_MESSAGE
  - HIPCHAT_MESSAGE_COLOR
  - HIPCHAT_ERROR_MESSAGE_COLOR
*/

/*
new version source: https://github.com/bitrise-io/steps-hipchat.git

inputs:
inputs:
- auth_token
- room_id
- from_name
- from_name_on_error
- message
- message_on_error
- color
- color_on_error
*/

// ConvertHipchat ...
func ConvertHipchat(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
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

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
