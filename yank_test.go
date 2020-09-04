package r2g2

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestGetClipboard(t *testing.T) {
	is := is.New(t)

	fp, err := getGoldenFilePath("yank.json")
	is.NoErr(err)

	data, err := ioutil.ReadFile(fp)
	is.NoErr(err) // Needs the mocked data.
	data = append(data, '\n')

	out := bytes.NewBuffer(data)
	in := new(bytes.Buffer)
	c := &Client{reader: out, writer: in}

	cb, err := c.GetClipboard()
	is.NoErr(err) // Should not fail
	is.Equal(cb.Address, int64(4294971501))
	is.Equal(cb.Bytes, "4881ec180600004989f74189fe488d85c0fdffff488945d085ff7f05e822310000")
}
