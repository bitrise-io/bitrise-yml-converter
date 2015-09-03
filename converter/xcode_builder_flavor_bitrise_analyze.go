package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: xcode-builder_flavor_bitrise_analyze
// new name: steps-activate-ssh-key

/*
old version source: https://github.com/bitrise-io/steps-activate-ssh-key.git

inputs:
  - XCODE_BUILDER_PROJECT_ROOT_DIR_PATH
  - XCODE_BUILDER_PROJECT_PATH
  - XCODE_BUILDER_SCHEME
	- XCODE_BUILDER_ACTION
	- XCODE_BUILDER_CERTIFICATE_URL
	- XCODE_BUILDER_CERTIFICATE_PASSPHRASE
	- XCODE_BUILDER_PROVISION_URL
	- XCODE_BUILDER_BUILD_TOOL
	- XCODE_BUILDER_CERTIFICATES_DIR
*/

/*
new version source: https://github.com/bitrise-io/steps-activate-ssh-key.git

inputs:
- workdir
- project_path
- scheme
- is_clean_build
*/

func convertXcodeBuilderFlavorBitriseAnalyze(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := newXcodeAnalyzeStepID
	inputConversionMap := map[string]string{
		"project_path": "XCODE_BUILDER_PROJECT_PATH",
		"scheme":       "XCODE_BUILDER_SCHEME",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
