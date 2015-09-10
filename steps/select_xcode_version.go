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
- SELECT_XCODE_VERSION_CHANNEL_ID
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
		"version_channel_id": "SELECT_XCODE_VERSION_CHANNEL_ID",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
