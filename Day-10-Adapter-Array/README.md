# Day-10-Adapter-Array: 

## Part 1
Simple searching for first available element in array

## Part 2
Looking for permutation of possible ways of plug connection. Create acyclic graph and count paths. It should be
efficient.

### First try
recursion with returning 1 when leaf found (max voltage value), 0 otherwise. In the parent node iteate throgh edges 
and sum child results. This approach is not efficient. Result found after min.

### Second try

## Run

From this path (`./advent-code-2020/Day-10-Adapter-Array`) just:

`make run`

## Run test

From this path (`./advent-code-2020/Day-10-Adapter-Array`) just:

`make test`
