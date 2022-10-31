package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'lint' subcommand")
		os.Exit(1)
	}

	var err error

	switch os.Args[1] {
	case "lint":
		// cmd args
		lintArgs := &LintCommandArgs{}
		lintCmd := flag.NewFlagSet("lint", flag.ExitOnError)
		lintCmd.StringVar(&lintArgs.OutDir, "outDir", REPORT_DIR_NAME, "output folder reports")
		lintCmd.StringVar(&lintArgs.Root, "root", CA_TRUST_LIST, "HTTP URL reference or file path to PEM chain with trusted certificates")

		lintCmd.Parse(os.Args[2:])

		lintArgs.Url = lintCmd.Arg(0)

		// run command
		err = RunLintCommand(lintArgs)
	default:
		fmt.Println("expected 'lint' subcommand")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("error on %s command running, %s\n", os.Args[1], err.Error())
		os.Exit(1)
	}
}
