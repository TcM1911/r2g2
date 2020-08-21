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
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestListOpenFiles(t *testing.T) {
	is := is.New(t)

	fp, err := getGoldenFilePath("open_files.json")
	is.NoErr(err)

	data, err := ioutil.ReadFile(fp)
	is.NoErr(err)
	data = append(data, '\n')

	out := bytes.NewBuffer(data)
	in := new(bytes.Buffer)
	c := &Client{reader: out, writer: in}

	files, err := c.ListOpenFiles()
	is.NoErr(err)
	is.Equal(len(files), 1)
}

func TestActiveFiles(t *testing.T) {
	is := is.New(t)

	fp, err := getGoldenFilePath("open_files.json")
	is.NoErr(err)

	data, err := ioutil.ReadFile(fp)
	is.NoErr(err)
	data = append(data, '\n')

	out := bytes.NewBuffer(data)
	in := new(bytes.Buffer)
	c := &Client{reader: out, writer: in}

	f, err := c.GetActiveOpenFile()
	is.NoErr(err)
	is.Equal(f.Path, "/bin/ls")
	is.Equal(f.Size, uint64(51888))
}
