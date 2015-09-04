package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldCocoapodsFlavorBitriseStepID ..
	OldCocoapodsFlavorBitriseStepID = "cocoapods_flavor_bitrise"
	// NewCocospodsInstallStepID ...
	NewCocospodsInstallStepID = "cocoapods-install"
)

//----------------------
// old name: cocoapods_flavor_bitrise
// new name: cocoapods-install

/*
old version source: https://github.com/bitrise-io/steps-cocoapods-and-repository-validator.git

inputs:
  - BITRISE_SOURCE_DIR
  - GATHER_PROJECTS
  - IS_UPDATE_COCOAPODS
  - REPO_VALIDATOR_SINGLE_BRANCH
*/

/*
new version source: https://github.com/bitrise-io/steps-cocoapods-install.git

inputs:
- source_root_path
- is_update_cocoapods
*/

// ConvertCocoapodsFlavorBitrise ...
func ConvertCocoapodsFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewCocospodsInstallStepID
	inputConversionMap := map[string]string{
		"source_root_path":    "BITRISE_SOURCE_DIR",
		"is_update_cocoapods": "IS_UPDATE_COCOAPODS",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
