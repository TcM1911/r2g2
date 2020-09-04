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
 * Copyright (C) Joakim Kennedy, 2020
 */

package r2g2

import (
	"encoding/json"
	"fmt"
)

const (
	getClipboardCMD = "yj"
)

// Clipboard is a representation of radare's clipboard.
type Clipboard struct {
	// Address is the address the bytes were yanked from.
	Address int64 `json:"addr"`
	// Bytes are the bytes yanked.
	Bytes string `json:"bytes"`
}

// GetClipboard returns the value in radare's clipboard.
func (c *Client) GetClipboard() (*Clipboard, error) {
	out, err := c.Run(getClipboardCMD)
	if err != nil {
		return nil, fmt.Errorf("command for getting clipboard failed: %w", err)
	}
	var cb *Clipboard
	if err = json.Unmarshal(out, &cb); err != nil {
		return nil, fmt.Errorf("failed to parse clipboard data: %w", err)
	}
	return cb, nil
}
