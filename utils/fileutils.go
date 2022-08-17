package utils

import (
	"encoding/json"
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

func FindDecoderTargetPosition(dec *json.Decoder, jsonkey string) error {
	for dec.More() {
		// Token returns the next JSON token in the input stream.
		t, err := dec.Token()
		if err != nil {
			return err
		}
		if t == jsonkey {
			_, err = dec.Token()
			return err
		}
	}
	return nil
}
