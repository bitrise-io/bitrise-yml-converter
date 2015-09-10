package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldFastlaneStepID ...
	OldFastlaneStepID = "fastlane"
	// NewFastlaneStepID ...
	NewFastlaneStepID = "fastlane"
)

//----------------------
// old name: fastlane
// new name: fastlane

/*
old version source: https://github.com/bitrise-io/steps-fastlane.git

inputs:
- fastlane_action
- work_dir
*/

/*
new version source: https://github.com/bitrise-io/steps-fastlane.git

inputs:
- lane
- work_dir
*/

// Convertfastlane ...
func Convertfastlane(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewFastlaneStepID
	inputConversionMap := map[string]string{
		"lane":     "fastlane_action",
		"work_dir": "work_dir",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
