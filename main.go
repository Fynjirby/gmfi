package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s %s %s\n", green("gmfi"), blue("<file>"), pink("[or more files]"))
		return
	}

	switch os.Args[1] {
	case "find", "search", "grep", "rg":
		if len(os.Args) < 3 {
			fmt.Printf("usage: %s %s\n", green("gmfi find"), blue("<pattern> [path]"))
			return
		}
		pattern := os.Args[2]
		path := "."
		if len(os.Args) >= 4 {
			path = os.Args[3]
		}
		searchIn(pattern, path)

	case "compare", "comp":
		if len(os.Args) != 4 {
			fmt.Printf("usage: %s %s\n", green("gmfi compare"), blue("<file> <another file>"))
			return
		}
		compareFiles(os.Args[2], os.Args[3])

	case "tree", "list":
		dir := "."
		if len(os.Args) >= 3 {
			dir = os.Args[2]
		}
		treeCommand(dir)

	case "biggest", "smallest":
		topN := 5
		dir := "."

		if len(os.Args) >= 3 {
			topN, _ = strconv.Atoi(os.Args[2])
			if len(os.Args) >= 4 {
				dir = os.Args[3]
			}
		}

		sortFiles(os.Args[1], topN, dir)

	case "--help", "-h":
		printHelp()
		return

	case "--version", "-v":
		fmt.Printf("gmfi %s\n", Version)
		return

	default:
		for _, file := range os.Args[1:] {
			meta, err := getMeta(file)
			if err != nil {
				fmt.Printf("%v\n", red(err))
				return
			}

			fmt.Printf("> %s (%s) - %s [%s] | %s\n", red(meta.Name), green(meta.Size), yellow(meta.Type), blue(meta.Perm), meta.Mod)
		}
	}
}

func printHelp() {
	fmt.Printf("usage:\n")
	fmt.Printf(" %s %s %s\n", green("gmfi"), blue("<file>"), pink("[or more files]"))

	fmt.Printf("commands:\n")
	fmt.Printf(" %s %s %s\n  %s\n", green(fmt.Sprintf("%-8s", "find")), blue("<pattern>"), pink("[path]"), "find files in directory")
	fmt.Printf(" %s %s\n  %s\n", green(fmt.Sprintf("%-8s", "compare")), blue("<file> <another file>"), "compare two files")
	fmt.Printf(" %s %s\n  %s\n", green(fmt.Sprintf("%-8s", "tree")), pink("[path]"), "display folder structure")
	fmt.Printf(" %s %s\n  %s\n", green(fmt.Sprintf("%-8s", "biggest")), pink("[count] [path]"), "show biggest files in a directory")
	fmt.Printf(" %s %s\n  %s\n", green(fmt.Sprintf("%-8s", "smallest")), pink("[count] [path]"), "show smallest files in a directory")

	fmt.Printf("flags:\n")
	fmt.Printf(" %s | %s\n", pink("-h"), pink("--help"))
	fmt.Printf(" %s | %s\n", pink("-v"), pink("--version"))
	fmt.Printf("%s\n", yellow("https://github.com/jvqtil/gmfi/"))
}
