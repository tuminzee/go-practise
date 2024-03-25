# Select Channel Program

## Overview
This Go program demonstrates the use of channels and goroutines to concurrently send and receive data. It creates two channels (`c1` and `c2`) and uses a goroutine to send data to both channels. It then continuously listens for data on both channels using a select statement and prints the received data. If no data is available on either channel, it sleeps for 1 second.

## Code Explanation
- The `main` function initializes two channels `c1` and `c2`.
- A goroutine is started to send data to both channels and then close them.
- The main loop continuously listens for data on both channels using a select statement.
- If data is received on channel `c1`, it prints "data from channel 1" followed by the received value.
- If data is received on channel `c2`, it prints "data from channel 2" followed by the received value.
- If no data is available on either channel, it prints "No data pushed inside the channel going to sleep for 1 sec" and sleeps for 1 second.


## Continuous Execution
The program runs continuously due to its main loop with a select statement. It continuously checks for data on the channels and prints messages if no data is available, effectively creating a loop that runs indefinitely until manually terminated.


