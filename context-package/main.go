package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getWindowsHostIP() string {
	out, err := exec.Command("bash", "-c", "ip route | grep default | awk '{print $3}'").Output()
	if err != nil {
		return "localhost"
	}
	return strings.TrimSpace(string(out))
}

type Result struct {
	workerID int
	value    string
	err      error
}

func worker(ctx context.Context, pool *pgxpool.Pool, id int, jobs <-chan string, results chan<- Result) {
	for query := range jobs {
		var output string
		err := pool.QueryRow(ctx, query).Scan(&output)
		if err != nil {
			results <- Result{workerID: id, err: err}
			continue
		}
		results <- Result{workerID: id, value: output}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	host := getWindowsHostIP()
	connStr := "postgres://postgres:postgres@" + host + ":5432/postgres?sslmode=disable"

	// pool instead of single conn — safe for concurrent use
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := []string{
		"SELECT version()",
		"SELECT current_database()",
		"SELECT current_user",
		"SELECT pg_postmaster_start_time()::text",
		"SELECT inet_server_addr()::text",
	}

	jobs := make(chan string, len(queries))
	results := make(chan Result, len(queries))

	// spin up 3 workers
	for i := range 3 {
		go worker(ctx, pool, i+1, jobs, results)
	}

	// send all jobs
	for _, q := range queries {
		jobs <- q
	}
	close(jobs) // no more jobs, workers will exit when done

	// collect all results
	for range queries {
		select {
		case res := <-results:
			if res.err != nil {
				fmt.Printf("worker %d error: %v\n", res.workerID, res.err)
			} else {
				fmt.Printf("worker %d → %s\n", res.workerID, res.value)
			}
		case <-ctx.Done():
			fmt.Println("timed out!")
			return
		}
	}
}
