package pipen

import (
	"io"
	"os/exec"
)

func StreamCommand(r io.Reader, program string, args ...string) (io.Reader, error) {
	progAbsPath, err := exec.LookPath(program)
	if err != nil {
		return nil, err
	}
	cmd := &exec.Cmd{
		Args: args,
		Path: progAbsPath,
		Dir:  ".",
	}

	closeOnErr := []io.Closer{}

	var latestErr error
	defer func() {
		if latestErr == nil {
			return
		}
		for _, c := range closeOnErr {
			_ = c.Close()
		}
	}()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		latestErr = err
		return nil, err
	}

	closeOnErr = append(closeOnErr, stdin)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		latestErr = err
		return nil, err
	}

	closeOnErr = append(closeOnErr, stdout)
	err = cmd.Start()
	if err != nil {
		latestErr = err
		stdin.Close()
	}

	pr, pw := io.Pipe()
	go func() {
		_, _ = io.Copy(stdin, r)
		_ = stdin.Close()
	}()

	go func(c *exec.Cmd) {
		_, _ = io.Copy(pw, stdout)
		_ = c.Wait()
		_ = pw.Close()
	}(cmd)

	return pr, nil
}
