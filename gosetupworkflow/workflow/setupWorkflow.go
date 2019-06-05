package workflow

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

//Setup function to setup the worfklow process
func Setup(bpmnDirectory string) {
	files, err := getBPMNFiles(bpmnDirectory)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var waitUploads sync.WaitGroup

	waitUploads.Add(len(files))
	defer waitUploads.Wait()

	for _, file := range files {
		go uploadWorkflows(file, &waitUploads)

	}

}

func getBPMNFiles(dir string) ([]string, error) {
	dirFd, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	files, err := dirFd.Readdir(-1)
	dirFd.Close()

	if err != nil {
		return nil, err
	}

	listOfFiles := make([]string, len(files), cap(files))
	for i, file := range files {
		listOfFiles[i] = dir + "/" + file.Name()
	}

	return listOfFiles, nil
}

func uploadWorkflows(file string, await *sync.WaitGroup) {
	fmt.Println("Uploading file", file)
	defer await.Done()
	extraParams := map[string]string{
		"taskService": "TMS",
	}
	request, err := newfileUploadRequest("http://localhost:7234/workflows/v1/definitions/", extraParams, "bpmnFile", file)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
	}

}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
