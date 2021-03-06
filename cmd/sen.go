package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

var (
	countTestFailed int
	countTestPassed int
)

// Sen defines object stores each row in csv
type Sen struct {
	Name     string
	Method   string
	EndPoint string
	Headers  string
	Body     string
	Response string
	Ignore   []string
}

// Flags defines flags of cli
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "Input file CSV",
		Value: "input",
		Usage: "Input file CSV with given format",
	},
}

// Action defines the main action for glod-cli
func Action(c *cli.Context) {
	if len(c.Args()) <= 0 {
		cli.ShowAppHelp(c)
		return
	}

	csvDirectory := c.Args()[0]
	listData, err := readCSV(csvDirectory)
	if err != nil {
		fmt.Println("Cannot read csv")
		os.Exit(1)
	}
	runningTest(listData)
	report()
}

func readCSV(fileName string) ([]Sen, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var listData = []Sen{}
	for i, each := range rawCSVdata {
		if i == 0 {
			continue
		}
		var sen Sen
		sen.Name = each[0]
		sen.Method = each[1]
		sen.EndPoint = each[2]
		sen.Headers = each[3]
		sen.Body = each[4]
		sen.Response = each[5]
		sen.Ignore = strings.Split(each[6], ",")
		listData = append(listData, sen)
	}
	return listData, nil
}

func runningTest(sens []Sen) {
	for _, each := range sens {

		if each.Name == "" {
			continue
		}

		fmt.Println("Running test case: " + each.Name)
		req, err := http.NewRequest(each.Method, each.EndPoint, bytes.NewBuffer([]byte(each.Body)))

		if err != nil {
			fmt.Println("Cannot make request")
			countTestFailed++
			return
		}
		if each.Headers != "" {
			keyValue := strings.Split(each.Headers, ":")
			req.Header.Add(keyValue[0], keyValue[1])
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Cannot make request")
			countTestFailed++
			return
		}
		defer resp.Body.Close()

		buffer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Cannot read body")
			countTestFailed++
			return
		}
		switch each.Method {
		case "GET":
			if sanitizeResponse(each.Response) == sanitizeResponse(string(buffer)) {
				fmt.Println("Test passed")
				countTestPassed++
			} else {
				fmt.Println("Test failed")
				countTestFailed++
			}
		case "POST":
			var expectedResponseMap map[string]interface{}
			err := json.Unmarshal([]byte(each.Response), &expectedResponseMap)
			if err != nil {
				fmt.Println("Cannot unmarshal response")
				countTestFailed++
				return
			}

			var actualResponseMap map[string]interface{}
			err = json.Unmarshal(buffer, &actualResponseMap)
			if err != nil {
				fmt.Println("Cannot unmarshal response")
				countTestFailed++
				return
			}

			isSame := true

			for k, _ := range expectedResponseMap {
				if checkSlicesContainValue(each.Ignore, k) {
					continue
				}

				if _, ok := actualResponseMap[k]; !ok {
					isSame = false
					break
				}

				if !compareUnknownObject(actualResponseMap[k], expectedResponseMap[k]) {
					isSame = false
					break
				}
			}

			if isSame {
				fmt.Println("Test passed")
				countTestPassed++
			} else {
				fmt.Println("Test failed")
				countTestFailed++
			}
		}
	}
}

func report() {
	fmt.Println("---------------------------------------")
	fmt.Println("Number test failed: " + strconv.Itoa(countTestFailed))
	fmt.Println("Number test passed: " + strconv.Itoa(countTestPassed))

}
