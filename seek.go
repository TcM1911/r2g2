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
 * Copyright (C) Joakim Kennedy, 2019
 */

package r2g2

import (
	"strconv"
)

var (
	seekCMD   = "s"
	space     = " "
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
