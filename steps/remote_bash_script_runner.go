package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldRemoteBashScriptRunnerStepID ...
	OldRemoteBashScriptRunnerStepID = "remote-bash-script-runner"
	// NewRemoteBashScriptRunnerStepID ...
	NewRemoteBashScriptRunnerStepID = "remote-script-runner"
)

//----------------------
// old name: remote-bash-script-runner
// new name: remote-script-runner

/*
old version source: https://github.com/bitrise-io/steps-remote-script-runner.git

inputs:
- DOWNLOAD_SCRIPT_URL
*/

/*
new version source: https://github.com/bitrise-io/steps-remote-script-runner.git

inputs:
- script_url
*/

// ConvertRemoteBashScriptRunner ...
func ConvertRemoteBashScriptRunner(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewRemoteBashScriptRunnerStepID
	inputConversionMap := map[string]string{
		"script_url": "DOWNLOAD_SCRIPT_URL",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
