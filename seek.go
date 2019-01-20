package r2g2

import (
	"strconv"
)

const (
	seekCMD   = "s"
	seekToCMD = seekCMD + space
)

// GetCurrentAddress returns the current address.
func (c *Client) GetCurrentAddress() (uint64, error) {
	data, err := c.Run(seekCMD)
	if err != nil {
		return 0, err
	}
	addr, err := stringToAddr(string(data))
	if err != nil {
		return 0, err
	}
	return addr, err
}

// SeekTo seeks to address addr.
func (c *Client) SeekTo(addr uint64) error {
	addrStr := addrToString(addr)
	_, err := c.Run(seekCMD + addrStr)
	return err
}

func addrToString(addr uint64) string {
	return strconv.FormatUint(addr, 16)
}

func stringToAddr(str string) (uint64, error) {
	return strconv.ParseUint(str, 16, 64)
}
