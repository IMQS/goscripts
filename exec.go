package goscripts

import (
	"os"
	"os/exec"
	"bytes"
)

// ExecCmd runs the command, and copies it's stdout and stderr
// to our stdout and stderr respectively.
// If mustSucceed is true, and the command returns a non-zero exit
// code, then the function panics.
// Returns the process exit code.
func ExecCmd(cmd *exec.Cmd, mustSucceed bool) int {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	os.Stdout.Write(stdout.Bytes())
	os.Stderr.Write(stderr.Bytes())
	if mustSucceed && cmd.ProcessState.ExitCode() != 0 {
		Check(err)
	}
	return cmd.ProcessState.ExitCode()
}

// Exec runs 'cmd' through ExecCmd(), and panics on failure
func Exec(cmd string, arg ...string) {
	ExecCmd(exec.Command(cmd, arg...), true)
}

// ExecOrTrue is like Exec, but does not panic if the child process exits with a non-zero exit code.
func ExecNoPanic(cmd string, arg ...string) int {
	return ExecCmd(exec.Command(cmd, arg...), false)
}
