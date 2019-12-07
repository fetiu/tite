# Project DejaVa

```python
main()=
{
    print("Hello world!")
}
```

`DejaVa` is a computer language that aims for general usage.
It is independent form all kinds of paradigms.
It can be script, or compile. It can be objective, or functional.

What it defines for now is a few sets of syntaxes that can embrace most styles and patterns found in programming languages.

## basics
This project is based on one simple principle.

> Don't overcomplicate simple things.

Sometimes using a proper symbol is more intuitive than a new keyword.

```go
item1:=5
item2:int=5
```

This is how you declare a variable. 
I believe you could infer it based on experience.

## loop

What if we can generate a list just by putting `~` between two numbers?

```python
for i from 1 ~ 10 {
}
```

even characters?

```python
for a from 'a'~'z' {
}
```

loops made easy.


## types 
new types can be simply defined.

```python
var:[x:int,y:int]=(0,0)
```

The expression above is to declare a variable with two fields.
Actually, we don't even need to define before we use it.
We don't need any declarator like `struct`.

but we can explicitly give them an alias.

```go
type Pos 
[
    x: int
    y: int
]

var:Pos=(0,0)
```

Pretty much legible than before.
Note that types use brackets, not the braces.


## functions

functions can be defined with familiar characters.

```go
myfunc1()={}

myfunc2(msg:str)int=
{
    print("myfunction: %s", msg);
}
```

**Functions** itself are the data. Like strings surrounded by `""`
Instead of double quotes, functions are surrounded by `{}`
But the name of the function works just like a **variable** that never be changed.
That's why functions and variable expressions are slightly similar.

Then there should be some functions that can be changed.

```go
myinterface:(msg:str)int

myinterface = myfunc2
myinterface = myfunc1 //error
```

`myinterface` can conatin `myfunc2` since they have the same signiture.
But in case of `myfunc1` signitures are diffrent. There should be error.

As you can see, the signiture of a function works with just the same principle of variable and its type.

## string format

```python
print("hello world %s".fmt(username));

```
