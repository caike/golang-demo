# ODevs Demo

This is a demo Go program showing some concurrency features.

## Concurrency On a Single CPU

To run this app using a single CPU:

`GOMAXPROCS=1 go run`

## Concurrency and Parallelism

To run this app using all the CPUs available on a machine:

`go run`


## Benchmarking

For those on Unix/Linux, the simplest way see the duration
of each run is to use the `time` tool. For example:

`GOMAXPROCS=1 time go run`

and

`time go run`
