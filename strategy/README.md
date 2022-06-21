# Strategy

Sometimes you have different objects that contains different attributes but behaves similarly, like Credit Card and Money, they have different attributes, but they serve as a payment method when you are shopping.

Implementing the strategy pattern, you standardize the behavior of these objects, so you are able to switch, add or remove these objects as you like, without impacting the other ones already created.

All you need to do is finding their behaviors similarities. Create an interface, so these objects can implement it. Create a manager, which receives these objects instances and then, implement your logic, which is responsible for choosing which one will be called.