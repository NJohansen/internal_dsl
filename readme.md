# Prototype of a HUE system - internal DSL

## Prototype 1 - Golang
The finally prototype i ended up with was an internal DSL written in Golang. 
The internal DSL is for a statemachine that should manage a Phillips Hue System with their light bulbs.

I've started by writing the DSL in Java but then converted it to the host language Golang.
It was done with inspiration from the Java version [Ulrik Pagh Schultz github repo]([https://link](https://github.com/ulrikpaghschultz/MDSD/tree/master/src/examples/mini_state_machine/)).

### How to run?
1. You will need to have a Go environment on your machine.
2. Then you can just write `go run main.go`. Be aware you have to locate the `main.go` file.
3. You can try changing the processing of event by changing one `When("some event")` and then write a new `hi.ProcessEvent("the event from the When method to process")`

### Considerations
There were some considerations when i wrote the DSL in Golang. 

* I was excited about how much code i needed in Go to get the same HUELight statemachine example.
* Golang does not provide the concept classes as in Java, i had to figure out how to write OOP in Golang for my internal DSL.
* A design thought i had was how did i create the statemachine generic but still with the thought of creating a DSL for a specific problem domain. An example of this is my IsLessThan(), IsMoreThan() or Equals() because what they should do is some integer comparison for brightness. The methods a the exact same, but i could not find a generic method name for one solution so i ended up with the three methods.
* I added lambda functions for the different To() methods in my method chaining, to show that this would be how i would execute code.

### Pros and cons

The pros of using Golang
* As i mentioned earlier i did not need the same amount of code to write the DSL.
* I don't have the same amount of types like ArrayList, HashMap, List and such with all their underlaying methods. I got map and slice and therefrom i had to implement the methods myself. Don't know if this is actually a pro or a con. I kind of liked it.
* Packages instead of classes.

The cons
* Method chaining and the builder pattern is not created for Golang since we in Golang can recieve an error in our method chain which would stop executing the chain and throw an error. Like in Java where we can throw exceptions anywhere and still return the builder class. This is not possible in Go as the only way is to return an error. What i needed was also to return my builder instance after every method, so i kept having it in scope. But when i first solved how i should point to the right builder instance it was fine.
* In Golang i had to keep referencing with the help of pointers as Golang per defeault does not know if it should initialize a new object or i am referencing to an existing Builder for example. This was kind of tough to find out, and i guess that is the complications that you can face by transferring an internal DSL to another host language. 

## Prototype 2 - Java without Variable class in Metamodel
Before i created my internal DSL in Golang i had to create the DSL in Java so i knew my statemachine example was working. 
It was done with inspiration from [Ulrik Pagh Schultz github repo]([https://link](https://github.com/ulrikpaghschultz/MDSD/tree/master/src/examples/mini_state_machine/)).

### How to run the example?
To run the example you just have to run the `HUELightBulbs.java` from the package `example_state_machine_without_variable`.
Where you then can see how the statemachine runs.

## Prototype 3 - Java with Variable class in Metamodel
I thought i had a good idea of creating an extended metamodel where I instead handled Variables from the HUE system in a class.
That way i had a generic way of creating variables like "Brightness" or "Color" variables. I did found out that this was kinda of the same as a state and it was easier to handle in state therefore i wrote **Prototype 2**
This was also done with inspiration from [Ulrik Pagh Schultz github repo]([https://link](https://github.com/ulrikpaghschultz/MDSD/tree/master/src/examples/mini_state_machine/)).

### How to run the example?
To run the example you just have to run the `HUELightBulbs.java` from the package `example_state_machine_with_variable`.
Where you then can see how the statemachine runs.