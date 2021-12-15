package util

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const CHA = "abcdefghijklmnopqrstuvwxyz0123456789"

// RandString n is the length. a-z 0-9
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = CHA[rand.Intn(len(CHA))]
	}
	return string(b)
}

//Does the file exist
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//Get template source file type
func GetFileType(path string) string {
	index := strings.LastIndex(".", path)
	if index == -1 {
		return ""
	}

	temp := []byte(path)
	return string(temp[index:])
}

func GetBody(client *http.Client, URL string) ([]byte, error) {
	resp, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// PostBody read post body
func PostBody(client *http.Client, URL string, data url.Values) ([]byte, error) {
	resp, err := client.PostForm(URL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
