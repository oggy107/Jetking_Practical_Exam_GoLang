Q1. Write a Golang program for implementing the concurrency with Goroutines.

TLDR: To demonstrate goroutines we are using goroutines to generate the sum of elements in array concurrently by spliting array in 2 halves

steps:
1. create a function named sum and pass in array and channel
2. calculate the sum of elements in array by iterating over it and pass the sum to the channel
3. im main funciton create an array and channel and split the array in two halves
4. pass halves into two sum functions along with channel and mark them as goroutines using go keyword so that they execute concurrently
5. listen for values returned by channel and store them in x and y as we have two goroutines
6. display the results. The sum of total elements in array would be addition of x and y
7. screenshot of output is in ss/q1.png