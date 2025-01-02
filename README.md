# Go Channel Hang Issue

This repository demonstrates a subtle issue involving channels in Go.  The program appears to correctly close the channel, yet it hangs indefinitely.

## Problem Description

A goroutine sends 100 boolean values to a channel. After sending all values, the channel is closed. The main goroutine reads from the channel until it's closed. However, the main goroutine hangs after reading all the values, seemingly not recognizing the channel's closed state.

## Solution

The solution addresses the unexpected behavior by adding a check to the loop's condition to ensure the channel's closed state is properly handled. This prevents the loop from continuing infinitely.

## How to reproduce

1. Clone the repository
2. Run `go run bug.go`
3. Observe that the program hangs indefinitely.
4. Run `go run bugSolution.go`
5. Observe that the program exits cleanly after processing all data. 

This demonstrates how to prevent a subtle goroutine related hang in Go.