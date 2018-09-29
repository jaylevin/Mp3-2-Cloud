package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	/* Vendors */
	"github.com/logrusorgru/aurora"
	"github.com/minio/minio-go"

	/* Internal */
	"github.com/jaylevin/Mp32Cloud/storage"
)

/* The absolute path to the mp3 file to upload. */
var mp3Path = flag.String("mp3Path", "", "absolute path to a mp3 file to conduct the test (upload) with")

func TestUploadTrack(t *testing.T) {
	if *mp3Path == "" {
		printError("Mp3 path argument missing!")
		printError(`Ex: go test -mp3Path="/Users/Admin/Desktop/file.mp3""`)
		t.Fail()
		return
	}

	// If these two variables do not equal each other at the end of the test, then FAIL.
	var originalFileSize int
	var uploadedFileSize int

	/* Do some tricky string/slice stuff to get the fileName and root directory it is located in.*/
	split := strings.Split(*mp3Path, "/")
	fileName := split[len(split)-1]
	split = split[:len(split)-1]

	log.Println("Attempting to upload mp3 file to cloud:", fileName)
	log.Println("Located in directory:", strings.Join(split, "/"))

	if !strings.HasSuffix(fileName, ".mp3") {
		printError("Mp3 path must end with .mp3!")
		printError(`Ex: go test -mp3Path="/Users/Admin/Desktop/file.mp3""`)
		t.Fail()
		return
	}

	mp3Bytes, err := ioutil.ReadFile(*mp3Path)
	if err != nil {
		t.Error(err)
		return
	}

	uploadRequest := Mp3UploadRequest{
		Mp3Name:  fileName,
		Mp3Bytes: mp3Bytes,
	}

	// Marshal the above struct instance into json bytes
	b, err := json.Marshal(uploadRequest)
	if err != nil {
		t.Error(err)
		return
	}

	/* Make the HTTP PUT Request to TMN-API */
	printMagenta("Sending mp3 file bytes to cloud server...")

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/upload", bytes.NewReader(b))
	httpResp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	/* Check for successful json response from server */
	var respBody map[string]interface{}
	err = json.NewDecoder(httpResp.Body).Decode(&respBody)

	success := respBody["success"]
	if !success.(bool) {
		log.Println(aurora.BgRed(aurora.Red("Mp3 file upload was unsuccessful, server returned success:false")))
		return
	}

	// Retrieve the file object from the cloud server to perform backwards checking
	minioClient := storage.NewClient()
	object, err := minioClient.GetObject("tmn", fileName, minio.GetObjectOptions{})
	if err != nil {
		t.Error(err)
		return
	}
	stat, _ := object.Stat()

	// Check if the original file size matches the uploaded file's size
	originalFileSize = len(mp3Bytes)
	uploadedFileSize = int(stat.Size)
	if originalFileSize != uploadedFileSize {
		printError("Test failed! Original file size != uploaded file size")
		t.Fail()
	} else {
		// Everything seems to have went smoothly
		printGreen("Mp3 file upload was successful!")

	}
}

func printGreen(args ...interface{}) {
	log.Println(aurora.Green(args))
}

func printMagenta(args ...interface{}) {
	log.Println(aurora.Magenta(args))
}

func printError(args ...interface{}) {
	log.Println(aurora.Red(args))
}
