package main

import (
	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use: "chr",
		Short: "chr [--engine libc] [--repo ./] [--tag latest] " +
			"echo Hello world",
		Args: cobra.MinimumNArgs(1),

		RunE: chrMain,
	}
)

func init() {

	flags := Command.Flags()
	flags.StringP("engine", "E", "libc", "the fake chroot engine name")
	flags.StringP("repo", "R", "./", "the ketos repo path")
	flags.StringP("tag", "T", "latest", "loading tag from ketos folder")
}

func chrMain(cmd *cobra.Command, args []string) error {

	engineName, err := cmd.Flags().GetString("engine")
	if err != nil {
		return err
	}
	repoPath, err := cmd.Flags().GetString("repo")
	if err != nil {
		return err
	}
	tagName, err := cmd.Flags().GetString("tag")
	if err != nil {
		return err
	}

	userCommand := args

	executor, err := NewChrootExecutor(engineName)
	if err != nil {
		return err
	}

	return executor.Execute(repoPath, tagName, userCommand)
}
