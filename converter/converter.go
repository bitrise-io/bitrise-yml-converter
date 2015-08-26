package converter

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	bitriseModels "github.com/bitrise-io/bitrise/models/models_1_0_0"
	"github.com/bitrise-io/go-utils/fileutil"
	oldmodels "github.com/bitrise-io/bitrise-yml-converter/models_0_9_0"
	"gopkg.in/yaml.v2"
)

// ReadOldWorkflowModel ...
func ReadOldWorkflowModel(pth string) (oldmodels.WorkflowModel, error) {
	bytes, err := fileutil.ReadBytesFromFile(pth)
	if err != nil {
		return oldmodels.WorkflowModel{}, err
	}

	if strings.HasSuffix(pth, ".json") {
		log.Debugln("=> Using JSON parser for: ", pth)
		return WorkflowModelFromJSONBytes(bytes)
	}

	log.Debugln("=> Using YAML parser for: ", pth)
	return WorkflowModelFromYAMLBytes(bytes)
}

// WorkflowModelFromYAMLBytes ...
func WorkflowModelFromYAMLBytes(bytes []byte) (workflow oldmodels.WorkflowModel, err error) {
	if err = yaml.Unmarshal(bytes, &workflow); err != nil {
		return
	}
	return
}

// WorkflowModelFromJSONBytes ...
func WorkflowModelFromJSONBytes(bytes []byte) (workflow oldmodels.WorkflowModel, err error) {
	if err = json.Unmarshal(bytes, &workflow); err != nil {
		return
	}
	return
}

// ConvertOldWorkfowModels ...
func ConvertOldWorkfowModels(oldWorkflows ...oldmodels.WorkflowModel) (bitriseModels.BitriseDataModel, error) {
	bitriseData := bitriseModels.BitriseDataModel{
		FormatVersion: "0.9.8",
		Workflows:     map[string]bitriseModels.WorkflowModel{},
	}

	hasDefaultSteplLibSource := true
	defaultSource := ""
	for idx, oldWorkflow := range oldWorkflows {
		newWorkflow := oldWorkflow.Convert()
		newWorkflowName := fmt.Sprintf("target_%d", idx)

		bitriseData.Workflows[newWorkflowName] = newWorkflow

		if defaultSource == "" {
			defaultSource = oldmodels.GetDefaultSteplibSource(oldWorkflow)
		} else if defaultSource != oldmodels.GetDefaultSteplibSource(oldWorkflow) {
			hasDefaultSteplLibSource = false
		}
	}

	if hasDefaultSteplLibSource {
		bitriseData.DefaultStepLibSource = defaultSource
	}

	return bitriseData, nil
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
