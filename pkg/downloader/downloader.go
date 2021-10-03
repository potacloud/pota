package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func createImageFile(uri string, dir string) (*os.File, string, error) {
	imagePath := path.Join(dir, path.Base(uri))
	out, err := os.Create(imagePath)
	if err != nil {
		return nil, "", err
	}

	return out, imagePath, nil
}

func getImageHttp(uri string) (*http.Response, error) {
	// download the image from url
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	// if status not equal to 200, return bad status
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("bad status: %s", resp.Status)
	}

	return resp, nil
}

func deepCopyImage(resp *http.Response, writer *os.File) error {
	// copy the downloaded response body image to file created before
	_, err := io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}
	writer.Close()
	resp.Body.Close()
	return nil
}

func getFileSize(uri string, dir string) (int64, error) {
	fs, err := os.Stat(path.Join(dir, path.Base(uri)))
	if err != nil {
		return 0, err
	}

	return fs.Size(), nil
}

func DownloadImage(uri string, dir string) (size int64, path string, err error) {
	writer, path, err := createImageFile(uri, dir)
	if err != nil {
		return 0, "", err
	}

	resp, err := getImageHttp(uri)
	if err != nil {
		return 0, "", err
	}

	err = deepCopyImage(resp, writer)
	if err != nil {
		return 0, "", err
	}

	size, err = getFileSize(uri, dir)
	if err != nil {
		return 0, "", err
	}

	return size, path, nil
}
