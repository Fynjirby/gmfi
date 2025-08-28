package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	pink   = color.New(color.FgMagenta).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\nusage: %s %s %s\n", green("gmfi"), blue("<filename>"), pink("[or more files]"))
		fmt.Printf(yellow("\ngithub.com/fynjirby/gmfi\n"))
		return
	}

	switch os.Args[1] {
	case "diff":
		if len(os.Args) != 4 {
			fmt.Printf("\nusage: %s %s %s\n", green("gmfi"), red("diff"), blue("<file1> <file2>"))
			return
		}
		diffFiles(os.Args[2], os.Args[3])

	case "view":
		if len(os.Args) < 3 {
			fmt.Println("\nusage: %s %s %s\n", green("gmfi"), red("view"), blue("<filename"))
		}
		file := os.Args[2]
		_, err := exec.LookPath("fat")
		if err != nil {
			fmt.Printf(red("fat is not installed — install it from github.com/Zuhaitz-dev/fat\n"))
			return
		}
		cmd := exec.Command("fat", file)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	case "big":
		dir := "."
		topN := 5

		if len(os.Args) >= 3 {
			if n, err := strconv.Atoi(os.Args[2]); err == nil {
				topN = n
			} else {
				dir = os.Args[2]
			}
		}

		if len(os.Args) >= 4 {
			if n, err := strconv.Atoi(os.Args[3]); err == nil {
				topN = n
			} else {
				dir = os.Args[3]
			}
		}

		bigFiles(dir, topN)

	case "tree":
		dir := "."
		if len(os.Args) >= 3 {
			dir = os.Args[2]
		}
		treeCommand(dir)

	case "help":
		printHelp()
		return

	default:
		for _, file := range os.Args[1:] {
			showInfo(file)
		}
	}
}

func printHelp() {
	fmt.Printf("\nusage: %s %s %s\n\n", green("gmfi"), blue("<filename>"), pink("[or more files]"))

	fmt.Printf("%s > %s\n", blue("diff"), "compare two files")
	fmt.Printf("%s > %s\n", blue("view"), "view file or archive contents (via FAT)")
	fmt.Printf("%s > %s\n", blue("tree"), "display folder structure")
	fmt.Printf("%s  > %s\n", blue("big"), "show biggest files in a directory")
	fmt.Printf("%s > %s\n", blue("help"), "show this help message")

	fmt.Printf("\n%s\n", yellow("github.com/fynjirby/gmfi"))
	return
}

func showInfo(file string) {
	meta, err := GetFileMeta(file)
	if err != nil {
		fmt.Printf(red("\nerror reading %s\n"), file)
		return
	}

	fmt.Printf("\n> %s (%s) - %s [%s]\n", red(meta.Name), green(meta.Size), yellow(meta.Type), blue(meta.Perm))
	fmt.Printf("%s * %s\n", pink(meta.Path), cyan(meta.Mod))
	return
}
