# Chain of responsibility

Sometimes you want to apply a set of logics in sequence to do something. The most common and simple way to implement this is adding an "if" after another "if".

But when your application grows and you add more steps, those "if" after "if" gets more complex, hard to test, hard to remove or to add.

Implementing chain of responsibility, you create different handlers to do what you want separately of each other and then calls the next handler if necessary.

The benefitis of applying this pattern is that you isolate one handler from another, so each handler can contain a spefic logic. You can easily change their calling order in the chain. Also adding or removing handlers don't impact (shouldn't) the others. 
