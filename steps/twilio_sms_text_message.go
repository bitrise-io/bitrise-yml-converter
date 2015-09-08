package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldTwilioSmsTextMessageStepID ...
	OldTwilioSmsTextMessageStepID = "twilio-sms-text-message"
	// NewTwilioSmsTextMessageStepID ...
	NewTwilioSmsTextMessageStepID = "sms-text-message"
)

//----------------------
// old name: twilio-sms-text-message
// new name: sms-text-message

/*
old version source: https://github.com/bitrise-io/steps-sms-text-message.git

inputs:
- TWILIO_ACCOUNT_SID
- TWILIO_AUTH_TOKEN
- TWILIO_SMS_FROM_NUMBER
- TWILIO_SMS_TO_NUMBER
- TWILIO_SMS_MESSAGE
*/

/*
new version source: https://github.com/bitrise-io/steps-sms-text-message.git

inputs:
- account_sid
- auth_token
- from_number
- to_number
- message
- sms_media
*/

// ConvertTwilioSmsTextMessage ...
func ConvertTwilioSmsTextMessage(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewTwilioSmsTextMessageStepID
	inputConversionMap := map[string]string{
		"account_sid": "TWILIO_ACCOUNT_SID",
		"auth_token":  "TWILIO_AUTH_TOKEN",
		"from_number": "TWILIO_SMS_FROM_NUMBER",
		"to_number":   "TWILIO_SMS_TO_NUMBER",
		"message":     "TWILIO_SMS_MESSAGE",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
