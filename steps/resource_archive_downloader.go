package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldResourceArchiveDownloaderStepID ...
	OldResourceArchiveDownloaderStepID = "resource-archive-downloader"
	// NewResourceArchiveDownloaderStepID ...
	NewResourceArchiveDownloaderStepID = "resource-archive"
)

//----------------------
// old name: resource-archive-downloader
// new name: resource-archive

/*
old version source: https://github.com/bitrise-io/steps-resource-archive.git

inputs:
- RESOURCE_ARCHIVE_URL
- EXTRACT_TO_PATH
*/

/*
new version source: https://github.com/bitrise-io/steps-resource-archive.git

inputs:
- archive_url
- extract_to_path
*/

// ConvertResourceArchiveDownloader ...
func ConvertResourceArchiveDownloader(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewResourceArchiveDownloaderStepID
	inputConversionMap := map[string]string{
		"archive_url":     "RESOURCE_ARCHIVE_URL",
		"extract_to_path": "EXTRACT_TO_PATH",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
