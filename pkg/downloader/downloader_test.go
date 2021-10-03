package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateImageFile(t *testing.T) {
	_, imagePath, err := createImageFile("https://ikhsanputra.com/assets/images/ava.png", "../../pota/lib/images")
	assert.NoError(t, err)
	assert.NotEmpty(t, imagePath)
}

func TestGetImageHttp(t *testing.T) {
	resp, err := getImageHttp("https://ikhsanputra.com/assets/images/ava.png")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeepCopyImage(t *testing.T) {
	uri := "https://ikhsanputra.com/assets/images/ava.png"

	resp, err := getImageHttp(uri)
	assert.NoError(t, err)

	writer, _, err := createImageFile(uri, "../../pota/lib/images")
	assert.NoError(t, err)

	err = deepCopyImage(resp, writer)
	assert.NoError(t, err)
}

func TestGetFileSize(t *testing.T) {
	size, err := getFileSize("https://ikhsanputra.com/assets/images/ava.png", "../../pota/lib/images")
	assert.NoError(t, err)
	assert.Greater(t, size, int64(0))
}

func TestDownloadImage(t *testing.T) {
	size, path, err := DownloadImage("https://ikhsanputra.com/assets/images/ava.png", "../../pota/lib/images")
	assert.NoError(t, err)
	assert.NotEmpty(t, path)
	assert.Greater(t, size, int64(0))
}
