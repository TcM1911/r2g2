package r2g2

import (
	"encoding/json"
	"errors"
)

/*
[{"raised":false,"fd":3,"uri":"/bin/ls","from":0,"writable":false,"size":137640},{"raised":true,"fd":4,"uri":"null://4848","from":0,"writable":true,"size":4848}]
*/

const (
	listOpenedFilesCMD = "oj"
)

var (
	// ErrNoActiveFile is returned if no active file is found.
	ErrNoActiveFile = errors.New("no active file")
)

type OpenFile struct {
	Active         bool   `json:"raised"`
	FileDescriptor int    `json:"fd"`
	Path           string `json:"uri"`
	From           int    `json:"from"`
	Writable       bool   `json:"writable"`
	Size           uint64 `json:"size"`
}

// ListOpenFiles returns a list of open files.
func (c *Client) ListOpenFiles() ([]*OpenFile, error) {
	data, err := c.Run(listOpenedFilesCMD)
	if err != nil {
		return nil, err
	}
	var fileList []*OpenFile
	err = json.Unmarshal(data, &fileList)
	return fileList, err
}

// GetActiveOpenFile returns the current active file in radare.
func (c *Client) GetActiveOpenFile() (*OpenFile, error) {
	fileList, err := c.ListOpenFiles()
	if err != nil {
		return nil, err
	}
	for _, v := range fileList {
		if v.Active {
			return v, err
		}
	}
	return nil, ErrNoActiveFile
}
