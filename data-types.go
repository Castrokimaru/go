//commom data types.
-quick cheat sheet for the types applicable 90% of the time.

type   Description  Example
int    Whole numbers  42,-7
float64  Decimal numbers  3.14,99.9
string  Text  "Hello"
bool  True/False  true


the "zero value concept"-in many languages, if you dont give a variable a value, it becomes null or undefined,which causes crashes.In Go, every varibale has a Zero value by default.
int defaults to 0
string defaults to ""  (empty string)
bool defaults to false.

//once a variable is an int, it stays an int forever.this is why Go is fast and safe.

