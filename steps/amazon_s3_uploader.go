package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldAmazonS3UploaderStepID ...
	OldAmazonS3UploaderStepID = "amazon-s3-uploader"
	// NewAmazonS3UploaderStepID ...
	NewAmazonS3UploaderStepID = "amazon-s3-upload"
)

//----------------------
// old name: amazon-s3-uploader
// new name: amazon-s3-upload

/*
old version source: https://github.com/bitrise-io/steps-amazon-s3-upload.git

inputs:
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- S3_UPLOAD_BUCKET
- S3_UPLOAD_LOCAL_PATH
- S3_ACL_CONTROL
*/

/*
new version source: https://github.com/bitrise-io/steps-amazon-s3-upload.git

inputs:
- access_key_id
- secret_access_key
- upload_bucket
- upload_local_path
- acl_control
*/

// ConvertAmazonS3Uploader ...
func ConvertAmazonS3Uploader(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewAmazonS3UploaderStepID
	inputConversionMap := map[string]string{
		"access_key_id":     "AWS_ACCESS_KEY_ID",
		"secret_access_key": "AWS_SECRET_ACCESS_KEY",
		"upload_bucket":     "S3_UPLOAD_BUCKET",
		"upload_local_path": "S3_UPLOAD_LOCAL_PATH",
		"acl_control":       "S3_ACL_CONTROL",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
