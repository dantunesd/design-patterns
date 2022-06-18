# Facade

Sometimes we want to execute a "simple" process, but them we see that this little process requires to instantiate many classes and call lots of methods in order.

We still have to do this "hard job", but creating a facade we hide all the creation and execution complexity behind a simplified interface.

So everytime you need to execute this process again, we simply call method that will do that hard work.