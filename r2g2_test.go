//+build integration

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

package r2g2_test

import (
	"path/filepath"
	"testing"

	"github.com/TcM1911/r2g2"
	"github.com/matryer/is"
)

const (
	lsFile = "ls-golden"
)

func TestAnalizeAllFunctions(t *testing.T) {
	is := is.New(t)
	c := setupClient(lsFile, is)

	err := c.AnalyzeAll()
	is.NoErr(err) // Analyzing all functions shouldn't fail.
}

func TestGetOffsetOfSymbol(t *testing.T) {
	is := is.New(t)
	c := setupClient(lsFile, is)

	err := c.AnalyzeAll()
	is.NoErr(err) // Analyzing all functions shouldn't fail.

	offset, err := c.GetSymbolOffset("main")
	is.NoErr(err)
	is.Equal(offset, uint64(0x100001060)) // Correct offset of main
}

func setupClient(name string, is *is.I) *r2g2.Client {
	fp, err := getGoldenFilePath(name)
	is.NoErr(err)

	client, err := r2g2.New(fp)
	is.NoErr(err)
	is.True(client != nil)

	if client == nil {
		is.Fail()
	}
	return client
}

func getGoldenFilePath(name string) (string, error) {
	return filepath.Abs(filepath.Join("testresources", name))
}
