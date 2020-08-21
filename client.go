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
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
)

// CommandHandler can send commands to radare and return the output.
type CommandHandler interface {
	Run(cmd string) ([]byte, error)
}

// New opens the given file with Radare2 and returns a Client handler
// to the instance.
func New(file string) (*Client, error) {
	if !filepath.IsAbs(file) {
		f, err := filepath.Abs(file)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve absolute path: %v", err)
		}
		file = f
	}
	r2cmd := exec.Command("r2", "-q0", file)
	stdin, err := r2cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to establish stdin pipe to radar2 process: %w", err)
	}
	stdout, err := r2cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to establish stdout pipe to radar2 process: %w", err)
	}
	if err := r2cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start radar2 process: %w", err)
	}

	// Let R2 startup and discard the inial outputs.
	if _, err := bufio.NewReader(stdout).ReadString('\x00'); err != nil {
		return nil, err
	}
	return &Client{
		writer: stdin,
		reader: stdout,
	}, nil
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
		if len(data) == 0 {
			return data, nil
		}
		return data[:len(data)-1], nil
	}
	if len(data) == 0 {
		return data, nil
	}
	return data[:len(data)-1], err
}
