package pkg_test

import (
	"testing"

	"github.com/r1cebucket/gopkg/cmd"
)

func TestCmdExec(t *testing.T) {
	cmd.Exec("mkdir test_folder")
}
