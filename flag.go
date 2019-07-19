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
	cmd := fmt.Sprintf("%s \"%s\" %d @ %d", flagCMD, name, length, offset)
	_, err := c.Run(cmd)
	return err
}
