package main

import (
	"os"
	"strings"
	"testing"
)

const testFile = "_--_tmp.txt"

func TestCat(t *testing.T) {

}

func TestCD(t *testing.T) {
	startDir := getAbsPath()
	CD("../")
	endDir := getAbsPath()
	if strings.EqualFold(startDir, endDir) {
		t.Fatalf("CD func failed to change directory string\nstartDir:%s\nendDir:%sn", startDir, endDir)
	}
}

func TestChmod(t *testing.T) {

	Touch(testFile)
	defer os.Remove(testFile)

	startFileInfo, err := os.Lstat(testFile)
	checkErr(err)

	startVal := int32(startFileInfo.Mode().Perm())
	Chmod(0777, testFile)
	defer Chmod(startVal, testFile)

	endFileInfo, err := os.Lstat(testFile)
	checkErr(err)

	endVal := int32(endFileInfo.Mode().Perm())
	if startVal == endVal {
		t.Fatalf("Chmod failed to change file permissions\nstartVal: %d\nendVal:%d\n", startVal, endVal)
	}
}

func TestExec(t *testing.T) {
	config := ExecConfig{CmdStr: "touch", Args: []string{testFile}}
	Exec(&config)
	defer os.Remove(testFile)
	config.CmdStr = "ls"
	config.Args = nil
	output := Exec(&config)
	if !strings.Contains(output, testFile) {
		t.Fatalf("ls output does not contain test file\n%s", output)
	}
}

func TestLs(t *testing.T) {
	res := Ls()
	if len(res) == 0 {
		t.Fatal()
	}
	t.Log(res)
}
