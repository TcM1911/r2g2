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
	"io"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestGenerateZignatureForSymbol(t *testing.T) {
	is := is.New(t)

	fp, err := getGoldenFilePath("zignature.json")
	is.NoErr(err)

	// Data returned when generating the zignatures.
	data0 := []byte{}

	data, err := ioutil.ReadFile(fp)
	is.NoErr(err)
	data = append(data, '\n', 0x00)

	in := new(bytes.Buffer)
	c := &Client{reader: &chainedResponse{data: [][]byte{data0, data}}, writer: in}

	z, err := c.ZignatureFunction("main")
	is.NoErr(err)
	is.Equal(z.Bytes, "554889e54157415641554154534881ec180600004989f74189fe488d85c0fdffff488945d085ff7f05e822310000488d35fe39000031ffe86a340000bf01000000e80634000085c0745ec705cc51000050000000488d3dd9390000e8bc3300004885c0740f803800740a4889c7e84a330000eb22488d55c8be68740840bf0100000031c0e8bd33000083f8ff740e0fb745ca85c07406890584510000c705e252000001000000eb26c605a151000001488d3d7e390000e8613300004885c0740e4889c7e8f4320000890552510000e85b33000085c07507c605765100000141bc100000004c8d2d51390000488d1d760700004489f74c89fe4c89eae8223300008d48db83f9530f87960300004863048b4801d8ffe0c7055152000001000000ebd1c7055d5200000100000031c0eb25c7057352000001000000c7054152000000000000c6050651000000eba631c089053452000089053a520000890538520000eb90c705f45100000100000031c08905f4510000890536520000e973ffffffc605de50000001e967ffffff31c08905f55100008905fb510000c705f551000001000000e94affffff4183e4ed4183cc02488d3dae380000488d35ae380000e8053200004489e181e1eefbffff84c0440f45e1e91bffffffc7058351000001000000e90cffffffc6057b50000001e900ffffffc605")
}

func TestGenerateZignatureAtOffset(t *testing.T) {
	is := is.New(t)

	fp1, err := getGoldenFilePath("function.json")
	is.NoErr(err)
	fp2, err := getGoldenFilePath("zignature.json")
	is.NoErr(err)

	// Data returned when generating the zignatures.
	data0 := []byte{}

	data1, err := ioutil.ReadFile(fp1)
	is.NoErr(err)
	data1 = append(data1, '\n', 0x00)

	data2, err := ioutil.ReadFile(fp2)
	is.NoErr(err)
	data2 = append(data2, '\n', 0x00)

	in := new(bytes.Buffer)
	c := &Client{reader: &chainedResponse{data: [][]byte{data1, data0, data2}}, writer: in}

	z, err := c.ZignatureFunctionOffset(0x1337)
	is.NoErr(err)
	is.Equal(z.Bytes, "554889e54157415641554154534881ec180600004989f74189fe488d85c0fdffff488945d085ff7f05e822310000488d35fe39000031ffe86a340000bf01000000e80634000085c0745ec705cc51000050000000488d3dd9390000e8bc3300004885c0740f803800740a4889c7e84a330000eb22488d55c8be68740840bf0100000031c0e8bd33000083f8ff740e0fb745ca85c07406890584510000c705e252000001000000eb26c605a151000001488d3d7e390000e8613300004885c0740e4889c7e8f4320000890552510000e85b33000085c07507c605765100000141bc100000004c8d2d51390000488d1d760700004489f74c89fe4c89eae8223300008d48db83f9530f87960300004863048b4801d8ffe0c7055152000001000000ebd1c7055d5200000100000031c0eb25c7057352000001000000c7054152000000000000c6050651000000eba631c089053452000089053a520000890538520000eb90c705f45100000100000031c08905f4510000890536520000e973ffffffc605de50000001e967ffffff31c08905f55100008905fb510000c705f551000001000000e94affffff4183e4ed4183cc02488d3dae380000488d35ae380000e8053200004489e181e1eefbffff84c0440f45e1e91bffffffc7058351000001000000e90cffffffc6057b50000001e900ffffffc605")
}

type chainedResponse struct {
	data   [][]byte
	round  int
	reader io.Reader
}

func (c *chainedResponse) Read(dst []byte) (int, error) {
	if c.round > len(c.data)-1 {
		return 0, io.EOF
	}

	if c.reader == nil {
		c.reader = bytes.NewReader(c.data[0])
	}

	n, err := c.reader.Read(dst)
	if err == io.EOF || n < len(dst) {
		c.round++
		if c.round < len(c.data) {
			c.reader = bytes.NewReader(c.data[c.round])
		}
	}
	return n, err
}
