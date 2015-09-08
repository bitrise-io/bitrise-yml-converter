package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldFtpUploadStepID ...
	OldFtpUploadStepID = "ftp-upload"
	// NewFtpUploadStepID ...
	NewFtpUploadStepID = "ftp-upload"
)

//----------------------
// old name: amazon-s3-uploader
// new name: amazon-s3-upload

/*
old version source: https://github.com/bitrise-io/steps-ftp-upload.git

inputs:
- FTP_HOSTNAME
- FTP_USERNAME
- FTP_PASSWORD
- FTP_UPLOAD_SOURCE_PATH
- FTP_UPLOAD_TARGET_PATH
*/

/*
new version source: https://github.com/bitrise-io/steps-ftp-upload.git

inputs:
- hostname
- username
- password
- upload_source_path
- upload_target_path
*/

// ConvertFtpUpload ...
func ConvertFtpUpload(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewAmazonS3UploaderStepID
	inputConversionMap := map[string]string{
		"hostname":           "FTP_HOSTNAME",
		"username":           "FTP_USERNAME",
		"password":           "FTP_PASSWORD",
		"upload_source_path": "FTP_UPLOAD_SOURCE_PATH",
		"upload_target_path": "FTP_UPLOAD_TARGET_PATH",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
