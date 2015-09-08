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
		steps.OldBashScriptRunnerStepID:                             steps.ConvertBashScriptRunner,
		steps.OldBitriseIosDeployStepID:                             steps.ConvertBitriseIosDeploy,
		steps.OldCocoapodsAndXcodeRepositoryValidatorFlavorBitrise:  steps.ConvertCocoapodsAndXcodeRepositoryValidatorFlavorBitrise,
		steps.OldCocoapodsFlavorBitriseStepID:                       steps.ConvertCocoapodsFlavorBitrise,
		steps.OldGenericScriptRunnerStepID:                          steps.ConvertGenericScriptRunner,
		steps.OldGitCloneFlavorBitriseStepID:                        steps.ConvertGitCloneFlavorBitrise,
		steps.OldGitCloneFlavorBitriseSSHStepID:                     steps.ConvertGitCloneFlavorBitriseSSH,
		steps.OldHipchatStepID:                                      steps.ConvertHipchat,
		steps.OldSlackMessageStepID:                                 steps.ConvertSlackMessage,
		steps.OldActivateSSHKeyFlavorBitriseStepID:                  steps.ConvertActivateSSHKeyFlavorBitrise,
		steps.OldXcodeBuilderFlavorBitriseAnalyzeStepID:             steps.ConvertXcodeBuilderFlavorBitriseAnalyze,
		steps.OlXcodeBuilderFlavorBitriseCreateArchiveStepID:        steps.ConvertXcodeBuilderFlavorBitriseCreateArchive,
		steps.OldXcodeBuilderFlavorBitriseUnittestStepID:            steps.ConvertXcodeBuilderFlavorBitriseUnittest,
		steps.OldCurlPingStepID:                                     steps.ConvertCurlPing,
		steps.OldItunesConnectDeployWithDeliverFlavorBitriseStepID:  steps.ConvertItunesConnectDeployWithDeliverFlavorBitrise,
		steps.OldItunesConnectDeployWithShenzhenFlavorBitriseStepID: steps.ConvertItunesConnectDeployWithShenzhenFlavorBitrise,
		steps.OldRandomQuoteStepID:                                  steps.ConvertRandomQuote,
		steps.OldAmazonS3UploaderStepID:                             steps.ConvertAmazonS3Uploader,
		steps.OldTestfairyDeployStepID:                              steps.ConvertTestfairyDeploy,
		steps.OldFtpUploadStepID:                                    steps.ConvertFtpUpload,
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

// ConvertOldWorkflow ...
func ConvertOldWorkflow(oldWorkflow oldmodels.WorkflowModel) (bitriseModels.WorkflowModel, string, error) {
	// Step conversion
	containsCertificateStep := false
	defaultStepLib := ""

	stepList := []bitriseModels.StepListItemModel{}
	for _, oldStep := range oldWorkflow.Steps {
		oldStepID := oldStep.ID
		newStep, err := oldStep.Convert()
		if err != nil {
			return bitriseModels.WorkflowModel{}, "", err
		}

		converterFunc, found := getNewStepIDAndConverter(oldStepID)
		if found {
			log.Infof("Convertable step found (%s)", oldStepID)
			fmt.Println()

			convertedStepListItems, err := converterFunc(newStep)
			if err != nil {
				return bitriseModels.WorkflowModel{}, "", err
			}

			for _, stepListItem := range convertedStepListItems {
				stepID, _, err := bitriseModels.GetStepIDStepDataPair(stepListItem)
				if err != nil {
					return bitriseModels.WorkflowModel{}, "", err
				}

				if strings.Contains(stepID, utils.CertificateStepID) {
					if containsCertificateStep {
						continue
					} else {
						containsCertificateStep = true
					}
				}

				defaultStepLib = utils.BitriseVerifiedStepLibGitURI
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
		return bitriseModels.WorkflowModel{}, "", err
	}

	return bitriseModels.WorkflowModel{
		Environments: environments,
		Steps:        stepList,
	}, defaultStepLib, nil
}

func generateTriggerMap(workflowIDs []string) []bitriseModels.TriggerMapItemModel {
	triggerMap := []bitriseModels.TriggerMapItemModel{}
	for _, workflowID := range workflowIDs {
		triggerItem := bitriseModels.TriggerMapItemModel{
			Pattern:              workflowID,
			IsPullRequestAllowed: true,
			WorkflowID:           workflowID,
		}
		triggerMap = append(triggerMap, triggerItem)
	}
	triggerItem := bitriseModels.TriggerMapItemModel{
		Pattern:              "*",
		IsPullRequestAllowed: true,
		WorkflowID:           "primary",
	}
	triggerMap = append(triggerMap, triggerItem)
	return triggerMap
}

// ConvertOldWorkfowModels ...
func ConvertOldWorkfowModels(oldWorkflowMap map[string]oldmodels.WorkflowModel) (bitriseModels.BitriseDataModel, error) {
	bitriseData := bitriseModels.BitriseDataModel{
		FormatVersion: "1.0.0",
		Workflows:     map[string]bitriseModels.WorkflowModel{},
	}

	defaultSource := ""
	workflowIDs := []string{}

	for workflowID, oldWorkflow := range oldWorkflowMap {
		workflowIDs = append(workflowIDs, workflowID)

		newWorkflow, defaultSteplib, err := ConvertOldWorkflow(oldWorkflow)
		if err != nil {
			return bitriseModels.BitriseDataModel{}, err
		}

		if defaultSteplib != "" {
			defaultSource = defaultSteplib
		}

		bitriseData.Workflows[workflowID] = newWorkflow
	}

	bitriseData.DefaultStepLibSource = defaultSource
	bitriseData.TriggerMap = generateTriggerMap(workflowIDs)

	return bitriseData, nil
}
