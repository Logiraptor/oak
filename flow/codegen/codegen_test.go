package codegen

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Logiraptor/oak/flow/loader"
	"github.com/Logiraptor/oak/flow/parser"
	"github.com/stretchr/testify/assert"

	"testing"
)

func createApp(t *testing.T, in string) loader.App {
	conf, err := parser.ParseReader(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}
	app, errs := loader.Load(conf)
	if errs != nil {
		t.Fatal(errs)
	}
	return app
}

func execCommand(sh []string, stdin string) (stdout string, stderr string) {
	cmd := exec.Command(sh[0], sh[1:]...)
	var (
		stdoutBuf = new(bytes.Buffer)
		stderrBuf = new(bytes.Buffer)
	)
	inPipe, _ := cmd.StdinPipe()
	outPipe, _ := cmd.StdoutPipe()
	errPipe, _ := cmd.StderrPipe()
	go io.Copy(inPipe, strings.NewReader(stdin))
	go io.Copy(stdoutBuf, outPipe)
	go io.Copy(stderrBuf, errPipe)
	cmd.Run()
	return stdoutBuf.String(), stderrBuf.String()
}

func TestCodegen(t *testing.T) {
	app := createApp(t, `
entry: start
imports:
  fmt: "fmt"
  base: "github.com/Logiraptor/oak/core/base"
  strings: "strings"
components:
  start: base.StringCLI
  caps: strings.ToUpper
  end: fmt.Println
flow:
  start: caps
  caps: end
`)
	tmp, err := ioutil.TempDir("", "flow-test")
	if err != nil {
		t.Fatal(err)
	}
	err = WriteFlowApp(app, tmp)
	assert.Nil(t, err)

	goFiles, _ := filepath.Glob(filepath.Join(tmp, "*.go"))

	cmd := []string{"go", "run"}
	cmd = append(cmd, goFiles...)
	cmd = append(cmd, "hello")

	stdout, stderr := execCommand(cmd, "")
	assert.Equal(t, "HELLO\n", stdout)
	assert.Equal(t, "", stderr)
}
