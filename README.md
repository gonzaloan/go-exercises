# Go Exercises and Examples

Just a repo with exercises and examples of advanced Golang Things :)


## Concurrency
 
### 1) Locks and Unlocks
Race Condition occurs when multiple threads try to access and modify the same data. If one thread tries to increase an integer and another thread tries to read it, this will cause a race condition. On the other hand, there won't be a race condition, if the variable is read-only.

You can check whether your code has potential race conditions using 
```go build --race class.go```

And one way of fixing this, is using Locks from the sync library. 



## Design Patterns


## Net

