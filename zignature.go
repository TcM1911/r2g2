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
	generateZignatureForFunctionCMD     = "zaf"
	generateZignatureForAllFunctionsCMD = "zaF"
	getZignaturesCMD                    = "zj"
	removeFunctionZignatureCMD          = "z-"
)

// Zignature is a signature type used by Radare2
type Zignature struct {
	Name     string            `json:"name"`
	Bytes    string            `json:"bytes"`
	Mask     string            `json:"mask"`
	Graph    Graph             `json:"graph"`
	Addr     uint64            `json:"addr"`
	RealName string            `json:"realname"`
	Refs     []string          `json:"refs"`
	XRefs    []string          `json:"xrefs"`
	Vars     []string          `json:"vars"`
	Hash     map[string]string `json:"hash"`
}

// Graph is a graph summary of a function.
type Graph struct {
	CC    int `json:"cc"`
	Nbbs  int `json:"nbbs"`
	Edges int `json:"edges"`
	Ebbs  int `json:"ebbs"`
	BBSum int `json:"bbsum"`
}

// ZignatureFunction creates a Zignature for a function at the given symbol.
func (c *Client) ZignatureFunction(symbol string) (*Zignature, error) {
	return genFuncZig(c, symbol)
}

// ZignatureFunctionOffset creates a Zignature for a function at the given offset.
func (c *Client) ZignatureFunctionOffset(offset uint64) (*Zignature, error) {
	f, err := c.GetFunctionAtOffset(offset)
	if err != nil {
		return nil, fmt.Errorf("faild to get function information at offset %d: %w", offset, err)
	}
	return genFuncZig(c, f.Name)
}

func genFuncZig(c *Client, symbol string) (*Zignature, error) {
	data, err := c.Run(fmt.Sprintf(
		"%s %s r2g2zig-%s",
		generateZignatureForFunctionCMD, symbol, symbol))
	if err != nil {
		return nil, fmt.Errorf("failed to generate zignature for \"%s\": %w", symbol, err)
	}
	var buf []json.RawMessage
	err = json.Unmarshal(data, &buf)
	if err != nil {
		return nil, fmt.Errorf("parsing json for \"%s\" failed: %w", symbol, err)
	}

	var z *Zignature
	zigName := fmt.Sprintf("r2g2zig-%s", symbol)
	for _, v := range buf {
		err = json.Unmarshal(v, &z)
		if err != nil {
			return nil, fmt.Errorf("parsing zignature for \"%s\" failed: %w", symbol, err)
		}
		if z.Name == zigName {
			break
		} else {
			z = nil
		}
	}
	if z == nil {
		return nil, fmt.Errorf("no matching zignature found")
	}

	_, err = c.Run(fmt.Sprintf("%s r2g2zig-%s", removeFunctionZignatureCMD, symbol))
	if err != nil {
		return nil, fmt.Errorf("error when cleaning up generated zignature: %w", err)
	}
	return z, nil
}
