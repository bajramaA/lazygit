package test

import (
	"os"
	"path/filepath"

	"github.com/go-errors/errors"

	"github.com/jesseduffield/lazygit/pkg/secureexec"
	"github.com/jesseduffield/lazygit/pkg/utils"
)

// GenerateRepo generates a repo from test/repos and changes the directory to be
// inside the newly made repo
func GenerateRepo(filename string) error {
	reposDir := "/test/repos/"
	testPath := utils.GetProjectRoot() + reposDir

	// workaround for debian packaging
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		cwd, _ := os.Getwd()
		testPath = filepath.Dir(filepath.Dir(cwd)) + reposDir
	}
	if err := os.Chdir(testPath); err != nil {
		return err
	}
	if output, err := secureexec.Command("bash", filename).CombinedOutput(); err != nil {
		return errors.New(string(output))
	}

	return os.Chdir(testPath + "repo")
}
