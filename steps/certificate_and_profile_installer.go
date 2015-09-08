package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldCertificateAndProfileInstallerStepID ...
	OldCertificateAndProfileInstallerStepID = "certificate-and-profile-installer"
	// NewCertificateAndProfileInstallerStepID ...
	NewCertificateAndProfileInstallerStepID = "certificate-and-profile-installer"
)

//----------------------
// old name: certificate-and-profile-installer
// new name: certificate-and-profile-installer

/*
old version source: https://github.com/bitrise-io/steps-certificate-and-profile-installer.git

inputs:
- certificate_url
- certificate_passphrase
- provisioning_profile_url
*/

/*
new version source: https://github.com/bitrise-io/steps-certificate-and-profile-installer.git

inputs:
- certificate_url
- certificate_passphrase
- provisioning_profile_url
- keychain_path
- keychain_password
*/

// ConvertCertificateAndProfileInstaller ...
func ConvertCertificateAndProfileInstaller(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewActivateSSHKeyStepID
	inputConversionMap := map[string]string{
		"certificate_url":          "certificate_url",
		"certificate_passphrase":   "certificate_passphrase",
		"provisioning_profile_url": "provisioning_profile_url",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
