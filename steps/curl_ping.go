package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldCurlPingStepID ...
	OldCurlPingStepID = "curl-ping"
	// NewCurlPingStepID ...
	NewCurlPingStepID = "curl-ping"
)

//----------------------
// old name: curl-ping
// new name: curl-ping

/*
old version source: https://github.com/bitrise-io/steps-curl-ping.git

inputs:
  - PING_URL
  - OPTIONAL_CURL_PARAMS
*/

/*
new version source: https://github.com/bitrise-io/steps-curl-ping.git

inputs:
- ping_url
- curl_params
*/

// ConvertCurlPing ...
func ConvertCurlPing(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewCurlPingStepID
	inputConversionMap := map[string]string{
		"ping_url":    "PING_URL",
		"curl_params": "OPTIONAL_CURL_PARAMS",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
