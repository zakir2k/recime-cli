package cmd

import "os"
import "os/exec"

// Build builds the bot.
func Build(dir string) error {
	cmd := exec.Command("npm", "run", "build")

	cmd.Dir = dir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
