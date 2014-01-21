// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"io"
	"os"
	"os/exec"
)

func SvgToPdf(svg string, pdfFile string) error {
	cmd := exec.Command("inkscape", "--without-gui", "--file=/dev/stdin", "--export-pdf="+pdfFile)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := stdin.Write([]byte(svg)); err != nil {
		return err
	}
	stdin.Close()
	io.Copy(os.Stdout, stdout)
	io.Copy(os.Stderr, stderr)
	return cmd.Wait()
}
