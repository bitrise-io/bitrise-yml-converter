package steps

import (
	"strings"

	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldXcodeBuilderFlavorBitriseUnittestStepID ...
	OldXcodeBuilderFlavorBitriseUnittestStepID = "xcode-builder_flavor_bitrise_unittest"
	// NewXcodeTest ...
	NewXcodeTest = "xcode-test"
)

//----------------------
// old name: xcode-builder_flavor_bitrise_unittest
// new name: xcode-test

/*
old version source: https://github.com/bitrise-io/steps-xcode-builder.git

inputs:
  - XCODE_BUILDER_PROJECT_ROOT_DIR_PATH
  - XCODE_BUILDER_PROJECT_PATH
  - XCODE_BUILDER_SCHEME
  - XCODE_BUILDER_ACTION
  - XCODE_BUILDER_CERTIFICATE_URL
  - XCODE_BUILDER_CERTIFICATE_PASSPHRASE
  - XCODE_BUILDER_PROVISION_URL
  - XCODE_BUILDER_BUILD_TOOL
  - XCODE_BUILDER_CERTIFICATES_DIR
# UnitTest specific inputs
  - XCODE_BUILDER_UNITTEST_PLATFORM_NAME
*/

/*
new version source: https://github.com/bitrise-io/steps-xcode-test.git

inputs:
- workdir
- project_path
- scheme
- simulator_device
- simulator_os_version
- is_clean_build
*/

/*
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
*/

// ConvertXcodeBuilderFlavorBitriseUnittest ...
func ConvertXcodeBuilderFlavorBitriseUnittest(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	simulatorOsVersion := ""
	simulatorDevice := ""

	// Converter function overwrites
	convertStepsInputs := func(originalInputs, diffInputs []envmanModels.EnvironmentItemModel, conversionMap map[string]string) ([]envmanModels.EnvironmentItemModel, error) {
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

			diffInput, found, err := utils.GetInputByKey(diffInputs, conversionInputKey)
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

			// Get ios version from old comma separated form
			if originalInputKey == "simulator_device" {
				simulatorDevice = diffInputValue
				if diffInputValue != "" {
					versionSplit := strings.Split(diffInputValue, ",")
					if len(versionSplit) == 2 {
						simulatorDevice = versionSplit[0]
						simulatorOsVersion = versionSplit[1]
						if strings.HasPrefix(simulatorOsVersion, "OS=") {
							simulatorOsVersion = strings.Replace(simulatorOsVersion, "OS=", "", -1)
						}
					}
				}
			}

			convertedInput := envmanModels.EnvironmentItemModel{
				originalInputKey:        diffInputValue,
				envmanModels.OptionsKey: originalInputOptions,
			}

			convertedInputs = append(convertedInputs, convertedInput)
		}

		for idx, input := range convertedInputs {
			specKey, _, err := input.GetKeyValuePair()
			if err != nil {
				return []envmanModels.EnvironmentItemModel{}, err
			}
			if specKey == "simulator_device" {
				input[specKey] = simulatorDevice
				convertedInputs[idx] = input
			}
			if specKey == "simulator_os_version" {
				input[specKey] = simulatorOsVersion
				convertedInputs[idx] = input
			}
		}

		return convertedInputs, nil
	}

	convertStep := func(convertedWorkflowStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) (stepmanModels.StepModel, string, error) {
		// The new StepLib version of step
		specStep, version, err := utils.GetStepFromNewSteplib(newStepID, utils.BitriseVerifiedStepLibGitURI)
		if err != nil {
			return stepmanModels.StepModel{}, "", err
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
			return stepmanModels.StepModel{}, "", err
		}
		specStep.Inputs = mergedInputs

		return specStep, version, nil
	}

	stepListItems, err := utils.CertificateStep()
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	// Convert Xcode test step
	newStepID := NewXcodeTest
	inputConversionMap := map[string]string{
		"workdir":          "XCODE_BUILDER_PROJECT_ROOT_DIR_PATH",
		"project_path":     "XCODE_BUILDER_PROJECT_PATH",
		"scheme":           "XCODE_BUILDER_SCHEME",
		"simulator_device": "XCODE_BUILDER_UNITTEST_PLATFORM_NAME",
	}

	newStep, version, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := newStepID + "@" + version
	stepListItems = append(stepListItems, bitriseModels.StepListItemModel{stepIDDataString: newStep})

	return stepListItems, nil
}
