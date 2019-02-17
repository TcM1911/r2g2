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
	"errors"
	"os"
	"strconv"
)

const (
	envPipeIn  = "R2PIPE_IN"
	envPipeOut = "R2PIPE_OUT"
)

var (
	ErrR2PipeNotAvailable  = errors.New("R2PIPEs are not available")
	ErrPipeOutNotAvailable = errors.New("Could not connect to Pipe out FD")
	ErrPipeInNotAvailable  = errors.New("Could not connect to Pipe in FD")
)

// CheckForR2Pipe returns true if the Environment variables R2PIPE_IN and
// R2PIPE_OUT has been set. This indicates that the process has been invoked
// from within r2 via "#!pipe".
func CheckForR2Pipe() bool {
	return (os.Getenv(envPipeIn) != "") && (os.Getenv(envPipeOut) != "")
}

func getFd(file string) (*os.File, error) {
	fd, err := strconv.Atoi(os.Getenv(file))
	if err != nil {
		return nil, err
	}
	return os.NewFile(uintptr(fd), file), nil
}

func OpenPipe() (*Client, error) {
	if !CheckForR2Pipe() {
		return nil, ErrR2PipeNotAvailable
	}
	in, err := getFd(envPipeIn)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, ErrPipeInNotAvailable
	}
	out, err := getFd(envPipeOut)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrPipeOutNotAvailable
	}
	return &Client{
		// XXX: This looks backwards, but it's the same as r2pipe.
		writer: out,
		reader: in,
	}, nil
}
