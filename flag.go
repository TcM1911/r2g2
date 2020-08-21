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
	"encoding/json"
	"fmt"
)

const (
	flagCMD               = "f \"%s\" %d @ %d"
	flagRemoveCMD         = "f-%s"
	flagRemoveAtOffsetCMD = "f-@%d"
	flagListAllCMD        = "fj"
	flagRenameCMD         = "fr %s %s" // fr [[old]] [new]
)

// Flag is a representation of an radare2 flag.
type Flag struct {
	Name     string `json:"name"`
	RealName string `json:"realname"`
	Size     uint64 `json:"size"`
	Offset   uint64 `json:"offset"`
}

// NewFlag creates a flag with the name at the offset with a length of 1.
func (c *Client) NewFlag(name string, offset uint64) error {
	return c.NewFlagWithLength(name, offset, uint64(1))
}

// NewFlagWithLength creates a flag with the name at the offset with the length.
func (c *Client) NewFlagWithLength(name string, offset, length uint64) error {
	cmd := fmt.Sprintf(flagCMD, name, length, offset)
	_, err := c.Run(cmd)
	return err
}

// GetFlags returns all the flags in the current session.
func (c *Client) GetFlags() ([]*Flag, error) {
	buf, err := c.Run(fmt.Sprintf(flagListAllCMD))
	if err != nil {
		return nil, err
	}

	var jsonData []json.RawMessage
	err = json.Unmarshal(buf, &jsonData)
	if err != nil {
		return nil, err
	}

	flags := make([]*Flag, 0, len(jsonData))
	for _, v := range jsonData {
		var f *Flag
		err = json.Unmarshal(v, &f)
		if err != nil {
			return nil, err
		}
		flags = append(flags, f)
	}

	return flags, nil
}
