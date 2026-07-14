package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold [ticket-id]",
	Short: "Scaffold boilerplate code across monorepos for a given Linear ticket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ticket := args[0]
		fmt.Printf("Starting Dual-Repo Scaffold for Ticket: %s\n", ticket)
		
		homeDir, _ := os.UserHomeDir()
		projectDir := filepath.Join(homeDir, "fde-autogen")
		pythonBin := filepath.Join(projectDir, "venv", "bin", "python")
		scriptPath := filepath.Join(projectDir, "fde_scaffold.py")

		execCmd := exec.Command(pythonBin, scriptPath, "--ticket", ticket)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin

		err := execCmd.Run()
		if err != nil {
			fmt.Printf("Error executing scaffold: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
}