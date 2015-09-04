package converter

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	oldmodels "github.com/bitrise-io/bitrise-yml-converter/old_models"
	"github.com/bitrise-io/bitrise-yml-converter/steps"
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

type stepConverter func(stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error)

// New step ID <-> Converter function
func getStepConverterFunctionMap() map[string]stepConverter {
	return map[string]stepConverter{
		steps.OldBashScriptRunnerStepID:                            steps.ConvertBashScriptRunner,
		steps.OldBitriseIosDeployStepID:                            steps.ConvertBitriseIosDeploy,
		steps.OldCocoapodsAndXcodeRepositoryValidatorFlavorBitrise: steps.ConvertCocoapodsAndXcodeRepositoryValidatorFlavorBitrise,
		steps.OldCocoapodsFlavorBitriseStepID:                      steps.ConvertCocoapodsFlavorBitrise,
		steps.OldGenericScriptRunnerStepID:                         steps.ConvertGenericScriptRunner,
		steps.OldGitCloneFlavorBitriseStepID:                       steps.ConvertGitCloneFlavorBitrise,
		steps.OldGitCloneFlavorBitriseSSHStepID:                    steps.ConvertGitCloneFlavorBitriseSSH,
		steps.OldHipchatStepID:                                     steps.ConvertHipchat,
		steps.OldSlackMessageStepID:                                steps.ConvertSlackMessage,
		steps.OldActivateSSHKeyFlavorBitriseStepID:                 steps.ConvertActivteSSHKey,
		steps.OldXcodeBuilderFlavorBitriseAnalyzeStepID:            steps.ConvertXcodeBuilderFlavorBitriseAnalyze,
		steps.OlXcodeBuilderFlavorBitriseCreateArchiveStepID:       steps.ConvertXcodeBuilderFlavorBitriseCreateArchive,
		steps.OldXcodeBuilderFlavorBitriseUnittestStepID:           steps.ConvertXcodeBuilderFlavorBitriseUnittest,
	}
}

func getNewStepIDAndConverter(oldStepID string) (stepConverter, bool) {
	converterFunctionMap := getStepConverterFunctionMap()
	converter, found := converterFunctionMap[oldStepID]
	if !found {
		return nil, false
	}
	return converter, true
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

		converterFunc, found := getNewStepIDAndConverter(oldStepID)
		if found {
			log.Infof("Convertable step found (%s)", oldStepID)
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

				if strings.Contains(stepID, utils.CertificateStepID) {
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
