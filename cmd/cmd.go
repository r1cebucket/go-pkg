package cmd

import (
	"os/exec"
	"strings"

	"github.com/r1cebucket/gopkg/log"
)

// command string
func Exec(s string) {
	tokens := strings.Split(s, " ")
	c := exec.Command(tokens[0], tokens[1:]...)
	err := c.Run()
	if err != nil {
		log.Err(err).Msg("Faile to exec instruction")
		return
	}
}
