package log

import (
	"fmt"
	"testing"
)

func TestAddElement(t *testing.T) {
	logHtml := &LogHtml{}
	logHtml.Create("C:\\workspace\\GoProject\\gotools\\src\\gotools\\util\\log\\View\\base.html")
	logHtml.AddOneLog("test")
	logs := []string{}
	logs = append(logs, "test1")
	logs = append(logs, "test2")
	logHtml.AddLogs(logs)
	res, _ := logHtml.GetString()
	fmt.Println(res)
	logHtml.GetFile("C:\\workspace\\GoProject\\gotools\\src\\gotools\\util\\log\\View\\baseTest.html")

}
