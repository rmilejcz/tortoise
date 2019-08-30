package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	OK               = 0
	GENERAL_FAILURE  = 1
	INCORRECT_USAGE  = 2
	CANNOT_EXECUTE   = 126
	INVALID_ARGUMENT = 128
)

var (
	touchFlag = flag.String("touch", "", "Create a file")
	sedFlag   = flag.String("sed", "", "Comma separated args for sed")
	catFlag   = flag.String("cat", "", "Comma separated list of file names / paths")
)

func main() {
	defer handleErr()
	flag.Parse()

	if *touchFlag != "" {
		Touch(*touchFlag)
	}

	if *sedFlag != "" {
		args := strings.Split(*sedFlag, ",")
		if len(args) != 3 {
			checkErr(errors.New("sed requires exactly 3 arguments"))
		} else {
			Sed(args[0], args[1], args[2])
		}
	}

	if *catFlag != "" {
		args := strings.Split(*catFlag, ",")
		if len(args) < 2 {
			checkErr(errors.New("cat requires at least 2 arguments"))
		} else {
			result := Cat(args...)
			fmt.Println(result)
		}
	}

	fileInfo, err := os.Lstat("tortoise")
	if err != nil {
		checkErr(err)
	}
	str := fmt.Sprintf("%#v\n", fileInfo)
	prettyPrint(str)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func handleErr() {
	err := recover()
	if err != nil {
		os.Stderr.WriteString(err.(error).Error())
		switch err {
		case os.ErrNotExist:
			os.Exit(INCORRECT_USAGE)

		case os.ErrPermission:
			os.Exit(CANNOT_EXECUTE)
		}
	}
}

func getAbsPath() string {
	dir, err := os.Getwd()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		//checkErr(GENERAL_FAILURE)
	}

	return dir
}

func prettyPrint(uglyString string) {
	//builder := strings.Builder{}
	strSlice := strings.Split(uglyString, " ")
	result := make([]string, len(strSlice))

	//bracketAndColon := "{|:"
	//var strPart string
	for i, str := range strSlice {
		if strings.Contains(str, "{") && strings.Contains(str, ":") {
			//strPart = strings.Replace(str, "")
		} else if strings.Contains(str, "{") {
			result[i] = strings.ReplaceAll(str, "{", " {\n")
		} else if strings.Contains(str, ":") {
			result[i] = strings.ReplaceAll(
				strings.ReplaceAll(str, ":", ": ,"),
				"\\",
				"",
			)
		}
	}
	log.Printf("%#v\n", strSlice)
}
