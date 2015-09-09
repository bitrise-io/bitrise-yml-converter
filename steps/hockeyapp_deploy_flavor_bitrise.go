package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldHockeyappDeployFlavorBitriseStepID ...
	OldHockeyappDeployFlavorBitriseStepID = "hockeyapp-deploy_flavor_bitrise"
	// NewHockeyappDeployFlavorBitriseStepID ...
	NewHockeyappDeployFlavorBitriseStepID = "hockeyapp-deploy"
)

//----------------------
// old name: hockeyapp-deploy_flavor_bitrise
// new name: hockeyapp-deploy

/*
old version source: https://github.com/bitrise-io/step-hockeyapp-deploy.git

inputs:
- BITRISE_IPA_PATH
- BITRISE_DSYM_PATH
- HOCKEYAPP_TOKEN
- HOCKEYAPP_APP_ID
- HOCKEYAPP_NOTES
- HOCKEYAPP_NOTES_TYPE
- HOCKEYAPP_NOTIFY
- HOCKEYAPP_STATUS
- HOCKEYAPP_MANDATORY
- HOCKEYAPP_TAGS
- HOCKEYAPP_COMMIT_SHA
- HOCKEYAPP_BUILD_SERVER_URL
- HOCKEYAPP_REPOSITORY_URL
*/

/*
new version source: https://github.com/bitrise-io/steps-hockeyapp-deploy.git

inputs:
- ipa_path
- dsym_path
- api_token
- app_id
- notes
- notes_type
- notify
- status
- mandatory
- tags
- commit_sha
- build_server_url
- repository_url
*/

// ConvertHockeyappDeployFlavorBitrise ...
func ConvertHockeyappDeployFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewHockeyappDeployFlavorBitriseStepID
	inputConversionMap := map[string]string{
		"ipa_path":         "BITRISE_IPA_PATH",
		"dsym_path":        "BITRISE_DSYM_PATH",
		"api_token":        "HOCKEYAPP_TOKEN",
		"app_id":           "HOCKEYAPP_APP_ID",
		"notes":            "HOCKEYAPP_NOTES",
		"notes_type":       "HOCKEYAPP_NOTES_TYPE",
		"notify":           "HOCKEYAPP_NOTIFY",
		"status":           "HOCKEYAPP_STATUS",
		"mandatory":        "HOCKEYAPP_MANDATORY",
		"tags":             "HOCKEYAPP_TAGS",
		"commit_sha":       "HOCKEYAPP_COMMIT_SHA",
		"build_server_url": "HOCKEYAPP_BUILD_SERVER_URL",
		"repository_url":   "HOCKEYAPP_REPOSITORY_URL",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
