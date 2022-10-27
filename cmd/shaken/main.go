package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'lint' or 'download' subcommands")
		os.Exit(1)
	}

	var err error

	switch os.Args[1] {
	case "lint":
		// cmd args
		lintCmd := flag.NewFlagSet("lint", flag.ExitOnError)
		var summary bool
		lintCmd.BoolVar(&summary, "summary", false, "enables summary reporting")

		lintCmd.Parse(os.Args[2:])
		certPath := lintCmd.Arg(0)

		// run command
		err = RunLintCommand(certPath, summary)
	case "download":
		// cmd args
		downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
		var outDir string
		downloadCmd.StringVar(&outDir, "outDir", "certs", "output folder for downloading certificates")
		var includeCa bool
		downloadCmd.BoolVar(&includeCa, "includeCa", false, "downloads CA certificates")

		downloadCmd.Parse(os.Args[2:])
		listPath := downloadCmd.Arg(0)

		// run command
		err = RunDownloadCommand(listPath, outDir, includeCa)
	default:
		fmt.Println("expected 'lint' or 'download' subcommands")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("error on %s command running, %s\n", os.Args[1], err.Error())
	}
}
