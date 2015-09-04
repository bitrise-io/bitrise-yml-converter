package utils

import (
	"testing"

	envmanModels "github.com/bitrise-io/envman/models"
)

func TestGetInputByKey(t *testing.T) {
	env1 := envmanModels.EnvironmentItemModel{
		"KEY1": "value",
	}

	env2 := envmanModels.EnvironmentItemModel{
		"KEY2": "value",
	}

	env3 := envmanModels.EnvironmentItemModel{
		"KEY3": "value",
	}

	envs := []envmanModels.EnvironmentItemModel{env1, env2, env3}

	_, found, err := GetInputByKey(envs, "KEY1")
	if err != nil {
		t.Fatal("Failed to get input by key (KEY1)")
	}
	if !found {
		t.Fatal("Should found env by key (KEY1)")
	}

	_, found, err = GetInputByKey(envs, "KEY4")
	if err != nil {
		t.Fatal("Failed to get input by key (KEY1)")
	}
	if found {
		t.Fatal("Should not found env by key (KEY4)")
	}

	_, found, err = GetInputByKey(envs, "")
	if err != nil {
		t.Fatal("Failed to get input by key ()")
	}
	if found {
		t.Fatal("Should not found env by key ()")
	}
}
