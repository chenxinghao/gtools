package main

import (
	"github.com/chenxinghao/gtools/util/ast"
	"github.com/chenxinghao/gtools/util/file"
	"strings"
	"time"
)

func main() {

	insertFuncName := "Tst"
	annoName := "Annotation"
	insertFuncMap := make(map[string][]string)
	insertFuncMap[insertFuncName] = []string{}

	var fop FileUtils.FolderOP
	dpath, _ := fop.GetCurrentDirectory()
	dirPath := fop.GetParentDirectory(dpath)

	dpath = "D:/work/GO_PROJECT/src/test/main"
	dirPath = "D:/work/GO_PROJECT/src/test"

	tempStrs := strings.Split(dirPath, "/")
	importStr := tempStrs[len(tempStrs)-1] + "/AnnotationsFunc"

	ast.CheckFunc(dirPath+"/AnnotationsFunc", insertFuncMap)

	ast.WalkAndHandler(dirPath, annoName, importStr, insertFuncName, insertFuncMap)

	time.Sleep(time.Duration(20) * time.Second)
}
