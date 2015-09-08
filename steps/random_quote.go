package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldRandomQuoteStepID ...
	OldRandomQuoteStepID = "random-quote"
	// NewRandomQuoteStepID ...
	NewRandomQuoteStepID = "random-quote"
)

//----------------------
// old name: random-quote
// new name: random-quote

/*
old version source: https://github.com/bitrise-io/steps-random-quote.git

inputs: []
*/

/*
new version source: https://github.com/bitrise-io/steps-random-quote.git

inputs: []
*/

// ConvertRandomQuote ...
func ConvertRandomQuote(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewItunesConnectDeployWithDeliverFlavorBitriseStepID
	inputConversionMap := map[string]string{}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
