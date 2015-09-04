package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
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

func convertCocoapodsFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := newCocospodsInstallStepID
	inputConversionMap := map[string]string{
		"source_root_path":    "BITRISE_SOURCE_DIR",
		"is_update_cocoapods": "IS_UPDATE_COCOAPODS",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
