/*
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) Joakim Kennedy, 2019-2020
 */

package r2g2

import (
	"fmt"
)

const (
	analyzeFunctionCMD          = "af %s @ %d"
	analyzeFunctionRecursiveCMD = "afr %s @ %d"
	analyzeAllCMD               = "aa"
	analyzeAllAndAutorenameCMD  = "aaa"
)

// AnalyzeFunction analyzing the function starting at the address addr.
func (c *Client) AnalyzeFunction(name string, addr uint64) error {
	_, err := c.Run(fmt.Sprintf(analyzeFunctionCMD, name, addr))
	return err
}

// AnalyzeFunctionRecursive analyzing the function recursivly starting at the address addr.
func (c *Client) AnalyzeFunctionRecursive(name string, addr uint64) error {
	_, err := c.Run(fmt.Sprintf(analyzeFunctionRecursiveCMD, name, addr))
	return err
}

// AnalyzeAll performs the "aaa" command.
func (c *Client) AnalyzeAll() error {
	_, err := c.Run(analyzeAllAndAutorenameCMD)
	return err
}
