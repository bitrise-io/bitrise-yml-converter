package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// BitriseVerifiedStepLibGitURI ...
	BitriseVerifiedStepLibGitURI = "https://github.com/bitrise-io/bitrise-steplib.git"
	// CertificateStepID ...
	CertificateStepID = "steps-certificate-and-profile-installer"
	// CertificateStepGitURI ...
	CertificateStepGitURI = "https://github.com/bitrise-io/steps-certificate-and-profile-installer.git"
)

//----------------------
// Common methods

func getInputByKey(inputs []envmanModels.EnvironmentItemModel, key string) (envmanModels.EnvironmentItemModel, bool, error) {
	for _, input := range inputs {
		aKey, _, err := input.GetKeyValuePair()
		if err != nil {
			return envmanModels.EnvironmentItemModel{}, false, err
		}
		if aKey == key {
			return input, true, nil
		}
	}
	return envmanModels.EnvironmentItemModel{}, false, nil
}

func convertStepsInputs(originalInputs, diffInputs []envmanModels.EnvironmentItemModel, conversionMap map[string]string) ([]envmanModels.EnvironmentItemModel, error) {
	mergedStepInputs := []envmanModels.EnvironmentItemModel{}
	for _, specInput := range originalInputs {
		specKey, _, err := specInput.GetKeyValuePair()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		workflowInputKey, found := conversionMap[specKey]
		if found == false {
			mergedStepInputs = append(mergedStepInputs, specInput)
			continue
		}

		workflowInput, found, err := getInputByKey(diffInputs, workflowInputKey)
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		_, workflowValue, err := workflowInput.GetKeyValuePair()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		if workflowValue == "" {
			continue
		}

		workflowOptions, err := workflowInput.GetOptions()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		workflowOptions.Title = nil
		workflowOptions.Description = nil
		workflowOptions.Summary = nil
		workflowOptions.ValueOptions = []string{}
		workflowOptions.IsRequired = nil
		workflowOptions.IsDontChangeValue = nil
		// workflowOptions.IsExpand should be keep

		mergedInput := envmanModels.EnvironmentItemModel{
			specKey:                 workflowValue,
			envmanModels.OptionsKey: workflowOptions,
		}

		mergedStepInputs = append(mergedStepInputs, mergedInput)
	}
	return mergedStepInputs, nil
}

func convertStep(convertedWorkflowStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) (stepmanModels.StepModel, error) {
	// The new StepLib version of step
	specStep, err := GetStepFromNewSteplib(newStepID, BitriseVerifiedStepLibGitURI)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	if convertedWorkflowStep.Title != nil && *convertedWorkflowStep.Title != "" {
		specStep.Title = pointers.NewStringPtr(*convertedWorkflowStep.Title)
	}
	if convertedWorkflowStep.IsAlwaysRun != nil {
		specStep.IsAlwaysRun = pointers.NewBoolPtr(*convertedWorkflowStep.IsAlwaysRun)
	}

	// Merge new StepLib version inputs, with old workflow defined
	mergedInputs, err := convertStepsInputs(specStep.Inputs, convertedWorkflowStep.Inputs, inputConversionMap)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	specStep.Inputs = mergedInputs

	return specStep, nil
}

func convertStepAndCreateStepListItem(convertedWorkflowStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) ([]bitriseModels.StepListItemModel, error) {
	newStep, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := BitriseVerifiedStepLibGitURI + "::" + newStepID

	stepListItem := bitriseModels.StepListItemModel{
		stepIDDataString: newStep,
	}

	return []bitriseModels.StepListItemModel{stepListItem}, nil
}

func certificateStep() ([]bitriseModels.StepListItemModel, error) {
	// Cerificate step separated in new StepLib
	// Step (https://github.com/bitrise-io/steps-certificate-and-profile-installer.git)
	// need to insert befor Xcode-Archive
	certificateStep, err := GetStepFromGit(CertificateStepGitURI)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}
	certificateStep.RunIf = pointers.NewStringPtr(".IsCI")
	certificateStep.Title = pointers.NewStringPtr(CertificateStepID)

	stepIDDataString := "git::" + CertificateStepGitURI + "@master"
	return []bitriseModels.StepListItemModel{
		bitriseModels.StepListItemModel{
			stepIDDataString: certificateStep,
		},
	}, nil
}
