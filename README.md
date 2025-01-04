# Min Heap to find top log its 

## Running the project

This is a Golang project. The main file is `main.go` in the root directory.

Example:

```
go run main.go test_files/example_input_data_1.data 5
```

In order to make good use of my resources, I solve this problem using a min-heap of the highest values. I read each row of the file without parsing the JSON, if the score is greater than the smallest value in my min-heap, it replaces that smallest value and the heap is rebalanced. 

