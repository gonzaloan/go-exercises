# Go Exercises and Examples

Just a repo with exercises and examples of advanced Golang Things :)


## Concurrency
 
### 1) Locks and Unlocks
Race Condition occurs when multiple threads try to access and modify the same data. If one thread tries to increase an integer and another thread tries to read it, this will cause a race condition. On the other hand, there won't be a race condition, if the variable is read-only.

You can check whether your code has potential race conditions using 
```go build --race class.go```

And one way of fixing this, is using Locks from the sync library. 


## Design Patterns

### Creational
Establish mechanisms aiming that the creation of objects can be reusable and flexible. Examples: Factory / Singleton

#### Factory
Allow us to create a "Factory" of objects from a base class and implement polymorphic behaviors to modify the inherited classes.
#### Singleton
Allow us to handle and restrict only one instance of a class. An example of its use, is to create database connections, and avoid creating multiple connections.
### 2) Estructural
Establish mechanisms of how create objects in larger structures without losing flexibility and reusability. Example: Adapter
### 3) Behavioral
Establish mechanisms of effective communication between objects and the assignation of responsibilities. Examples: Observer and Strategy


## Net

