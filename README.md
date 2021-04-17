Communicating Services
======================
golang-soa

# Basic Idea
Golang is by design not an object-oriented programming language. It explicitly tries to avoid the pitfalls of OOP languages
like Java that much too often end in inheritance cascades. While providing something quite similar to inheritance
(anonymous fields in structs), it makes composition the default case, not inheritance. That's definitely a good point
because it reflects the state of the art in modern software development and avoids the often seen horror of long inheritance chains.

Smalltalk is sometimes considered one of the first OOP languages, because its central concept is that of an object
I consider two things to be central to the idea of object orientation:
* encapsulation,
* coordination via communication, not via sharing.

Now, Golang makes it quite difficult to encapsulate data because it supports encapsulation only at the package level.

The other idea is maybe more important to Golang. It is especially crucial to Go because it makes the development of concurrent
workflows much easier.

*The main idea is that of small asynchronous services communicating by messages sent through channels.*

# Challenges in Asynchronous Programming
* Avoid *data races* and *deadlocks*.
* Unit testing
* Ordered shutdown.

# Architectural patterns

## Pipelining services
Managing asynchronous tasks and reasoning about the workflow should be easy and natural.

### every service consists of a struct and a (possibly) endless for-loop.

    type MyService struct {}
    func NewMyStruct() *MyService {
        go func() {
            for true {
                // do stuff here
            }
        }()
        return &MyService{}
    }

### every channel has an exclusive owner with exclusive to send messages
To make the messages available to receivers, the channel owner does not offer a traditional
getter function but a getter for the channel as receiver:


    type Node interface {
        Messages() <-chan Message
    }

Important: Communicate via channels, not via getters. A channel can transmit more information that a getter, especially
the information whether there are more messages to be sent. If there are no more messages, the sender can close the channel
thereby telling all the receivers that the last messages have been sent.

### use closing of channels for ending the process
Sometimes people use the concept of a `Context` to coordinate the application shutdown. Every Goroutine (or every struct 
used in such a goroutine) has access to some main context. To close the program the context is cancelled, thereby giving
all the goroutines the signal to end their work.
Unfortunately, that doesn't lead to a well-ordered shutdown. Sometime goroutines will end their job before all the tasks
that are already queued have been sent to them. The effect is some unordered loosing of tasks.

A better way is to use the means offered by channels: Every service closes its outgoing channels as soon as all the work
is done. The services relying on them will recognize that no more work is coming when they see the incoming channels getting
closed. They can then finish their job and close their outgoing channels before returning from their main-loop.
