package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldSelectXcodeVersionStepID ...
	OldSelectXcodeVersionStepID = "select-xcode-version"
	// NewSelectXcodeVersionStepID ...
	NewSelectXcodeVersionStepID = "select-xcode-version"
)

//----------------------
// old name: select-xcode-version
// new name: select-xcode-version

/*
old version source: https://github.com/bitrise-io/steps-select-xcode-version.git

inputs:
- SELECT_XCODE_VERSION_CHANNEL_IDó
*/

/*
new version source: https://github.com/bitrise-io/steps-select-xcode-version.git

inputs:
- version_channel_id
*/

// ConvertSelectXcodeVersion ...
func ConvertSelectXcodeVersion(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewSelectXcodeVersionStepID
	inputConversionMap := map[string]string{
		"ssh_rsa_private_key":        "SSH_RSA_PRIVATE_KEY",
		"ssh_key_save_path":          "SSH_KEY_SAVE_PATH",
		"is_remove_other_identities": "IS_REMOVE_OTHER_IDENTITIES",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
