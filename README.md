# Aporeto Internship Quiz Solution

Tests report for internship summer 2016 @ Aporeto Inc.

## Problem1 (Bash Shell Script sample)

- Detect if the certain file exists or not first, then create/overwrite file according to user's input.

## Problem2 (Python/Go sample)

- Because the length distribution of sentences is not clear (some or all of the sentences can be very long), fast hash digest function (md5) is used here for saving memory of hash list. But it might be overweight (mostly short sentence) or not enough (hash collision), so hash function is encapsulated as static method for further modification.
- Hash list is based on Set for its performance in Python.

## Problem3 (Go sample)

- Send HTTP(s) request to get file content, apply regexp to extract word and use hash list to count word.
- Concatenate output contents with `bytes.Buffer` for performance.
- Parallel Operation for each url is implemented in a Goroutine (named `worker`).
- Main function is blocked until all Goroutine emit finishSingal in a certain channel (called `finishChan`).