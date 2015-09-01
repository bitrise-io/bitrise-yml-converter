package cli

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/bitrise-yml-converter/converter"
	oldModels "github.com/bitrise-io/bitrise-yml-converter/old_models"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/codegangsta/cli"
)

func getWorkflowNameFromPath(pth string) string {
	_, file := filepath.Split(pth)
	if strings.HasSuffix(file, ".yml") {
		fileSplit := strings.Split(file, ".yml")
		file = fileSplit[0]
	}
	return file
}

func convert(c *cli.Context) {
	// Input validation
	src := c.String(SourceKey)
	if src == "" {
		log.Fatal("Missing source")
	}

	sources := []string{}
	srcSlice := strings.Split(src, ",")
	if len(srcSlice) > 1 {
		// Comma separated sources
		log.Info("Converting workflows at:", srcSlice)
		sources = srcSlice
	} else {
		isDir, err := pathutil.IsDirExists(src)
		if err != nil {
			log.Fatal("Failed to check path:", err)
		}
		if isDir {
			// Converting workflows in directory
			log.Info("Converting workflows in dir:", src)
			if err := filepath.Walk(src, func(path string, f os.FileInfo, err error) error {
				if filepath.Ext(path) == ".yml" {
					sources = append(sources, path)
				}
				return nil
			}); err != nil {
				log.Fatal("Faild to collect workflow pathes")
			}
			log.Info("Converting workflows at:", sources)
		} else {
			// Converting single workflow
			log.Info("Converting single workflows at:", src)
			sources = append(sources, src)
		}
	}

	dstPth := c.String(DestinationKey)
	if dstPth == "" {
		log.Fatal("Missing destination")
	}

	// Read old workflow
	oldWorkflowMap := map[string]oldModels.WorkflowModel{}
	for _, srcPth := range sources {
		log.Infoln("Converting workflow at:", srcPth)
		oldWorkflow, err := converter.ReadOldWorkflowModel(srcPth)
		if err != nil {
			log.Fatal("Failed to read old workflow:", err)
		}
		oldWorkflowID := getWorkflowNameFromPath(srcPth)

		log.Debugln("Old workflow:")
		log.Debugf("%#v", oldWorkflow)
		oldWorkflowMap[oldWorkflowID] = oldWorkflow
	}

	// Convert workflow
	newConfig, err := converter.ConvertOldWorkfowModels(oldWorkflowMap)
	if err != nil {
		log.Fatal("Failed to convert old workflow:", err)
	}
	log.Debugln("New workflow:")
	log.Debugf("%#v", newConfig)

	// Write new wokrflow to file
	if err := converter.WriteNewWorkflowModel(dstPth, newConfig); err != nil {
		if err != nil {
			log.Fatal("Failed to write new workflow:", err)
		}
	}

	log.Infoln("Converted workflow path:", dstPth)
}
