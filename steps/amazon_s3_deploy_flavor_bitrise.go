package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldAmazonS3DeployFlavorBitriseStepID ...
	OldAmazonS3DeployFlavorBitriseStepID = "amazon-s3-deploy_flavor_bitrise"
	// NewAmazonS3DeployFlavorBitriseStepID ...
	NewAmazonS3DeployFlavorBitriseStepID = "amazon-s3-deploy"
)

//----------------------
// old name: amazon-s3-deploy_flavor_bitrise
// new name: amazon-s3-deploy

/*
old version source: https://github.com/bitrise-io/steps-amazon-s3-deploy.git

inputs:
- BITRISE_IPA_PATH
- BITRISE_DSYM_PATH
- BITRISE_APP_SLUG
- BITRISE_BUILD_SLUG
- S3_DEPLOY_AWS_ACCESS_KEY
- S3_DEPLOY_AWS_SECRET_KEY
- S3_BUCKET_NAME
- S3_PATH_IN_BUCKET
- S3_FILE_ACCESS_LEVEL
*/

/*
new version source: https://github.com/bitrise-io/steps-amazon-s3-deploy.git

inputs:
- ipa_path
- dsym_path
- app_slug
- build_slug
- aws_access_key
- aws_secret_key
- bucket_name
- path_in_bucket
- file_access_level
*/

// ConvertAmazonS3DeployFlavorBitrise ...
func ConvertAmazonS3DeployFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewAmazonS3DeployFlavorBitriseStepID
	inputConversionMap := map[string]string{
		"ipa_path":          "BITRISE_IPA_PATH",
		"dsym_path":         "BITRISE_DSYM_PATH",
		"app_slug":          "BITRISE_APP_SLUG",
		"build_slug":        "BITRISE_BUILD_SLUG",
		"aws_access_key":    "S3_DEPLOY_AWS_ACCESS_KEY",
		"aws_secret_key":    "S3_DEPLOY_AWS_SECRET_KEY",
		"bucket_name":       "S3_BUCKET_NAME",
		"path_in_bucket":    "S3_PATH_IN_BUCKET",
		"file_access_level": "S3_FILE_ACCESS_LEVEL",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
