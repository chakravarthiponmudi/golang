package workflow

import (
	"fmt"
	"os"
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
	await.Done()
}
