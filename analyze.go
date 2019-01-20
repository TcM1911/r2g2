package r2g2

import (
	"fmt"
)

const (
	analyzeFunctionCMD         = "af"
	analyzeAllCMD              = "aa"
	analyzeAllAndAutorenameCMD = "aaa"
)

// AnalyzeFunction analyzing the function starting at the address addr.
func (c *Client) AnalyzeFunction(name string, addr uint64) error {
	_, err := c.Run(fmt.Sprintf("af %s @ %d", name, addr))
	return err
}
