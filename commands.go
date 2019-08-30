package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

type ExecConfig struct {
	CmdStr  string
	Args    []string
	Outfile string
}

// Cat accepts a list of file names, and returns the contents of those files combined by a newline character.
//
// If a filename is invalid, it will be included in the result
//
// When used via CLI, will log to stdout
func Cat(targets ...string) string {
	var fileContents []byte
	var err error
	strs := make([]string, 0)
	for _, f := range targets {
		fileContents, err = ioutil.ReadFile(f)
		if err == nil {
			strs = append(strs, string(fileContents))
		} else {
			strs = append(strs, f)
		}
	}
	return strings.Join(strs, "\n")
}

// Attempts to change current working directory
func CD(path string) {
	err := os.Chdir(path)
	checkErr(err)
}

// A simple convenience wrapper for os.Chmod so that all shell commands can be called from the same package
//
// It is important to note that the argument order is reversed compared to go. This is because this is how the command
// would be invoked in a shell, and that is how shelljs does it as well.
func Chmod(mode int32, name string) {
	err := os.Chmod(name, os.FileMode(mode))
	checkErr(err)
}

// Exec runs a literal command strings with arguments
//
// Can be run silently, if silent is false will return output as a string
func Exec(config *ExecConfig) string {
	cmd := exec.Command(config.CmdStr, config.Args...)

	if config.Outfile != "" {
		stdoutPipe, err := cmd.StdoutPipe()
		checkErr(err)

		outfile := Touch("./out.txt")
		defer outfile.Close()
		writer := bufio.NewWriter(outfile)
		go io.Copy(writer, stdoutPipe)
		cmd.Wait()

		return ""
	} else {
		output, err := cmd.Output()
		checkErr(err)
		return string(output)
	}
}

// Ls returns a list of os.FileInfo structs, essentially calling Lstat on each file in the directory
func Ls() []os.FileInfo {
	config := ExecConfig{CmdStr: "ls"}
	output := Exec(&config)
	outputSlice := strings.Split(output, "\n")
	sort.Strings(outputSlice)

	files := make([]os.FileInfo, len(outputSlice))
	var fileTmp os.FileInfo
	var err error
	for i, fileName := range outputSlice {
		fileTmp, err = os.Lstat(fileName)
		checkErr(err)
		files[i] = fileTmp
	}

	return files
}

// Create a file in the current directory
func Touch(fileName string) *os.File {
	outFile, err := os.Create(fileName)
	checkErr(err)

	return outFile
}

// Sed accepts three strings:
//
// a valid regexp to be compiled, the replacement string (supports capture groups), and a path to the target file
func Sed(target string, regexStr string, replace string) {
	fileInfo, err := os.Lstat(target)
	if err != nil {
		// assume target is string literal
		sedString(target, regexStr, replace)
	}

	file, err := ioutil.ReadFile(target)
	checkErr(err)

	rgx, err := regexp.Compile(regexStr)
	checkErr(err)

	fileStr := string(file)
	result := rgx.ReplaceAllString(fileStr, replace)
	err = ioutil.WriteFile(target, []byte(result), fileInfo.Mode())
	checkErr(err)
}

func sedString(target string, regexStr string, replace string) string {
	rgx, err := regexp.Compile(regexStr)
	checkErr(err)
	result := rgx.ReplaceAllString(replace, target)
	return result
}
