package utils

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
	CertificateStepID = "certificate-and-profile-installer"
)

//----------------------
// Common methods

//
// Converted inputs should:
// * contain all spec inputs
// * keep all original input value (except empty value: "")
// * keep original IsExpand value
//
// originalInputs: readed from StepLib
// diffInputs: readed from workflow to convert
//
func convertStepsInputs(originalInputs, diffInputs []envmanModels.EnvironmentItemModel, conversionMap map[string]string) ([]envmanModels.EnvironmentItemModel, error) {
	convertedInputs := []envmanModels.EnvironmentItemModel{}

	for _, originalInput := range originalInputs {
		originalInputKey, originalInputValue, err := originalInput.GetKeyValuePair()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		originalInputOptions, err := originalInput.GetOptions()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		conversionInputKey, found := conversionMap[originalInputKey]
		if found == false {
			convertedInputs = append(convertedInputs, originalInput)
			continue
		}

		diffInput, found, err := GetInputByKey(diffInputs, conversionInputKey)
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		if !found {
			convertedInputs = append(convertedInputs, originalInput)
			continue
		}

		_, diffInputValue, err := diffInput.GetKeyValuePair()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}
		if diffInputValue == "" {
			diffInputValue = originalInputValue
		}

		diffInputOptions, err := diffInput.GetOptions()
		if err != nil {
			return []envmanModels.EnvironmentItemModel{}, err
		}

		if diffInputOptions.IsExpand != nil {
			originalInputOptions.IsExpand = pointers.NewBoolPtr(*diffInputOptions.IsExpand)
		}

		convertedInput := envmanModels.EnvironmentItemModel{
			originalInputKey:        diffInputValue,
			envmanModels.OptionsKey: originalInputOptions,
		}

		convertedInputs = append(convertedInputs, convertedInput)
	}

	return convertedInputs, nil
}

// ConvertStep ...
//
// Converted step should:
// * keep original Title
// * keep original IsAlwaysRun
//
// newStepID: the id of the new step in new StepLib
// diffStep: readed from workflow to convert
//
func ConvertStep(diffStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) (stepmanModels.StepModel, string, error) {
	originalStep, version, err := GetStepFromNewSteplib(newStepID, BitriseVerifiedStepLibGitURI)
	if err != nil {
		return stepmanModels.StepModel{}, "", err
	}
	if diffStep.Title != nil {
		originalStep.Title = pointers.NewStringPtr(*diffStep.Title)
	}

	if diffStep.IsAlwaysRun != nil {
		originalStep.IsAlwaysRun = pointers.NewBoolPtr(*diffStep.IsAlwaysRun)
	}

	// Merge new StepLib version inputs, with old workflow defined
	mergedInputs, err := convertStepsInputs(originalStep.Inputs, diffStep.Inputs, inputConversionMap)
	if err != nil {
		return stepmanModels.StepModel{}, "", err
	}
	originalStep.Inputs = mergedInputs

	return originalStep, version, nil
}

// ConvertStepAndCreateStepListItem ...
func ConvertStepAndCreateStepListItem(diffStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) ([]bitriseModels.StepListItemModel, error) {
	newStep, version, err := ConvertStep(diffStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := newStepID + "@" + version

	return []bitriseModels.StepListItemModel{
		bitriseModels.StepListItemModel{
			stepIDDataString: newStep,
		},
	}, nil
}

// CertificateStep ...
//
// Cerificate step separated in new StepLib
// so need to insert befor new Xcode steps
func CertificateStep() ([]bitriseModels.StepListItemModel, error) {
	certificateStep, version, err := GetStepFromNewSteplib(CertificateStepID, BitriseVerifiedStepLibGitURI)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	certificateStep.RunIf = pointers.NewStringPtr(".IsCI")

	stepIDDataString := CertificateStepID + "@" + version

	return []bitriseModels.StepListItemModel{
		bitriseModels.StepListItemModel{
			stepIDDataString: certificateStep,
		},
	}, nil
}
