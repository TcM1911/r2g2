package r2g2

import "bufio"
import "io"

// CommandHandler can send commands to radare and return the output.
type CommandHandler interface {
	Run(cmd string) ([]byte, error)
}

// Client is a client for interacting with radare2.
type Client struct {
	writer io.Writer
	reader io.Reader
}

// Run executes a radare command and returns the result as a byte slice.
func (c *Client) Run(cmd string) ([]byte, error) {
	if _, err := c.writer.Write([]byte(cmd + "\n")); err != nil {
		return nil, err
	}
	data, err := bufio.NewReader(c.reader).ReadBytes(0x00)
	if err == io.EOF {
		return data[:len(data)-1], nil
	}
	return data[:len(data)-1], err
}
