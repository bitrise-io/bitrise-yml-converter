package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldFabricCrashlyticsBetaDeployStepID ...
	OldFabricCrashlyticsBetaDeployStepID = "fabric-crashlytics-beta-deploy"
	// NewFabricCrashlyticsBetaDeployStepID ...
	NewFabricCrashlyticsBetaDeployStepID = "fabric-crashlytics-beta-deploy"
)

//----------------------
// old name: fabric-crashlytics-beta-deploy
// new name: fabric-crashlytics-beta-deploy

/*
old version source: https://github.com/bitrise-io/steps-fabric-crashlytics-beta-deploy.git

inputs:
- STEP_CRASHLYTICS_API_KEY
- STEP_CRASHLYTICS_BUILD_SECRET
- STEP_CRASHLYTICS_IPA_PATH
- STEP_CRASHLYTICS_EMAIL_LIST
- STEP_CRASHLYTICS_GROUP_ALIASES_LIST
- STEP_CRASHLYTICS_NOTIFICATION
- STEP_CRASHLYTICS_RELEASE_NOTES
- STEP_CERT_ACTIVATOR_CERTIFICATE_URL
- STEP_CERT_ACTIVATOR_CERTIFICATE_PASSPHRASE
- STEP_CERT_ACTIVATOR_CERTIFICATES_DIR
- STEP_CERT_ACTIVATOR_KEYCHAIN_PATH
- STEP_CERT_ACTIVATOR_KEYCHAIN_PSW
*/

/*
new version source: https://github.com/bitrise-io/steps-fabric-crashlytics-beta-deploy.git

inputs:
- api_key
- build_secret
- ipa_path
- email_list
- group_aliases_list
- notification
- release_notes
- STEP_CERT_ACTIVATOR_CERTIFICATE_URL
- STEP_CERT_ACTIVATOR_CERTIFICATE_PASSPHRASE
- STEP_CERT_ACTIVATOR_CERTIFICATES_DIR
- STEP_CERT_ACTIVATOR_KEYCHAIN_PATH
- STEP_CERT_ACTIVATOR_KEYCHAIN_PSW
*/

// ConvertFabricCrashlyticsBetaDeploy ...
func ConvertFabricCrashlyticsBetaDeploy(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewFabricCrashlyticsBetaDeployStepID
	inputConversionMap := map[string]string{
		"api_key":            "STEP_CRASHLYTICS_API_KEY",
		"build_secret":       "STEP_CRASHLYTICS_BUILD_SECRET",
		"ipa_path":           "STEP_CRASHLYTICS_IPA_PATH",
		"email_list":         "STEP_CRASHLYTICS_EMAIL_LIST",
		"group_aliases_list": "STEP_CRASHLYTICS_GROUP_ALIASES_LIST",
		"notification":       "STEP_CRASHLYTICS_NOTIFICATION",
		"release_notes":      "STEP_CRASHLYTICS_RELEASE_NOTES",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
