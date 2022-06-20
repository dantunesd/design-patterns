# State

Sometimes we have to create objects which have a complex state flow to handle.

The most common way to handle it is creating a lot of "ifs" and "switch cases" in the same file/function. But going on this way you will probably end up creating a very hard and complex logic to understand.

In this example, a car has a complex state logic. If you accelerate it while off, nothing should happen. But if it's on, it should leave from where it stands.

Applying the State pattern, you'll split the complexity to handle states into many different classes. Each class handles only its own state. And when a state must be changed, the current state delegates to the next state class.

Another benefit of applying it is: creating new states is easy and has none to little impact on the other state classes already created.

