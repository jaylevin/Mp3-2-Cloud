# Mp3-2-Cloud
A single endpoint REST server written in Go that demonstrates basic mp3 file uploading a cloud storage server using https://github.com/minio/minio-go

# Setup Guide
1) Clone the repository
2) Execute `make` in the project root directory. You may need to `go get ...` some external dependencies.
3) Execute the resulting binary file named `mp32cloud` in the same directory. The REST server should now be running and listening for requests. Ctrl + C out of the server
4) A configuration file named `config.json` should've been generated. Open it and modify the appropriate fields with your DigitalOcean's access key, secret key, and endpoint.
5) Restart the server by executing the binary file named `mp32cloud`. You should be informed that your configuration file has been loaded with the changes that you made in step 4.
6) Open up a new terminal window and navigate to the project's root directory, leave the server running.
7) Execute `go test -mp3Path="/Users/Desktop/Admin/file.mp3"`, where the mp3Path argument is an actual path to an mp3 file on your local machine. This
8) This will perform a PUT request to http://localhost:8080/upload, which should handle the upload process.



# Upload Endpoint Documentation
Method: PUT 
Endpoint: http://localhost:8080/upload
Request body is expected to a be JSON object of the form
```
{
  "mp3Name": "file.mp3",
  "mp3Bytes": BLOB
}
```

The JSON response will always be in the form
```
{
  "success": bool
}
```
If `{"success":true}` the upload was successful.
