package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "time"
    "sort"
)


type APIResponse struct {
	GroupNames string `json:"group_names"`
	Teams []string `json:"teams"`
}

type Result struct {
    URL       string
    GroupNames string
    Teams     []string
}

func consumer(results <-chan Result) {
    var allTeams []string

    // Collect & flatten
    for r := range results {
        allTeams = append(allTeams, r.Teams...)
    }

    // Sort descending (Z → A)
    sort.Slice(allTeams, func(i, j int) bool {
        return allTeams[i] > allTeams[j]
    })

    fmt.Println("")
    for _, team := range allTeams {
        fmt.Println(team)
    }
}


func producer(ctx context.Context, jobs chan<- string) {
	groups:= []string{"a", "b", "c", "d", "e", "f"}
    defer close(jobs)

	for _, group:= range groups{
		select {
        case <-ctx.Done():
            return
        case jobs <- fmt.Sprintf("http://localhost:5000/group%s", group):
        }

	}
}

func worker(
    ctx context.Context,
    wg *sync.WaitGroup,
    client *http.Client,
    jobs <-chan string,
    results chan<- Result,
) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            return

        case url, ok := <-jobs:
            if !ok {
                return
            }

            req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
            if err != nil {
                continue
            }

            resp, err := client.Do(req)
            if err != nil {
                continue
            }

            if resp.StatusCode != http.StatusOK {
                resp.Body.Close()
                continue
            }

            var apiResp APIResponse
            if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
                resp.Body.Close()
                continue
            }
            resp.Body.Close()

            // Transform API → internal type
            
            results <- Result{
                URL:       url,
                GroupNames: apiResp.GroupNames,
                Teams:     apiResp.Teams,
            }

            
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client := &http.Client{Timeout: 5 * time.Second}

    jobs := make(chan string)
    results := make(chan Result)

    var wg sync.WaitGroup

    // ---- start workers ----
    workerCount := 6
    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go worker(ctx, &wg, client, jobs, results)
    }

    // ---- start producer ----
    go producer(ctx, jobs)

    // ---- close results when workers finish ----
    go func() {
        wg.Wait()
        close(results)
    }()

// fmt.Println(results)
    // ---- start consumer (blocking) ----
    consumer(results)
}

