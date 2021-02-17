package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"errors"
	"os"
)

const dropboxContentBaseUrl = "https://content.dropboxapi.com/2/"
const dropboxFileDownloadUrl = dropboxContentBaseUrl + "files/download"

const dropboxAccessTokenEnvVar = "DROPBOX_ACCESS_TOKEN"

// DownloadFile retrieves a file from dropbox at the given path
// If is the caller's responsibility to close the file stream
func DownloadFile(filePath string) (io.ReadCloser, error) {
	token := os.Getenv(dropboxAccessTokenEnvVar)
	if token == "" {
		return nil, errors.New("Could not download file; " + dropboxAccessTokenEnvVar + " env var not set")
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", dropboxFileDownloadUrl, nil)
	addAuthHeader(token, request)

	arg := fmt.Sprintf(`{"path":"%v"}`, filePath)

	request.Header.Add("Dropbox-API-Arg", arg)

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if (resp.StatusCode != 200) {
		respBody, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, errors.New("Error downloading file: " + string(respBody))
	}

	return resp.Body, nil
}

func addAuthHeader(token string, request *http.Request) {
	request.Header.Add("Authorization", "Bearer " + token)
}