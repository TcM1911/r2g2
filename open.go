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
	"encoding/json"
	"errors"
)

const (
	listOpenedFilesCMD = "oj"
)

var (
	// ErrNoActiveFile is returned if no active file is found.
	ErrNoActiveFile = errors.New("no active file")
)

// OpenFile holds information about an open file.
type OpenFile struct {
	Active         bool   `json:"raised"`
	FileDescriptor int    `json:"fd"`
	Path           string `json:"uri"`
	From           int    `json:"from"`
	Writable       bool   `json:"writable"`
	Size           uint64 `json:"size"`
}

// ListOpenFiles returns a list of open files.
func (c *Client) ListOpenFiles() ([]*OpenFile, error) {
	data, err := c.Run(listOpenedFilesCMD)
	if err != nil {
		return nil, err
	}
	var fileList []*OpenFile
	err = json.Unmarshal(data, &fileList)
	return fileList, err
}

// GetActiveOpenFile returns the current active file in radare.
func (c *Client) GetActiveOpenFile() (*OpenFile, error) {
	fileList, err := c.ListOpenFiles()
	if err != nil {
		return nil, err
	}
	for _, v := range fileList {
		if v.Active {
			return v, err
		}
	}
	return nil, ErrNoActiveFile
}
