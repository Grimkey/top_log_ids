# Emerald Cloud Lab homework

## Running the project

This is a Golang project. The main file is `main.go` in the root directory.

Example:

```
go run main.go test_files/example_input_data_1.data 5
```

## CPU Resource Efficiency

In order to make good use of my resources, I solve this problem using a min-heap of the highest values. I read each row of the file without parsing the JSON, if the score is greater than the smallest value in my min-heap, it replaces that smallest value and the heap is rebalanced. This should give me O(n) efficiency on datasets with small values for N. As N, gets higher, the operational complexity becomes O(n log n) because the heap's complexity has to be included.

I considered writing a custom parser and using Regex for finding the "id" to return. I rejected both ideas because they are fragile and I didn't know if there would be any variation in real logs. As part of collecting the information to return, I parse record string as JSON and pull out the "id".

## Memory Resource Efficiency

The memory efficiency primarily comes from not holding onto memory for each line in the file. Only the top N lines are kept, and so I rely upon Go's garbage collector to cleanup memory. At no time, should we need more than the top N plus one lines in memory. For small Ns, this should be very memory efficient, as N grows, it requires keeping an increasingly number of lines in memory.

One optimization that we could make for large values of N is to parse the record JSON earlier to keep a small amount of memory. But given the test file sizes, I did not believe it would be more useful than the CPU gain for parsing only when we were done.

## Assumption for unstated requirements

The error condition section did not specify how to handle when a parameter to the function was missing. I chose to return a helpful note incidating what the paramters were if one or both were missing.
