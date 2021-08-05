FOR loop
3 components separated by ;
Init statement: executed before the first iteration
Condition expression: evaluated before every iteration
Post statement: executed at the end of every iteration
Gotcha: init and post statements are optional: for ; sum < 100; {...}
Gotcha 2: for loop is while in Go if we omit the ;  for sum < 100 {...}

IF statement
Looks like for loop. Expression comes naked
We can declare a variable in the if statement. ; is used for separation
if a := 3; a < myVar {...}





GORSE:
Packages:
Cobra - used for building cli for the project
Zap - used for logging (like pino or winston)

GO COURSE UDEMY
To add values to a slice just use curly braces
mySLice := [ ]string{“one”, “two”, “three”}
append is used for adding elements to an array
append doesn’t modify the array but returns a new array with the new values

for i, card := range cards {...}
2 things to note here. 1. i is not initiated the standard way, we are just providing a variable and for loop is just giving it an index position. 
2. What range does is take the index and make a copy of value and assign it to the card. I just noticed that we are declaring i and card variables together. 
Range obviously returns index and card. This is what I found from golang docs:`A "for" statement with a "range" clause iterates through all entries of an array, s
lice, string or map, or values received on a channel. For each entry it assigns iteration values to corresponding iteration variables if present and then executes the block.`

20. What comes right before the function name is called a receiver aka a receiver on a function.
49. Declaring maps is map[keyType]valueType - colors := map[string]string{“red”:”#ff0000”}. So as you can see there is a small gotcha here. With map first type 
is for key and the second one is for values
50. To access map value by its key we always use square braces syntax. Like mayMap[3] = 33 or fmt.Println(myMap[“keyName”])
There is a built-in function called delete which we can use to delete key-value fields from a map. Like this: delete(myMapName, keyName)

54. When we are writing struct method func we don't have to declare a value that points to struct that we can use in the func. We can just write only the struct 
type like this: func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there!"
}
Above, usually we write a value and type inside a receiver statement like (sb englishBot) but if we don't use sb we can just omit it.

55. So I understood what the honorary member idea is all about. Now I am going to make my own explanation for this interface thing.
Case: we have two structs and both of them have two similar functions and they basically do the same thing. The structs also are very identical. But instead of 
making one function a method of theirs we make it a regular function. But we have a problem here. This function accepts a struct type variable but since we have 
two structs we can only choose one of them in our function statement. So to solve this we make an interface which has this new function name which accepts struct 
type as argument. We were initially torn between the choice of choosing only one struct as argument type where instead we wanted to make this function accept 
either struct types. So this new interface has a method which the two structs also have. Now if you can see we are kinda making those structs relative using an 
interface. How so? This interface has a method, two structs also have a method and the method name is the same! Get it? Now the new function accepts interface 
type as argument and now we can pass either struct as argument to this new function. And this variable passed has the method that both structs have so we can access it.
56. Methods of interfaces don’t need to declare argument variable name but just argument type:
type bot interface { getGreeting(string, int) (string, error)

58. Struct values can be interfaces that means this value can belong to this struct as long as its value satisfies this interface.
Interfaces can be a value of interfaces we just need their name as value not type or return type needed

Golang gotchas:
* all the functions and files related to one package must reside in one directory
* package functions starting with lowercase letter is a private function and can only be used inside the owner package, if otherwise can be used in other packages as well


74. go command creates routines but the function that spawns the routines may exist before they are completed. That’s why we use channels. Channels are the only way 
the main routine speaks to all child routines.
Golang cool-tip:
* make returns a variable of the given type

76. Sending and receiving data from and to the channel. 
Receiving data from a channel: getFromChan := <- channelName
Sending data from channel: channelName <- data

getFromChan := <- channelName ::: channelName <- data

If you want to get something from channel then that channel name is on the left
If you want to send something through the channel then channel name is on the right
