package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Grimkey/ecl_hq/ecl_heap"
)

// ParseLine assumes a valid line is `score: record` and will attempt to convert it into a LogElement.
func ParseLine(line string) (*ecl_heap.LogElement, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return nil, errors.New("invalid line format. No ':' symbol found")
	}

	score, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, errors.New("score is not a number")
	}

	record := strings.TrimSpace(parts[1])
	if len(record) == 0 {
		return nil, errors.New("no record found")
	}

	return &ecl_heap.LogElement{
		Score:  score,
		Record: record,
	}, nil
}

// BuildHeap traverses the file and creates a top n heap.
func BuildHeap(file *os.File, heap *ecl_heap.TopLogHeap) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignore empty lines
		if line == "" {
			continue
		}

		element, err := ParseLine(line)
		if err != nil {
			return err
		}
		heap.TryAdd(element)
	}
	return nil

}

// isFromJSON is a helper to pull the "id" field from the JSON.
func idFromJSON(jsonStr string) (string, error) {
	var data map[string]any

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	if id, ok := data["id"].(string); ok {
		return id, nil
	}

	return "", errors.New("'id' not found in string")
}

type output struct {
	Score int    `json:"score"`
	ID    string `json:"id"`
}

func FormatOutput(logHeap *ecl_heap.TopLogHeap) ([]output, error) {
	var result []output

	logs := logHeap.Write()
	for _, log := range logs {
		id, err := idFromJSON(log.Record)
		if err != nil {
			return nil, err
		}
		result = append(result, output{
			Score: log.Score,
			ID:    id,
		})
	}

	return result, nil
}

func main() {
	expected_parameters := 3
	if len(os.Args) != expected_parameters {
		fmt.Println("Usage: go run main.go <filename> <number>")
		return
	}

	filename := os.Args[1]
	top_n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}

	logHeap := ecl_heap.NewLogHeap(top_n)
	err = BuildHeap(file, &logHeap)
	file.Close() // Avoid defer because os.Exit does not invoke it.

	if err != nil {
		os.Exit(2)
	}

	output_list, err := FormatOutput(&logHeap)
	if err != nil {
		os.Exit(2)
	}

	jsonResult, err := json.Marshal(output_list)
	if err != nil {
		os.Exit(2)
	}

	fmt.Println(string(jsonResult))
}
