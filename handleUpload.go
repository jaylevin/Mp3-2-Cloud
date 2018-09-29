package main

import (
	"net/http"
	"log"
	"encoding/json"
	"bytes"

	/* Internal */
	"github.com/jaylevin/Mp3-2-Cloud/storage"
)

type Mp3UploadRequest struct {
	Mp3Name  string // should end with .mp3
	Mp3Bytes []byte
}

// the digital ocean's "space" you want to upload the incoming mp3 files.
var bucketName = "tmn"

func handleMp3Upload(resp http.ResponseWriter, req *http.Request) {
	var request Mp3UploadRequest

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println("Failed to decode mp3 upload request:", err.Error())
		resp.Write([]byte(`{"success":false}`))
		return
	}
	defer req.Body.Close()

	client := storage.NewClient(conf)
	_, err = client.UploadMp3(bucketName, request.Mp3Name, bytes.NewReader(request.Mp3Bytes), int64(len(request.Mp3Bytes)))
	if err != nil {
		log.Println("Track upload failed:", err.Error())
		resp.Write([]byte(`{"success":false}`))
		return
	}

	resp.Write([]byte(`{"success":true}`))
}
