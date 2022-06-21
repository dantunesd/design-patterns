# Strategy

Sometimes you have different objects which contains different attributes but that behaves similarly, like Credit Card and Money, they have different attributes, but they serve as a payment method when you are shopping.

Implementing the strategy pattern, you standardize the behavior of these objects, so you are able to switch, add or remove these objects as you like, without impacting the other objects already created.

All you need to do is finding their behaviors similarities. Create an interface, so these objects implements it. Create a manager, which receives these objects instances and then, implement logic, which will be responsible to choose which one will be called.