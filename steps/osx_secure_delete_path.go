package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldOsxSecureDeletePathStepID ...
	OldOsxSecureDeletePathStepID = "osx-secure-delete-path"
	// NewOsxSecureDeletePathStepID ...
	NewOsxSecureDeletePathStepID = "secure-delete-path"
)

//----------------------
// old name: osx-secure-delete-path
// new name: secure-delete-path

/*
old version source: https://github.com/bitrise-io/steps-secure-delete-path.git

inputs:
- SECURE_DELETE_PATH
- SECURE_DELETE_WITHSUDO
*/

/*
new version source: https://github.com/bitrise-io/steps-secure-delete-path.git

inputs:
- path
- with_sudo
*/

// ConvertOsxSecureDeletePath ...
func ConvertOsxSecureDeletePath(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewAmazonS3UploaderStepID
	inputConversionMap := map[string]string{
		"path":      "SECURE_DELETE_PATH",
		"with_sudo": "SECURE_DELETE_WITHSUDO",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
