package sql

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/Logiraptor/oak/flow/backends/test"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type DB struct {
	proc *exec.Cmd
	Conn *sql.DB
}

func NewPostgresDB() (DB, error) {

	tempDir, err := ioutil.TempDir(os.TempDir(), "test")
	if err != nil {
		return DB{}, err
	}

	startCmd := exec.Command("docker", "run",
		"-p", "127.0.0.1:9999:5432",
		"-e", "POSTGRES_PASSWORD=password",
		"-e", "PGDATA="+tempDir,
		"postgres")

	stderr, err := startCmd.StderrPipe()
	if err != nil {
		return DB{}, err
	}
	stdout, err := startCmd.StdoutPipe()
	if err != nil {
		return DB{}, err
	}
	go io.Copy(os.Stdout, stderr)
	go io.Copy(os.Stdout, stdout)

	err = startCmd.Start()
	if err != nil {
		return DB{}, err
	}

	// wait for container to come online
	db, err := sql.Open("postgres", "postgres://postgres:password@127.0.0.1:9999?sslmode=disable")
	if err != nil {
		return DB{}, err
	}

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		fmt.Println(err.Error())
		time.Sleep(time.Millisecond * 100)
	}

	return DB{
		proc: startCmd,
		Conn: db,
	}, nil
}

func (d *DB) Close() {
	d.Conn.Close()
	d.proc.Process.Signal(os.Interrupt)
	d.proc.Process.Wait()
}

func TestSQLStorage(t *testing.T) {
	db, err := NewPostgresDB()
	assert.NoError(t, err)
	defer db.Close()

	test.StorageContract(t, &SQLStorage{
		Conn: db.Conn,
	})
}
