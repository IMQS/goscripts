package goscripts

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// ExecFlags are command execution flags
type ExecFlags int

const (
	ExecFlagNone        ExecFlags = iota // If the command does not succeed (i.e. exit code non-zero), panic
	ExecFlagMustSucceed                  // If the command does not succeed (i.e. exit code non-zero), panic
	ExecFlagTrimSpace                    // Trim the spaces off the resulting stdout and stderr (for PExec)
)

// ExecCmd runs the command, and copies it's stdout and stderr
// to our stdout and stderr respectively.
// Returns the process exit code
func ExecCmd(cmd *exec.Cmd, flags ExecFlags) int {
	code, stdout, stderr := ExecCmdP(cmd, flags)
	os.Stdout.Write([]byte(stdout))
	os.Stderr.Write([]byte(stderr))
	return code
}

// Exec runs 'cmd' through ExecCmd(), and panics on failure
func Exec(cmd string, arg ...string) {
	ExecCmd(exec.Command(cmd, arg...), ExecFlagMustSucceed)
}

// ExecOrTrue is like Exec, but does not panic if the child process exits with a non-zero exit code.
// Returns exit code
func ExecNoPanic(cmd string, arg ...string) int {
	return ExecCmd(exec.Command(cmd, arg...), ExecFlagNone)
}

// ExecCmdP runs the command, and returns it's exit code, stdout, and stderr.
// If mustSucceed is true, and the command returns a non-zero exit
// code, then the function panics.
func ExecCmdP(cmd *exec.Cmd, flags ExecFlags) (exitCode int, stdout string, stderr string) {
	var stdoutB bytes.Buffer
	var stderrB bytes.Buffer
	cmd.Stdout = &stdoutB
	cmd.Stderr = &stderrB
	err := cmd.Run()
	if flags&ExecFlagMustSucceed != 0 && cmd.ProcessState.ExitCode() != 0 {
		Check(err)
	}
	exitCode = cmd.ProcessState.ExitCode()
	stdout = string(stdoutB.Bytes())
	stderr = string(stderrB.Bytes())
	if flags&ExecFlagTrimSpace != 0 {
		stdout = strings.TrimSpace(stdout)
		stderr = strings.TrimSpace(stderr)
	}
	return
}

// Exec runs 'cmd' through ExecCmdP(), and return stdout, stderr
func ExecP(cmd string, arg ...string) (stdout string, stderr string) {
	_, stdout, stderr = ExecCmdP(exec.Command(cmd, arg...), ExecFlagMustSucceed)
	return
}
