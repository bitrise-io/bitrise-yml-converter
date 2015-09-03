package converter

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	oldmodels "github.com/bitrise-io/bitrise-yml-converter/old_models"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	oldBitriseIosDeployStepID = "bitrise-ios-deploy"
	newBitriseIosDeployStepID = "bitrise-ios-deploy"

	oldHipchatStepID = "hipchat"
	newHipchatStepID = "hipchat"

	oldSlackMessageStepID = "slack-message"
	newSlackStepID        = "slack"

	oldGenericScriptRunnerStepID = "generic-script-runner"
	oldBashScriptRunnerStepID    = "bash-script-runner"
	newScriptStepID              = "script"

	olXcodeBuilderFlavorBitriseCreateArchiveStepID = "xcode-builder_flavor_bitrise_create-archive"
	newXcodeArchiveStepID                          = "xcode-archive"

	oldXcodeBuilderFlavorBitriseUnittestStepID = "xcode-builder_flavor_bitrise_unittest"
	newXcodeTest                               = "xcode-test"

	oldActivateSSHKeyFlavorBitriseStepID = "activate-ssh-key_flavor_bitrise"
	newActivateSSHKeyStepID              = ""

	oldXcodeBuilderFlavorBitriseAnalyzeStepID = "xcode-builder_flavor_bitrise_analyze"
	newXcodeAnalyzeStepID                     = ""
)

type stepConverter func(stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error)

// New step ID <-> Converter function
func getStepConverterFunctionMap() map[string]stepConverter {
	return map[string]stepConverter{
		oldHipchatStepID:                               convertHipchat,
		oldSlackMessageStepID:                          convertSlackMessage,
		oldGenericScriptRunnerStepID:                   convertGenericScriptRunner,
		oldBashScriptRunnerStepID:                      convertBashScriptRunner,
		olXcodeBuilderFlavorBitriseCreateArchiveStepID: convertXcodeBuilderFlavorBitriseCreateArchive,
		oldXcodeBuilderFlavorBitriseUnittestStepID:     convertXcodeBuilderFlavorBitriseUnittest,
		oldBitriseIosDeployStepID:                      convertBitriseIosDeploy,
		oldActivateSSHKeyFlavorBitriseStepID:           convertActivteSSHKey,
		oldXcodeBuilderFlavorBitriseAnalyzeStepID:      convertXcodeBuilderFlavorBitriseAnalyze,
	}
}

// old step ID <-> New step ID
func getStepConversionMap() map[string]string {
	return map[string]string{
		oldHipchatStepID:                               newHipchatStepID,
		oldSlackMessageStepID:                          newSlackStepID,
		oldGenericScriptRunnerStepID:                   newScriptStepID,
		oldBashScriptRunnerStepID:                      newScriptStepID,
		olXcodeBuilderFlavorBitriseCreateArchiveStepID: newXcodeArchiveStepID,
		oldXcodeBuilderFlavorBitriseUnittestStepID:     newXcodeTest,
		oldBitriseIosDeployStepID:                      newBitriseIosDeployStepID,
		oldActivateSSHKeyFlavorBitriseStepID:           newActivateSSHKeyStepID,
		oldXcodeBuilderFlavorBitriseAnalyzeStepID:      newXcodeAnalyzeStepID,
	}
}

func getNewStepIDAndConverter(oldStepID string) (string, stepConverter, bool) {
	stepConversionMap := getStepConversionMap()
	newID, found := stepConversionMap[oldStepID]
	if !found {
		return "", nil, false
	}

	converterFunctionMap := getStepConverterFunctionMap()
	converter, found := converterFunctionMap[oldStepID]
	if !found {
		return "", nil, false
	}
	return newID, converter, true
}

// GetDefaultSteplibSource ...
func GetDefaultSteplibSource(workflow oldmodels.WorkflowModel) string {
	defaultSource := ""
	for _, step := range workflow.Steps {
		if defaultSource == "" {
			defaultSource = step.SteplibSource
		} else if defaultSource != step.SteplibSource {
			return ""
		}
	}
	return defaultSource
}

// ConvertOldWorkflow ...
func ConvertOldWorkflow(oldWorkflow oldmodels.WorkflowModel) (bitriseModels.WorkflowModel, error) {
	// Step conversion
	containsCertificateStep := false

	stepList := []bitriseModels.StepListItemModel{}
	for _, oldStep := range oldWorkflow.Steps {
		oldStepID := oldStep.ID
		newStep, err := oldStep.Convert()
		if err != nil {
			return bitriseModels.WorkflowModel{}, err
		}

		newStepID, converterFunc, found := getNewStepIDAndConverter(oldStepID)
		if found {
			log.Infof("Convertable step found (%s) -> (%s)", oldStepID, newStepID)
			fmt.Println()

			convertedStepListItems, err := converterFunc(newStep)
			if err != nil {
				return bitriseModels.WorkflowModel{}, err
			}

			for _, stepListItem := range convertedStepListItems {
				stepID, _, err := bitriseModels.GetStepIDStepDataPair(stepListItem)
				if err != nil {
					return bitriseModels.WorkflowModel{}, err
				}

				if strings.Contains(stepID, CertificateStepID) {
					if containsCertificateStep {
						continue
					} else {
						containsCertificateStep = true
					}
				}

				stepList = append(stepList, stepListItem)
			}
		} else {
			log.Infof("Step (%s) not convertable", oldStepID)
			fmt.Println()

			_, _, version := oldStep.GetStepLibIDVersionData()

			stepIDDataString := "_::" + newStep.Source.Git + "@" + version

			stepListItem := bitriseModels.StepListItemModel{
				stepIDDataString: newStep,
			}
			stepList = append(stepList, stepListItem)
		}
	}

	// Workflow environments
	environments, err := oldWorkflow.GetEnvironments()
	if err != nil {
		return bitriseModels.WorkflowModel{}, err
	}

	return bitriseModels.WorkflowModel{
		Environments: environments,
		Steps:        stepList,
	}, nil
}

// ConvertOldWorkfowModels ...
func ConvertOldWorkfowModels(oldWorkflowMap map[string]oldmodels.WorkflowModel) (bitriseModels.BitriseDataModel, error) {
	bitriseData := bitriseModels.BitriseDataModel{
		FormatVersion: "1.0.0",
		Workflows:     map[string]bitriseModels.WorkflowModel{},
	}

	hasDefaultSteplLibSource := true
	defaultSource := ""
	workflowIDs := []string{}

	for workflowID, oldWorkflow := range oldWorkflowMap {
		workflowIDs = append(workflowIDs, workflowID)

		newWorkflow, err := ConvertOldWorkflow(oldWorkflow)
		if err != nil {
			return bitriseModels.BitriseDataModel{}, err
		}

		bitriseData.Workflows[workflowID] = newWorkflow

		if defaultSource == "" {
			defaultSource = GetDefaultSteplibSource(oldWorkflow)
		} else if defaultSource != GetDefaultSteplibSource(oldWorkflow) {
			hasDefaultSteplLibSource = false
		}
	}

	if hasDefaultSteplLibSource {
		bitriseData.DefaultStepLibSource = defaultSource
	}

	triggerMap := []bitriseModels.TriggerMapItemModel{}
	for _, workflowID := range workflowIDs {
		triggerItem := bitriseModels.TriggerMapItemModel{
			Pattern:              workflowID,
			IsPullRequestAllowed: true,
			WorkflowID:           workflowID,
		}
		triggerMap = append(triggerMap, triggerItem)
	}
	bitriseData.TriggerMap = triggerMap

	return bitriseData, nil
}
