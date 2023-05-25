package network

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFile(link string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	u, _ := url.Parse(strings.ReplaceAll(link, "raw/gh-pages/", ""))
	os.MkdirAll(filepath.Join("data", filepath.Dir(u.Path)), os.ModePerm)

	err = os.WriteFile(filepath.Join("data", u.Path), body, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Println("File saved successfully:", strings.TrimLeft(u.Path, "/"))
	return nil
}
