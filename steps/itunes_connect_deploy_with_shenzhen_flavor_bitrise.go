package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldItunesConnectDeployWithShenzhenFlavorBitriseStepID ...
	OldItunesConnectDeployWithShenzhenFlavorBitriseStepID = "itunes-connect-deploy-with-shenzhen_flavor_bitrise"
	// NewItunesConnectDeployWithShenzhenFlavorBitriseStepID ...
	NewItunesConnectDeployWithShenzhenFlavorBitriseStepID = "deploy-to-itunesconnect-shenzhen"
)

//----------------------
// old name: itunes-connect-deploy-with-shenzhen_flavor_bitrise
// new name: deploy-to-itunesconnect-shenzhen

/*
old version source: https://github.com/bitrise-io/steps-deploy-to-itunesconnect-shenzhen.git

inputs:
- STEP_SHENZHEN_DEPLOY_IPA_PATH
- STEP_SHENZHEN_DEPLOY_ITUNESCON_USER
- STEP_SHENZHEN_DEPLOY_ITUNESCON_PASSWORD
- STEP_SHENZHEN_DEPLOY_ITUNESCON_APP_ID
*/

/*
new version source: https://github.com/bitrise-io/steps-deploy-to-itunesconnect-shenzhen.git

inputs:
- ipa_path
- itunescon_user
- password
- app_id
*/

// ConvertItunesConnectDeployWithShenzhenFlavorBitrise ...
func ConvertItunesConnectDeployWithShenzhenFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewItunesConnectDeployWithShenzhenFlavorBitriseStepID
	inputConversionMap := map[string]string{
		"ipa_path":       "STEP_SHENZHEN_DEPLOY_IPA_PATH",
		"itunescon_user": "STEP_SHENZHEN_DEPLOY_ITUNESCON_USER",
		"password":       "STEP_SHENZHEN_DEPLOY_ITUNESCON_PASSWORD",
		"app_id":         "STEP_SHENZHEN_DEPLOY_ITUNESCON_APP_ID",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
