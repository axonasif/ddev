package ddevapp

import (
	"github.com/drud/ddev/pkg/fileutil"
	"github.com/drud/ddev/pkg/globalconfig"
	"github.com/drud/ddev/pkg/util"
	"os"
	"path/filepath"
)

// PopulateCustomCommandFiles sets up the needed directories and files
func PopulateCustomCommandFiles(app *DdevApp) error {

	sourceGlobalCommandPath := filepath.Join(globalconfig.GetGlobalDdevDir(), "commands")
	err := os.MkdirAll(sourceGlobalCommandPath, 0755)
	if err != nil {
		return nil
	}

	projectCommandPath := app.GetConfigPath("commands")
	// Make sure our target global command directory is empty
	copiedGlobalCommandPath := app.GetConfigPath(".global_commands")
	err = os.RemoveAll(copiedGlobalCommandPath)
	if err != nil {
		util.Error("Unable to remove %s: %v", copiedGlobalCommandPath, err)
		return nil
	}

	err = fileutil.CopyDir(sourceGlobalCommandPath, copiedGlobalCommandPath)
	if err != nil {
		return err
	}

	if !fileutil.FileExists(projectCommandPath) || !fileutil.IsDirectory(projectCommandPath) {
		return nil
	}
	return nil
}
