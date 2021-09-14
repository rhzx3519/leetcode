package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestTempFile(t *testing.T) {
	tempDir := "/tmp/test_temp/"
	f, err := ioutil.TempFile(tempDir, "temp_*")
	if err != nil {
		log.Fatalln(err.Error())
	}
	os.Rename(f.Name(), strings.Join([]string{"/tmp/test_temp/", "renamed_temp_file"}, ""))
	f.Close()
}
