package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldMailgunEmailStepID ...
	OldMailgunEmailStepID = "mailgun-email"
	// NewMailgunEmailStepID ...
	NewMailgunEmailStepID = "email-with-mailgun"
)

//----------------------
// old name: mailgun-email
// new name: email-with-mailgun

/*
old version source: https://github.com/bitrise-io/steps-email-with-mailgun.git

inputs:
- MAILGUN_API_KEY
- MAILGUN_DOMAIN
- MAILGUN_SEND_TO
- MAILGUN_EMAIL_SUBJECT
- MAILGUN_EMAIL_MESSAGE
*/

/*
new version source: https://github.com/bitrise-io/steps-email-with-mailgun.git

inputs:
- api_key
- domain
- send_to
- subject
- message
*/

// ConvertMailgunEmail ...
func ConvertMailgunEmail(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewMailgunEmailStepID
	inputConversionMap := map[string]string{
		"api_key": "MAILGUN_API_KEY",
		"domain":  "MAILGUN_DOMAIN",
		"send_to": "MAILGUN_SEND_TO",
		"subject": "MAILGUN_EMAIL_SUBJECT",
		"message": "MAILGUN_EMAIL_MESSAGE",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
