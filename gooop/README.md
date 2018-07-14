# Object-oriented Programming with Go Course

This course presents a strategy to take advantage of go objects and teach how to design with the package scope in mind.

[pt-br] Esse curso apresenta uma estratégia para tirar vantagem dos objectos to go e ensina como planejar o software com o escopo do pacote em mente.

## Notes

### Object Oriented Programming History

* late 1950 (Conceptualization)
* 1967 (Simula67)
* 1970 (Smalltalk)

> Those who don't know history are doomed to repeat it
> Edmon Burke

### Challenges of OOP in Go

* Go is not a *class* and *object* language
* Go does not have *private* data
* Go does not support *inheritance*
* Go does not have *abstract* base types

### Go have

* Methods
* Package oriented design
* Type embedding
* Interfaces !

### Topics

* Encapsulation
* Message Passing
* Composition
* Polymorphism

### Encapsulation

Accessing a service on an object *without knowing* how that service is implemented.

Delivery to the object the control what the consumer can interact with.

**Package Oriented Design** - treat packages as the lowest organizational unit.

### Interface

A data structure is not accessible outside the package. To make it available define an interface to expose methods to make data available.

### Message Passing

Sending a message to an object but letting that object determine what to do with it.

The more Sender is ignorant about Receiver services, the more decoupled the system is.

**Can work by Interface or by Channel**

**Interface** - Use the service directly by a service interface. You can change the interface underlying.

**Channel** - Pass a channel to object constructor and then use the channel. Inside the constructor, there is the logic applied to the channel.

### Composition

Behavior reuse strategy where a type contains objects that have desired functionality. The type delegates call to those objects to use their behaviors.

* Inheritance not supported
* Composition (Embedding) - compose a type with other types

 :warning: To resolve name conflicts between two items that compound an object, use a "local function" on the object.

### Polymorphism

The ability to transparently substitute a family of types that implement a common set of behaviors.

 :memo: Declare the interface in the Consumer scope, it allows to use the interface across packages. 

## Course details

| | |
| - | - |
| Source | [Pluralsight](pluralsight.com) |
| Title | [Object-oriented Programming with Go](https://app.pluralsight.com/library/courses/go-object-oriented-programming/description) |
| Author | Mike Van Sickle |
| Level | Intermediate |