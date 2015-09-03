package converter

import (
	"strings"

	bitriseModels "github.com/bitrise-io/bitrise/models"
	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/go-utils/pointers"
	stepmanModels "github.com/bitrise-io/stepman/models"
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

func convertXcodeBuilderFlavorBitriseUnittest(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	simulatorOsVersion := ""

	// Converter function overwrites
	convertStepsInputs := func(originalInputs, diffInputs []envmanModels.EnvironmentItemModel, conversionMap map[string]string) ([]envmanModels.EnvironmentItemModel, error) {
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

			// Get ios version from old comma separated form
			if specKey == "simulator_device" {
				if workflowValue != "" {
					versionSplit := strings.Split(workflowValue, ",")
					simulatorOsVersion = versionSplit[len(versionSplit)-1]
				}
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

		for idx, input := range mergedStepInputs {
			specKey, _, err := input.GetKeyValuePair()
			if err != nil {
				return []envmanModels.EnvironmentItemModel{}, err
			}
			if specKey == "simulator_os_version" {
				input[specKey] = simulatorOsVersion
				mergedStepInputs[idx] = input
				break
			}
		}

		return mergedStepInputs, nil
	}

	convertStep := func(convertedWorkflowStep stepmanModels.StepModel, newStepID string, inputConversionMap map[string]string) (stepmanModels.StepModel, error) {
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

	stepListItems, err := certificateStep()
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	// Convert Xcode test step
	newStepID := newXcodeTest
	inputConversionMap := map[string]string{
		"project_path":     "XCODE_BUILDER_PROJECT_PATH",
		"scheme":           "XCODE_BUILDER_SCHEME",
		"simulator_device": "XCODE_BUILDER_UNITTEST_PLATFORM_NAME",
	}

	newStep, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := BitriseVerifiedStepLibGitURI + "::" + newStepID
	stepListItems = append(stepListItems, bitriseModels.StepListItemModel{stepIDDataString: newStep})

	return stepListItems, nil
}
