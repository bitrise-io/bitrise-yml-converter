package converter

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	oldmodels "github.com/bitrise-io/bitrise-yml-converter/old_models"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldSlackGitURI ...
	OldSlackGitURI = "https://github.com/bitrise-io/steps-slack-message.git"
	// NewSlackStepID ...
	NewSlackStepID = "slack"

	// OldHipchatGitURI ...
	OldHipchatGitURI = "https://github.com/bitrise-io/steps-hipchat.git"
	// NewHipchatStepID ...
	NewHipchatStepID = "hipchat"

	// OldGenericScriptRunnerGitURI ...
	OldGenericScriptRunnerGitURI = "https://github.com/bitrise-io/steps-generic-script-runner.git"
	// NewScriptStepID ...
	NewScriptStepID = "script"

	// OlXcodeBuilderFlavorBitriseCreateArchiveGitURI ...
	OlXcodeBuilderFlavorBitriseCreateArchiveGitURI = "https://github.com/bitrise-io/steps-xcode-builder.git"
	// NewXcodeArchiveStepID ...
	NewXcodeArchiveStepID = "xcode-archive"
)

type stepConverter func(stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error)

// New step ID <-> Converter function
func getStepConverterFunctionMap() map[string]stepConverter {
	return map[string]stepConverter{
		NewSlackStepID:        convertSlack,
		NewHipchatStepID:      convertHipchat,
		NewScriptStepID:       converScript,
		NewXcodeArchiveStepID: converXcodeArchive,
	}
}

// Old step git URI <-> New step ID
func getStepConversionMap() map[string]string {
	return map[string]string{
		OldSlackGitURI:                                 NewSlackStepID,
		OldHipchatGitURI:                               NewHipchatStepID,
		OldGenericScriptRunnerGitURI:                   NewScriptStepID,
		OlXcodeBuilderFlavorBitriseCreateArchiveGitURI: NewXcodeArchiveStepID,
	}
}

func getNewStepIDAndConverter(stepGitURI string) (string, stepConverter, bool) {
	stepConversionMap := getStepConversionMap()
	newID, found := stepConversionMap[stepGitURI]
	if !found {
		return "", nil, false
	}

	converterFunctionMap := getStepConverterFunctionMap()
	converter, found := converterFunctionMap[newID]
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
	environments, err := oldWorkflow.GetEnvironments()
	if err != nil {
		return bitriseModels.WorkflowModel{}, err
	}

	newWorkflow := bitriseModels.WorkflowModel{
		Environments: environments,
	}

	stepList := []bitriseModels.StepListItemModel{}
	for _, oldStep := range oldWorkflow.Steps {
		newStep, err := oldStep.Convert()
		if err != nil {
			return bitriseModels.WorkflowModel{}, err
		}

		newStepID, converterFunc, found := getNewStepIDAndConverter(newStep.Source.Git)
		if found {
			log.Infof("Convertable step found (%s) -> (%s)", newStep.Source.Git, newStepID)
			fmt.Println()

			convertedStepListItems, err := converterFunc(newStep)
			if err != nil {
				return bitriseModels.WorkflowModel{}, err
			}

			for _, stepListItem := range convertedStepListItems {
				stepList = append(stepList, stepListItem)
			}

		} else {
			log.Infof("Step (%s) not convertable", newStep.Source.Git)
			fmt.Println()

			_, _, version := oldStep.GetStepLibIDVersionData()

			stepIDDataString := "_::" + newStep.Source.Git + "@" + version

			stepListItem := bitriseModels.StepListItemModel{
				stepIDDataString: newStep,
			}
			stepList = append(stepList, stepListItem)
		}
	}
	newWorkflow.Steps = stepList

	return newWorkflow, nil
}

// ConvertOldWorkfowModels ...
func ConvertOldWorkfowModels(oldWorkflowMap map[string]oldmodels.WorkflowModel) (bitriseModels.BitriseDataModel, error) {
	bitriseData := bitriseModels.BitriseDataModel{
		FormatVersion: "1.0.0",
		Workflows:     map[string]bitriseModels.WorkflowModel{},
	}

	hasDefaultSteplLibSource := true
	defaultSource := ""
	for workflowID, oldWorkflow := range oldWorkflowMap {
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

	return bitriseData, nil
}
