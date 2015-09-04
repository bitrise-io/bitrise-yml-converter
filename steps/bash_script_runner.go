package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldBashScriptRunnerStepID ...
	OldBashScriptRunnerStepID = "bash-script-runner"
	// NewScriptStepID ...
	NewScriptStepID = "script"
)

//----------------------
// old name: bash-script-runner
// new name: script

/*
old version source: https://github.com/bitrise-io/steps-bash-script.git

inputs:
	- content
  - __INPUT_FILE__
  - BASH_SCRIPT_WORKING_DIR
*/

/*
new version source: https://github.com/bitrise-io/steps-script.git

inputs:
- content
- runner_bin
- is_debug
- working_dir
- script_file_path
*/

// ConvertBashScriptRunner ...
func ConvertBashScriptRunner(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewScriptStepID

	inputConversionMap := map[string]string{}

	_, found, err := utils.GetInputByKey(convertedWorkflowStep.Inputs, "content")
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}
	if found {
		inputConversionMap = map[string]string{
			"content":     "content",
			"working_dir": "BASH_SCRIPT_WORKING_DIR",
		}
	} else {
		inputConversionMap = map[string]string{
			"content":     "__INPUT_FILE__",
			"working_dir": "BASH_SCRIPT_WORKING_DIR",
		}
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
