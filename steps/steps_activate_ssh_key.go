package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldActivateSSHKeyFlavorBitriseStepID ...
	OldActivateSSHKeyFlavorBitriseStepID = "activate-ssh-key_flavor_bitrise"
	// NewActivateSSHKeyStepID ...
	NewActivateSSHKeyStepID = "activate-ssh-key"
)

//----------------------
// old name: steps-activate-ssh-key
// new name: activate-ssh-key

/*
old version source: https://github.com/bitrise-io/steps-activate-ssh-key.git

inputs:
- SSH_RSA_PRIVATE_KEY
- SSH_KEY_SAVE_PATH
- IS_REMOVE_OTHER_IDENTITIES
*/

/*
new version source: https://github.com/bitrise-io/steps-activate-ssh-key.git

inputs:
- ssh_rsa_private_key
- ssh_key_save_path
- is_remove_other_identities
*/

// ConvertActivteSSHKey ...
func ConvertActivteSSHKey(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewActivateSSHKeyStepID
	inputConversionMap := map[string]string{
		"ssh_rsa_private_key":        "SSH_RSA_PRIVATE_KEY",
		"ssh_key_save_path":          "SSH_KEY_SAVE_PATH",
		"is_remove_other_identities": "IS_REMOVE_OTHER_IDENTITIES",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
