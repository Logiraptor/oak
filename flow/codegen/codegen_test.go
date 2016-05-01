package codegen

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/Logiraptor/oak/flow/loader"
	"github.com/Logiraptor/oak/flow/parser"
	"github.com/stretchr/testify/assert"

	"testing"
)

func execCommand(sh []string, stdin string) (stdout string, stderr string) {
	cmd := exec.Command(sh[0], sh[1:]...)
	var (
		stdoutBuf = new(bytes.Buffer)
		stderrBuf = new(bytes.Buffer)
	)
	inPipe, _ := cmd.StdinPipe()
	outPipe, _ := cmd.StdoutPipe()
	errPipe, _ := cmd.StderrPipe()
	go func() {
		io.Copy(inPipe, strings.NewReader(stdin))
		inPipe.Close()
	}()
	go io.Copy(stdoutBuf, outPipe)
	go io.Copy(stderrBuf, errPipe)
	cmd.Run()
	return stdoutBuf.String(), stderrBuf.String()
}

func TestCodegen(t *testing.T) {
	paths := []string{}
	filepath.Walk("../../examples", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".yml" {
			paths = append(paths, path)
		}
		return nil
	})
	t.Logf("Found %d examples", len(paths))

	for _, path := range paths {
		t.Logf("Generating: %s", path)
		prog, err := parser.ParseFile(path)
		assert.Nil(t, err)
		app, errs := loader.Load(prog)
		assert.Len(t, errs, 0)

		tmp, err := ioutil.TempDir("", "flow-test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmp)

		WriteFlowApp(app, tmp)

		type TestConfig struct {
			Args   []string
			Stdin  string
			Stdout string
			Stderr string
		}

		var tc TestConfig

		buf, err := ioutil.ReadFile(filepath.Join("testdata", filepath.Base(path)+".out"))
		assert.Nil(t, err)

		err = yaml.Unmarshal(buf, &tc)
		assert.Nil(t, err)

		goFiles, _ := filepath.Glob(filepath.Join(tmp, "*.go"))

		cmd := []string{"go", "run"}
		cmd = append(cmd, goFiles...)
		cmd = append(cmd, tc.Args...)

		t.Logf("Executing %s", strings.Join(cmd, " "))

		stdout, stderr := execCommand(cmd, tc.Stdin)
		assert.Equal(t, tc.Stdout, stdout)
		assert.Equal(t, tc.Stderr, stderr)

	}
}
