package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldRubyScriptWithGemfileStepID ...
	OldRubyScriptWithGemfileStepID = "ruby-script-with-gemfile"
	// NewRubyScriptWithGemfileStepID ...
	NewRubyScriptWithGemfileStepID = "ruby-script"
)

//----------------------
// old name: activate-ssh-key_flavor_bitrise
// new name: activate-ssh-key

/*
old version source: https://github.com/bitrise-io/steps-ruby-script.git

inputs:
- GEMFILE_CONTENT
- __INPUT_FILE__
*/

/*
new version source: https://github.com/bitrise-io/steps-ruby-script.git

inputs:
- gemfile_content
- ruby_content
- script_run_dir
*/

// ConvertRubyScriptWithGemfile ...
func ConvertRubyScriptWithGemfile(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewRubyScriptWithGemfileStepID
	inputConversionMap := map[string]string{
		"gemfile_content": "GEMFILE_CONTENT",
		"ruby_content":    "__INPUT_FILE__",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
