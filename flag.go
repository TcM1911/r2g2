package r2g2

import "fmt"

const (
	flagCMD = "f"
)

// NewFlag ..
func (c *Client) NewFlag(name string, offset uint64) error {
	return c.NewFlagWithLength(name, offset, uint64(1))
}

// NewFlagWithLength  ..
func (c *Client) NewFlagWithLength(name string, offset, length uint64) error {
	_, err := c.Run(fmt.Sprintf("'%s %s %d @ %d'", flagCMD, name, length, offset))
	return err
}
