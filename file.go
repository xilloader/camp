package camp

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func GetLocalFileData(filename string) []byte {
	data, _ := GetLocalFileDataE(filename)
	return data
}

func GetLocalFileDataE(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetOnlineFileData(filename string) []byte {
	data, _ := GetOnlineFileDataE(filename)
	return data
}

func GetOnlineFileDataE(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FileExtension(filename string) string {
	ext := filepath.Ext(filename) // return ".suffix"
	if len(ext) <= 1{
		ext = ""
	} else {
		ext = ext[1:]
	}
	return ext
}
