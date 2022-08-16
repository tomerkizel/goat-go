package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var (
	//temp file prefix
	tempPrefix = "jfrog.cli.temp."
	//temp file root dir
	tempDirBase string
)

//init the tempDirBase to the OS temp directory
func init() {
	tempDirBase = os.TempDir()
}

func CreateTempFile() (*os.File, error) {
	if tempDirBase == "" {
		return nil, fmt.Errorf("temp file cannot be created in an empty base dir")
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fd, err := ioutil.TempFile(tempDirBase, tempPrefix+"-"+timestamp+"-")
	return fd, err
}
