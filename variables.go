//two ways to declare variables in go. a formal way and a shortcut.
a.//the formal way(var)
-we use this when you want to declare a variable but maybe wait to give it a value later or when you are declairing things outside of a fuction 

var age int = 25
var name string = "Gopher"

b.//the shortcut way(:=)
-most common way to declare variables inside a function.
-Go infers the type based on the value you give it.

score := 100 //Go knows this as an 'int'

message := "High score!" //Go knows this is a 'string'
