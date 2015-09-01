package converter

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
	oldModels "github.com/bitrise-io/bitrise-yml-converter/old_models"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	"github.com/bitrise-io/go-utils/cmdex"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	stepmanModels "github.com/bitrise-io/stepman/models"
	"gopkg.in/yaml.v2"
)

// GetStepFromNewSteplib ...
func GetStepFromNewSteplib(stepID, stepLibGitURI string) (stepmanModels.StepModel, error) {
	// Activate step - get step.yml
	tempStepCloneDirPath, err := pathutil.NormalizedOSTempDirPath("step_clone")
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	tempStepYMLDirPath, err := pathutil.NormalizedOSTempDirPath("step_yml")
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	tempStepYMLFilePath := path.Join(tempStepYMLDirPath, "step.yml")

	logLevel := log.GetLevel().String()
	args := []string{"--debug", "--loglevel", logLevel, "activate", "--collection", stepLibGitURI, "--id", stepID, "--path", tempStepCloneDirPath, "--copyyml", tempStepYMLFilePath, "--update"}
	if err := cmdex.RunCommand("stepman", args...); err != nil {
		return stepmanModels.StepModel{}, err
	}

	specStep, err := ReadSpecStep(tempStepYMLFilePath)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}

	// Cleanup
	if err := cmdex.RemoveDir(tempStepCloneDirPath); err != nil {
		return stepmanModels.StepModel{}, errors.New(fmt.Sprint("Failed to remove step clone dir: ", err))
	}
	if err := cmdex.RemoveDir(tempStepYMLDirPath); err != nil {
		return stepmanModels.StepModel{}, errors.New(fmt.Sprint("Failed to remove step clone dir: ", err))
	}

	return specStep, nil
}

// GetStepFromGit ...
func GetStepFromGit(stepGitURI string) (stepmanModels.StepModel, error) {
	// Activate step - get step.yml
	tempStepCloneDirPath, err := pathutil.NormalizedOSTempDirPath("step_clone")
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	tempStepYMLDirPath, err := pathutil.NormalizedOSTempDirPath("step_yml")
	if err != nil {
		return stepmanModels.StepModel{}, err
	}
	tempStepYMLFilePath := path.Join(tempStepYMLDirPath, "step.yml")

	if err := cmdex.GitClone(stepGitURI, tempStepCloneDirPath); err != nil {
		return stepmanModels.StepModel{}, err
	}

	if err := cmdex.CopyFile(path.Join(tempStepCloneDirPath, "step.yml"), tempStepYMLFilePath); err != nil {
		return stepmanModels.StepModel{}, err
	}

	specStep, err := ReadSpecStep(tempStepYMLFilePath)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}

	// Cleanup
	if err := cmdex.RemoveDir(tempStepCloneDirPath); err != nil {
		return stepmanModels.StepModel{}, errors.New(fmt.Sprint("Failed to remove step clone dir: ", err))
	}
	if err := cmdex.RemoveDir(tempStepYMLDirPath); err != nil {
		return stepmanModels.StepModel{}, errors.New(fmt.Sprint("Failed to remove step clone dir: ", err))
	}

	return specStep, nil
}

// ReadSpecStep ...
func ReadSpecStep(pth string) (stepmanModels.StepModel, error) {
	if isExists, err := pathutil.IsPathExists(pth); err != nil {
		return stepmanModels.StepModel{}, err
	} else if !isExists {
		return stepmanModels.StepModel{}, errors.New(fmt.Sprint("No file found at path", pth))
	}

	bytes, err := fileutil.ReadBytesFromFile(pth)
	if err != nil {
		return stepmanModels.StepModel{}, err
	}

	var stepModel stepmanModels.StepModel
	if err := yaml.Unmarshal(bytes, &stepModel); err != nil {
		return stepmanModels.StepModel{}, err
	}

	if err := stepModel.Normalize(); err != nil {
		return stepmanModels.StepModel{}, err
	}

	if err := stepModel.ValidateStep(false); err != nil {
		return stepmanModels.StepModel{}, err
	}

	if err := stepModel.FillMissingDefaults(); err != nil {
		return stepmanModels.StepModel{}, err
	}

	return stepModel, nil
}

// ReadOldWorkflowModel ...
func ReadOldWorkflowModel(pth string) (oldModels.WorkflowModel, error) {
	bytes, err := fileutil.ReadBytesFromFile(pth)
	if err != nil {
		return oldModels.WorkflowModel{}, err
	}

	if strings.HasSuffix(pth, ".json") {
		log.Debugln("=> Using JSON parser for: ", pth)
		return WorkflowModelFromJSONBytes(bytes)
	}

	log.Debugln("=> Using YAML parser for: ", pth)
	return WorkflowModelFromYAMLBytes(bytes)
}

// WorkflowModelFromYAMLBytes ...
func WorkflowModelFromYAMLBytes(bytes []byte) (workflow oldModels.WorkflowModel, err error) {
	if err = yaml.Unmarshal(bytes, &workflow); err != nil {
		return
	}
	return
}

// WorkflowModelFromJSONBytes ...
func WorkflowModelFromJSONBytes(bytes []byte) (workflow oldModels.WorkflowModel, err error) {
	if err = json.Unmarshal(bytes, &workflow); err != nil {
		return
	}
	return
}

// WriteNewWorkflowModel ...
func WriteNewWorkflowModel(pth string, newWorkflow bitriseModels.BitriseDataModel) error {
	bytes, err := yaml.Marshal(newWorkflow)
	if err != nil {
		return err
	}
	if err := fileutil.WriteBytesToFile(pth, bytes); err != nil {
		return err
	}
	return nil
}
