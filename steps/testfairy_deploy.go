package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldTestfairyDeployStepID ...
	OldTestfairyDeployStepID = "testfairy-deploy"
	// NewTestfairyDeployStepID ...
	NewTestfairyDeployStepID = "testfairy-deploy"
)

//----------------------
// old name: testfairy-deploy
// new name: testfairy-deploy

/*
old version source: https://github.com/bitrise-io/steps-testfairy-deploy.git

inputs:
- TESTFAIRY_API_KEY
- TESTFAIRY_IPA_PATH
- TESTFAIRY_DSYM_PATH
- TESTFAIRY_TESTER_GROUPS
- TESTFAIRY_NOTIFY
- TESTFAIRY_AUTO_UPDATE
- TESTFAIRY_MAX_TEST_DURATION
- TESTFAIRY_VIDEO_RECORDING
- TESTFAIRY_COMMENT
*/

/*
new version source: https://github.com/bitrise-io/steps-testfairy-deploy.git

inputs:
- api_key
- ipa_path
- dsym_path
- tester_groups
- notify
- auto_update
- max_test_duration
- video_recording
- comment
*/

// ConvertTestfairyDeploy ...
func ConvertTestfairyDeploy(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewTestfairyDeployStepID
	inputConversionMap := map[string]string{
		"api_key":           "TESTFAIRY_API_KEY",
		"ipa_path":          "TESTFAIRY_IPA_PATH",
		"dsym_path":         "TESTFAIRY_DSYM_PATH",
		"tester_groups":     "TESTFAIRY_TESTER_GROUPS",
		"notify":            "TESTFAIRY_NOTIFY",
		"auto_update":       "TESTFAIRY_AUTO_UPDATE",
		"max_test_duration": "TESTFAIRY_MAX_TEST_DURATION",
		"video_recording":   "TESTFAIRY_VIDEO_RECORDING",
		"comment":           "TESTFAIRY_COMMENT",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
