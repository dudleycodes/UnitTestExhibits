// +build integration

package postgresql

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	// The docker used by these integration tests
	_dockerImage = "docker.io/library/postgres:13"

	_password = "some_password"
	_dbname   = "some_db"
)

var (
	// The System Under Test that will be used for all PostgreSQL integration tests
	_sut *Broker = nil
)

func TestMain(m *testing.M) {
	// Get a free TCP port from the operating system
	port := 0
	if addr, err := net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		if ln, err := net.ListenTCP("tcp", addr); err == nil {
			ln.Close()
			port = ln.Addr().(*net.TCPAddr).Port
		} else {
			fmt.Printf("Failed to connect to TCP on localhost: %s", err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Printf("Failed to resolve address for localhost: %s", err.Error())
		os.Exit(1)
	}

	closer, err := startPostgreSQL(port)
	if err != nil {
		fmt.Printf("Failed to create Docker container for %s: %s", _dockerImage, err.Error())
		os.Exit(1)
	}

	cfg := Config{
		Host:     "localhost",
		Port:     port,
		User:     "postgres",
		Password: _password,
		DBName:   _dbname,
	}

	_sut, err = New(cfg)
	if err != nil {
		fmt.Printf("Failed to create new PostgreSQL broker: %s", err)
		closer() // os.Exit() will bypass defer statements; closer() must be explicitly called
		os.Exit(1)
	}

	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)

		if _sut.Ping() == true {
			break
		}

		if i > 6 {
			fmt.Printf("PostgreSQL broker was unable to ping remote service in %d attempts", i)
			closer()
			os.Exit(1)
		}
	}

	exitCode := m.Run()

	closer()
	os.Exit(exitCode)
}

func TestRowCount(t *testing.T) {
	t.Parallel()

	// Create an empty table used only for this test
	tableName := fmt.Sprintf("rowcount_%d", time.Now().Unix())
	createTableQuery := fmt.Sprintf("CREATE TABLE %s (id SERIAL PRIMARY KEY, k VARCHAR (8))", tableName)

	if _sut.database == nil {
		t.Errorf("it was nil")
	}

	if _, err := _sut.database.Exec(createTableQuery); err != nil {
		t.Fatalf("Failed to create table %q for row count testing: %s", tableName, err)
	}

	t.Run("empty table should count 0 rows", func(t *testing.T) {
		count, err := _sut.RowCount(tableName)

		if err != nil {
			t.Errorf("Failed while executing RowCount(): %s", err.Error())
		}

		if count != 0 {
			t.Errorf("Counting the rows of an empty table should result in `0` not `%d`", count)
		}
	})

	t.Run("table with 6 rows should count 6 rows", func(t *testing.T) {
		for i := 0; i < 6; i++ {
			if _, err := _sut.database.Exec(fmt.Sprintf("INSERT INTO %s (k) VALUES ('%d')", tableName, i)); err != nil {
				t.Fatalf("Failed to insert row with k = `%d`", i)
			}
		}

		count, err := _sut.RowCount(tableName)

		if err != nil {
			t.Errorf("Failed while executing RowCount(): %s", err.Error())
		}

		if count != 6 {
			t.Errorf("Counting the rows of an empty table should result in `6` not `%d`", count)
		}
	})
}

// start up an instance of PostgreSQL in a Docker container; returns a function to stop and remove it.
func startPostgreSQL(port int) (closer func(), err error) {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("Failed to get docker client, %w", err)
	}

	reader, err := dockerClient.ImagePull(context.Background(), _dockerImage, types.ImagePullOptions{})
	if err != nil {
		return nil, fmt.Errorf("Failed to pull docker image, %w", err)
	}

	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return nil, fmt.Errorf("couldn't fetch docker image %q: %w", _dockerImage, err)
	}

	containerCfg := container.Config{
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", _password),
			fmt.Sprintf("POSTGRES_DB=%s", _dbname),
		},
		Image: _dockerImage,
	}

	hostCfg := container.HostConfig{
		AutoRemove: true,
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{HostPort: fmt.Sprintf("%d/tcp", port)},
			},
		},
	}

	containerName := fmt.Sprintf("test_postgresql_%d", time.Now().Unix())
	cont, err := dockerClient.ContainerCreate(context.Background(), &containerCfg, &hostCfg, nil, containerName)

	if err != nil {
		return nil, fmt.Errorf("Failed to create container, %w", err)
	}

	closeContainer := func() {
		if err := dockerClient.ContainerRemove(context.Background(), cont.ID, types.ContainerRemoveOptions{
			RemoveVolumes: true,
			RemoveLinks:   true,
			Force:         true,
		}); err != nil {
			fmt.Printf("failed to remove container: %s", err.Error())
		}
	}

	if err := dockerClient.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
		closeContainer()
		return nil, fmt.Errorf("Failed to start container, %w", err)
	}

	return func() {
		closeContainer()

		timeout := 10 * time.Second

		if err := dockerClient.ContainerStop(context.Background(), cont.ID, &timeout); err != nil {
			fmt.Printf("failed to stop container: %s", err.Error())
		}
	}, nil
}
