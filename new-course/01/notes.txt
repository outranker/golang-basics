8. variables & functions
- when you declare type int on a 64bit computer, then type int is automatically int64. it goes same with 32nit or 16bit 
computersdefault int depends on platform
- when defining multiple return types for a function, wrap them with brackets, inside brackets seperate them with 
commas -> func sum(a int) (int, string) {}
9. pointers
- case: imaginw we have a variable, say string, and we want to pass it to another function. we pass it using ampersand symbol
right before the variable like this: editString(&myString). and when we declare a function that receives a pointer we put
asterisk symbol to its type like this func editString(s * string) {...}
10. types and structs
- when there is a global variable in a module and some func takes arg named exactly like global variable
then whatever value passed is used instead of the global one. this is called varibale shadowing. seems natural to me
- struct type is kinda weird, not like that of JS. you declare struct in this way -> type Name struct { a string \n b string}
when initializing it like this -> user := Name{}. either pass values in object format with double colon and commas and stuff
or leave it empty. btw if struct fields are capitalized we can use it in other packages. when initializing a struct last 
field declaration should end with a comma. note: we are using varibale declaration and initialization. that's why we have
to open and close curly braces.
- when we are declaring a variable of type some struct we do it the JS way like: var myVar myStruct \n myVar.FirstName = "h" 
11. receivers: structs with functions
- when we declare receivers, we can pass type struct as pointer
12. other data structures: maps and slices
- 2 ways to create maps: first: var myMap map[string]string second: myOtherMap := make(map[string]string). second is better
- how does map syntax declaration work? first type in square brackets is what type its key should be and the other type
outside square brackets is what data type its value should be.
- position/order of map keys are randomized. cannot rely that their order that they were insert in sustain throughout its
lifetime
- creating slice is easy. pushing value is done via append(). 2 args are required. first one is slice itself and the second
one is new value to be pushed
- if we declare slice using var we can only (i think) add values using append (just like push) if we use declare&initialize
method we can use curly brackets appended to the type and type out values -> n := []int{1,2,3}
13. decision structures
- we don't need break for switch cases
14. loops and ranging over data
- as i know range returns index and value from arrays but turns out we can also use it with maps, strings, slices
15. interfaces
16. packages
17. channels
- channels are created using make built-in function
- imagine we want to create many go routines - functions that each happen on different thread than the current one
concurrently. but we cannot communicate with them because we don't have means to do that. channels come into play
in this situation. we pass the channel to the go routine and get the return value from that channel
18. reading and writing json
- typically when we call some api/external service, we receive data in structured format ie json and we cannot use it 
right away, because it's just a string. but here is catch we have to define a struct for that received data. and mathc
all keys with with keys of json data. then we can unmarshal that into our struct. to unmarshal json, we need a variable
to store the unmarshalled data so we declare it beforehand. then using enconding/json pkg, we use json.Unmarshal() func
it receives a byte as first argument and the variable to store it as second argument. btw, we give second value's 
address instead of plain itself.
- but how do we write json from struct. you guessed it right, we use json.Marshal() or we can use MarshalIndent
to pretty print it