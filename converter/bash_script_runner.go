package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: bash-script-runner
// new name: script

/*
old version source: https://github.com/bitrise-io/steps-bash-script.git

inputs:
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

func convertBashScriptRunner(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewScriptStepID
	inputConversionMap := map[string]string{
		"content":     "__INPUT_FILE__",
		"working_dir": "BASH_SCRIPT_WORKING_DIR",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
