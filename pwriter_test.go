package goat

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type outputRecord struct {
	IntKey  int    `json:"intKey"`
	StrKey  string `json:"strKey"`
	BoolKey bool   `json:"boolKey"`
}

var records = []outputRecord{
	{1, "1", true},
	{2, "2", false},
	{3, "3", true},
	{4, "4", false},
	{5, "5", false},
	{6, "6", true},
	{7, "7", true},
	{8, "8", true},
	{9, "9", false},
	{10, "10", false},
	{11, "11", false},
	{12, "12", true},
	{13, "13", false},
	{14, "14", true},
	{15, "15", true},
	{16, "16", true},
	{17, "17", false},
	{18, "18", true},
	{19, "19", false},
	{20, "20", false},
	{21, "21", true},
	{22, "22", true},
	{23, "23", true},
	{24, "24", false},
	{25, "25", false},
	{26, "26", false},
	{27, "27", true},
	{28, "28", false},
	{29, "29", true},
	{30, "30", true},
}

type Response struct {
	Arr []outputRecord `json:"arr"`
}

func writeTestRecords(t *testing.T, pw *PersistentWriter) {
	var sendersWaiter sync.WaitGroup
	for i := 0; i < len(records); i += 3 {
		sendersWaiter.Add(1)
		go func(start, end int) {
			defer sendersWaiter.Done()
			for j := start; j < end; j++ {
				pw.write(records[j])
			}
		}(i, i+3)
	}
	sendersWaiter.Wait()
	assert.NoError(t, pw.close())
}

func TestPersistentWriter(t *testing.T) {
	writer, err := newPersistentWriter("arr", true, 50000)
	assert.NoError(t, err)
	writeTestRecords(t, writer)
	of, err := os.Open(writer.GetFilePath())
	assert.NoError(t, err)
	byteValue, _ := ioutil.ReadAll(of)
	var response Response
	assert.NoError(t, json.Unmarshal(byteValue, &response))
	assert.NoError(t, of.Close())
	assert.NoError(t, writer.removeOutputFilePath())
	for i := range records {
		assert.Contains(t, response.Arr, records[i], "record %s missing", records[i].StrKey)
	}
}
