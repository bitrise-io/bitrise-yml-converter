package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: cocoapods-and-xcode-repository-validator_flavor_bitrise
// new name: cocoapods-and-repository-validator

/*
old version source: https://github.com/bitrise-io/steps-cocoapods-and-repository-validator.git

inputs:
- BITRISE_SOURCE_DIR
- GATHER_PROJECTS
- IS_UPDATE_COCOAPODS
- REPO_VALIDATOR_SINGLE_BRANCH
*/

/*
new version source: https://github.com/bitrise-io/steps-cocoapods-and-repository-validator.git

inputs:
- source_root_path
- is_update_cocoapods
- scan_only_branch
*/

func convertCocoapodsAndXcodeRepositoryValidatorFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := newCocoapodsAndRepositoryValidatorStepID
	inputConversionMap := map[string]string{
		"source_root_path":    "BITRISE_SOURCE_DIR",
		"is_update_cocoapods": "IS_UPDATE_COCOAPODS",
		"scan_only_branch":    "REPO_VALIDATOR_SINGLE_BRANCH",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
