package FileUtils

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

type Info struct {
}

//第一个bool指返回的文件是否存在，第二个bool值返回路径是否是文件夹
func (i *Info) IsFileExist(path string) (bool, bool) {
	info, err := os.Stat(path)
	if err != nil {
		return false, false
	}
	return true, info.IsDir()
}

func (i *Info) GetSystemNewlineChar() string {
	fmt.Println(runtime.GOOS)
	systemName := runtime.GOOS
	switch systemName {
	case "windows":
		return "\r\n"
	case "linux":
		return "\n"
	default:
		return "\n"
	}

}

func (i *Info) GetSystemFilePathDelimiter() string {
	fmt.Println(runtime.GOOS)
	systemName := runtime.GOOS
	switch systemName {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	default:
		return "/"
	}

}
func (i *Info) IsEmptyDir(path string) bool {
	dir, _ := ioutil.ReadDir(path)
	return len(dir) == 0
}
