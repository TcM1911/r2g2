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
