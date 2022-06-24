# Observer

Sometimes you want to do some things after something else has happened. For example, logging an object state change, publish an event to a message broker, notify another system about what ocurred, and so on.

One of the common ways to do that is calling all these things you need after your desired operation ocurrency. However, with your application growth, such class or function probably will end up with too many dependencies, doing a lot of things, with many ifs spread out, etc.

Implementing observer pattern, you can register as many listeners as you like (classes/structs) for certain events(and you define it), and after the event dispatch, these listeners will be notified about the occurred and then they will do what they were supposed to do.

Some of the benefits of this pattern is that you can add or remove as many listeners related to an event as you want, without having to change the existing ones. You don't have to change the class that dispaches these events.
