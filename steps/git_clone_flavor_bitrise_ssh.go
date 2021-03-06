package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldGitCloneFlavorBitriseSSHStepID ...
	OldGitCloneFlavorBitriseSSHStepID = "git-clone_flavor_bitrise_ssh"
)

//----------------------
// old name: git-clone_flavor_bitrise_ssh
// new name: git-clone-extended

/*
old version source: https://github.com/bitrise-io/steps-git-clone.git

inputs:
  - GIT_REPOSITORY_URL
  - BITRISE_GIT_COMMIT
  - BITRISE_GIT_TAG
  - BITRISE_GIT_BRANCH
  - BITRISE_SOURCE_DIR
  - AUTH_USER
  - AUTH_PASSWORD
  - AUTH_SSH_PRIVATE_KEY
  - AUTH_SSH_PRIVATE_KEY_BASE64
  - GIT_CLONE_FORMATTED_OUTPUT_FILE_PATH
  - GIT_CLONE_IS_EXPORT_OUTPUTS
*/

/*
new version source: https://github.com/bitrise-io/steps-git-clone.git

inputs:
- repository_url
- commit
- tag
- branch
- pull_request_id
- clone_into_dir
- auth_ssh_private_key
*/

// ConvertGitCloneFlavorBitriseSSH ...
func ConvertGitCloneFlavorBitriseSSH(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewGitCloneFlavorBitriseStepID
	inputConversionMap := map[string]string{
		"repository_url":       "GIT_REPOSITORY_URL",
		"commit":               "BITRISE_GIT_COMMIT",
		"tag":                  "BITRISE_GIT_TAG",
		"branch":               "BITRISE_GIT_BRANCH",
		"clone_into_dir":       "BITRISE_SOURCE_DIR",
		"auth_ssh_private_key": "AUTH_SSH_PRIVATE_KEY",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
