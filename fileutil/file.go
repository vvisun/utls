package fileutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/vvisun/utls/errutil"
)

func CopyFile(dst, src string) int32 {
	if dst == "" || src == "" {
		return errutil.Sys_InvalidArg
	}

	srcDir, _ := filepath.Split(src)
	// get properties of source dir
	srcDirInfo, err := os.Stat(srcDir)
	if err != nil {
		return errutil.Sys_UnKnown
	}

	dstDir, _ := filepath.Split(dst)

	MakeDirIfNeed(dstDir, srcDirInfo.Mode())

	sf, err := os.Open(src)
	if err != nil {
		return errutil.Sys_UnKnown
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return errutil.Sys_UnKnown
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		return errutil.Sys_UnKnown
	}

	return errutil.Succ
}

func CopyDir(dst string, src string) int32 {
	// get properties of source dir
	srcDirInfo, err := os.Stat(src)
	if err != nil {
		return errutil.Sys_UnKnown
	}

	// create dest dir
	err = MakeDirIfNeed(dst, srcDirInfo.Mode())
	if err != nil {
		return errutil.Sys_UnKnown
	}

	srcDir, _ := os.Open(src)
	objs, err := srcDir.Readdir(-1)
	if err != nil {
		return errutil.Sys_UnKnown
	}

	const sep = string(filepath.Separator)
	for _, obj := range objs {
		srcFile := src + sep + obj.Name()
		dstFile := dst + sep + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			if e := CopyDir(dstFile, srcFile); e != errutil.Succ {
				return e
			}
			continue
		}

		e := CopyFile(dstFile, srcFile)
		if e != errutil.Succ {
			return e
		}
	}

	return errutil.Succ
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func MakeDirIfNeed(dir string, mode os.FileMode) error {
	dir = strings.TrimRight(dir, "/")

	if FileExists(dir) {
		return nil
	}

	err := os.MkdirAll(dir, mode)
	return err
}

func Unused(args ...interface{}) {}

func RunCmd(cmdName string, workingDir string, args ...string) (string, error) {
	const duration = time.Second * 7200

	cmd := exec.Command(cmdName, args...)

	if workingDir != "" {
		cmd.Dir = workingDir
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}

	chanErr := make(chan error)
	go func() {
		multiReader := io.MultiReader(stdout, stderr)
		in := bufio.NewScanner(multiReader)
		for in.Scan() {
			buf.Write(in.Bytes())
			buf.WriteString("\n")
		}

		if err := in.Err(); err != nil {
			chanErr <- err
			return
		}

		close(chanErr)

	}()

	// wait or timeout
	chanDone := make(chan error)

	go func() {
		chanDone <- cmd.Wait()
	}()
	select {
	case <-time.After(duration):
		cmd.Process.Kill()
		return "", fmt.Errorf("run command: %s failed with timeout", cmdName)

	case err, ok := <-chanErr:
		if ok {
			return "", err
		}

	case e := <-chanDone:
		fmt.Printf("error %+v\n", e)
	}

	return buf.String(), nil
}
