package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldItunesConnectDeployWithDeliverFlavorBitriseStepID ...
	OldItunesConnectDeployWithDeliverFlavorBitriseStepID = "itunes-connect-deploy-with-deliver_flavor_bitrise"
	// NewItunesConnectDeployWithDeliverFlavorBitriseStepID ...
	NewItunesConnectDeployWithDeliverFlavorBitriseStepID = "deploy-to-itunesconnect-deliver"
)

//----------------------
// old name: itunes-connect-deploy-with-deliver_flavor_bitrise
// new name: deploy-to-itunesconnect-deliver

/*
old version source: https://github.com/bitrise-io/steps-deploy-to-itunesconnect-deliver.git

inputs:
- STEP_DELIVER_DEPLOY_IPA_PATH
- STEP_DELIVER_DEPLOY_ITUNESCON_USER
- STEP_DELIVER_DEPLOY_ITUNESCON_PASSWORD
- STEP_DELIVER_DEPLOY_ITUNESCON_APP_ID
- STEP_DELIVER_DEPLOY_IS_SUBMIT_FOR_BETA
*/

/*
new version source: https://github.com/bitrise-io/steps-deploy-to-itunesconnect-deliver.git

inputs:
- ipa_path
- itunescon_user
- password
- app_id
- submit_for_beta
*/

// ConvertItunesConnectDeployWithDeliverFlavorBitrise ...
func ConvertItunesConnectDeployWithDeliverFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewItunesConnectDeployWithDeliverFlavorBitriseStepID
	inputConversionMap := map[string]string{
		"ipa_path":        "STEP_DELIVER_DEPLOY_IPA_PATH",
		"itunescon_user":  "STEP_DELIVER_DEPLOY_ITUNESCON_USER",
		"password":        "STEP_DELIVER_DEPLOY_ITUNESCON_PASSWORD",
		"app_id":          "STEP_DELIVER_DEPLOY_ITUNESCON_APP_ID",
		"submit_for_beta": "STEP_DELIVER_DEPLOY_IS_SUBMIT_FOR_BETA",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
