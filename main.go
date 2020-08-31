package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"time"
)

var (
	extentions = map[string]string{
		"py":   "python3",
		"js":   "node",
		"ex":   "elixir",
		"cpp":  "gcc",
		"go":   "go",
		"java": "java",
	}

	testCases []testCase
)

type testCase struct {
	Args []string `json:"args"`
	Truth string `json:"truth"`
}

func init() {
	jsonFile, err := os.Open("sample_test_case.json")
	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan() {
		var test testCase

		json.Unmarshal([]byte(scanner.Text()), &test)
		testCases = append(testCases, test)
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(testCases)
	fmt.Println(len(testCases))
}

func main() {
	cwd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(fmt.Sprintf("%s/submissions", cwd))
	for _, file := range files {
		fn := fmt.Sprintf("submissions/%s", file.Name())
		cmnd := GetFileCommand(file.Name())

		var failed bool

		start := time.Now()

		for _, testcase := range testCases {
			c := exec.Command(cmnd)

			if cmnd == "go" {
				c.Args = append(c.Args, "run")
				c.Args = append(c.Args, fn)
			} else {
				c.Args = append(c.Args, fn)
			}


			args := testcase.Args
			truth := testcase.Truth

			for _, arg := range args {
				c.Args = append(c.Args, arg)
			}

			out, err := c.Output()

			if err != nil {
				fmt.Println(err)
			}

			resp := strings.ToLower(strings.Replace(string(out), "\n", "", -1))

			if resp != truth {
				failed = true
				fmt.Printf("testcase: %v truth: %s found: %s\n", args, truth, resp)
				break
			}
		}

		if failed == true {
			fmt.Printf("%v failed!	 %v\n", fn, time.Now().Sub(start))
		} else {
			fmt.Printf("%v correct!  %v\n", fn, time.Now().Sub(start))
		}
	}
}

func GetFileCommand(filename string) string {
	s := strings.Split(filename, ".")
	return extentions[s[len(s)-1]]
}
