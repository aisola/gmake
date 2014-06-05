package main

import "flag"
import "fmt"
import "io/ioutil"
import "os"
import "os/exec"

const (
	help_text string = `
    Usage: gmake [OPTION]...
    
    A very lightweight build tool.

          --help     display this help and exit
          --version  output version information and exit
    `
	version_text = `
    gmake (aisola/gmake) 0.1

    Copyright (C) 2014 Abram C. Isola.
    This program comes with ABSOLUTELY NO WARRANTY; for details see
    LICENSE. This is free software, and you are welcome to redistribute 
    it under certain conditions in LICENSE.
`
)

var AST []*Directive

func getDirective(dname string) *Directive {
	for i := 0; i < len(AST); i++ {
		if AST[i].Name == dname {
			return AST[i]
		}
	}
	return nil
}

func combiner(strs []string) string {
	var fullstr string
	for i := 0; i < len(strs); i++ {
		fullstr = fullstr + strs[i] + " "
	}
	return fullstr
}

// Starts processing
func main() {
	help := flag.Bool("help", false, help_text)
	version := flag.Bool("version", false, version_text)
	flag.Parse()

	if *help {
		fmt.Println(help_text)
		os.Exit(0)

	} else if *version {
		fmt.Println(version_text)
		os.Exit(0)

	} else {
		// get contents
		buf, err := ioutil.ReadFile("GMakefile")
		if err != nil {
			fmt.Println("gmake: fatal: could not read GMakefile")
			return
		}

		// scan then parse
		_, tokens := Lexer("GMAKE", string(buf))
		AST = Parse("GMAKE", tokens)

		args := flag.Args()

		var dir *Directive
		if len(args) == 0 {
			dir = getDirective("all")
			if dir == nil {
				fmt.Println("gmake: fatal: no all directive defined")
				os.Exit(1)
			}
		} else {
			dir = getDirective(args[0])
			if dir == nil {
				fmt.Printf("gmake: fatal: no '%s' directive defined")
				os.Exit(1)
			}
		}

		for i := 0; i < len(dir.Commands); i++ {
			fmt.Println(combiner(dir.Commands[i].Parts))

			cm, parts := dir.Commands[i].Parts[0], dir.Commands[i].Parts[1:]

			cmd := exec.Command(cm, parts...)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("gmake: fatal: '%s'", err)
				os.Exit(1)
			}
		}

	}
}
