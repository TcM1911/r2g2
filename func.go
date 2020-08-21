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
	"fmt"
)

const (
	getFunctionDetailCMD   = "pdfj"
	getFunctionAtOffsetCMD = "pdfj @ %d"
)

// Function holds information about a radare function.
type Function struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`
	Addr uint64 `json:"addr"`
	Ops  []*Op  `json:"ops"`
}

// Op holds information about an operation.
type Op struct {
	Offset   uint64   `json:"offset"`
	Ptr      uint64   `json:"ptr"`
	ESIL     string   `json:"esil"`
	RefPtr   bool     `json:"refptr"`
	FcnAddr  uint64   `json:"fcn_addr"`
	FcnLast  uint64   `json:"fcn_last"`
	Size     uint64   `json:"size"`
	OpCode   string   `json:"opcode"`
	Disasm   string   `json:"disasm"`
	Bytes    string   `json:"bytes"`
	Family   string   `json:"family"`
	Type     string   `json:"type"`
	TypeNum  uint64   `json:"type_num"`
	Type2Num uint64   `json:"type2_num"`
	Flags    []string `json:"flags"`
	Jump     uint64   `json:"jump"`
	Fail     uint64   `json:"fail"`
	XRefs    []*XRef  `json:"xrefs"`
}

// XRef holds information about a cross-reference.
type XRef struct {
	Addr uint64 `json:"addr"`
	Type string `jsoin:"type"`
}

// GetCurrentFunction returns information about the current function.
func (c *Client) GetCurrentFunction() (*Function, error) {
	data, err := c.Run(getFunctionDetailCMD)
	if err != nil {
		return nil, err
	}
	return parseFunctionResp(data)
}

// GetFunctionAtOffset returns function information about a function
// located at the given offset.
func (c *Client) GetFunctionAtOffset(offset uint64) (*Function, error) {
	data, err := c.Run(fmt.Sprintf(getFunctionAtOffsetCMD, offset))
	if err != nil {
		return nil, err
	}
	return parseFunctionResp(data)
}

func parseFunctionResp(data []byte) (*Function, error) {
	var function Function
	err := json.Unmarshal(data, &function)
	return &function, err
}
