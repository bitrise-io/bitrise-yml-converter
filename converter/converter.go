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
	// OldBitriseIosDeployStepID ...
	OldBitriseIosDeployStepID = "bitrise-ios-deploy"
	// NewBitriseIosDeployStepID ...
	NewBitriseIosDeployStepID = "bitrise-ios-deploy"

	// OldHipchatStepID ...
	OldHipchatStepID = "hipchat"
	// NewHipchatStepID ...
	NewHipchatStepID = "hipchat"

	// OldSlackMessageStepID ...
	OldSlackMessageStepID = "slack-message"
	// NewSlackStepID ...
	NewSlackStepID = "slack"

	// OldGenericScriptRunnerStepID ...
	OldGenericScriptRunnerStepID = "generic-script-runner"
	// NewScriptStepID ...
	NewScriptStepID = "script"

	// OlXcodeBuilderFlavorBitriseCreateArchiveStepID ...
	OlXcodeBuilderFlavorBitriseCreateArchiveStepID = "xcode-builder_flavor_bitrise_create-archive"
	// NewXcodeArchiveStepID ...
	NewXcodeArchiveStepID = "xcode-archive"

	// OldXcodeBuilderFlavorBitriseUnittestStepID ...
	OldXcodeBuilderFlavorBitriseUnittestStepID = "xcode-builder_flavor_bitrise_unittest"
	// NewXcodeTest ...
	NewXcodeTest = "xcode-test"

	// OldBashScriptRunnerStepID ...
	OldBashScriptRunnerStepID = "bash-script-runner"
)

type stepConverter func(stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error)

// New step ID <-> Converter function
func getStepConverterFunctionMap() map[string]stepConverter {
	return map[string]stepConverter{
		OldHipchatStepID:                               convertHipchat,
		OldSlackMessageStepID:                          convertSlackMessage,
		OldGenericScriptRunnerStepID:                   convertGenericScriptRunner,
		OldBashScriptRunnerStepID:                      convertBashScriptRunner,
		OlXcodeBuilderFlavorBitriseCreateArchiveStepID: convertXcodeBuilderFlavorBitriseCreateArchive,
		OldXcodeBuilderFlavorBitriseUnittestStepID:     convertXcodeBuilderFlavorBitriseUnittest,
		OldBitriseIosDeployStepID:                      convertBitriseIosDeploy,
	}
}

// Old step ID <-> New step ID
func getStepConversionMap() map[string]string {
	return map[string]string{
		OldHipchatStepID:                               NewHipchatStepID,
		OldSlackMessageStepID:                          NewSlackStepID,
		OldGenericScriptRunnerStepID:                   NewScriptStepID,
		OldBashScriptRunnerStepID:                      NewScriptStepID,
		OlXcodeBuilderFlavorBitriseCreateArchiveStepID: NewXcodeArchiveStepID,
		OldXcodeBuilderFlavorBitriseUnittestStepID:     NewXcodeTest,
		OldBitriseIosDeployStepID:                      NewBitriseIosDeployStepID,
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
