package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: generic-script-runner
// new name: script

/*
old version source: https://github.com/bitrise-io/steps-generic-script-runner.git

inputs:
  - GENERIC_SCRIPT_RUNNER_CONTENT
  - GENERIC_SCRIPT_RUNNER_BIN
  - GENERIC_SCRIPT_RUNNER_WORKING_DIR
  - GENERIC_SCRIPT_RUNNER_SCRIPT_TMP_PATH
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

func convertGenericScriptRunner(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := newScriptStepID
	inputConversionMap := map[string]string{
		"content":          "GENERIC_SCRIPT_RUNNER_CONTENT",
		"runner_bin":       "GENERIC_SCRIPT_RUNNER_BIN",
		"working_dir":      "GENERIC_SCRIPT_RUNNER_WORKING_DIR",
		"script_file_path": "GENERIC_SCRIPT_RUNNER_SCRIPT_TMP_PATH",
	}

	return convertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
