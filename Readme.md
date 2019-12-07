# Project DejaVa

```python
main()=
{
    print("Hello world!")
}
```

`DejaVa` is a computer language that aims for general usage.

It is independent from all kinds of paradigms.
It can be script, or compile. It can be objective, or functional.

What it defines for now is a few sets of syntaxes that can embrace most styles and patterns found in programming languages.

## basics

This project is based on one principle.

> *Don't overcomplicate simple things.*

Sometimes using a proper symbol is more intuitive than a new keyword.

```rust
a:=0
b:int=5
```

This is how you declare a variable. 
I believe you could infer it based on experience.

## loop

What if we can generate a list just by putting `~` between two numbers?

```rust
for i:int 1 ~ 10 {
}
```

even characters?

```rust
for a:char 'a'~'z' {
}
```

loops made easy.


## types 

New types can be simply defined.

```python
point:[x:int,y:int]=(0,0)
```

The expression above is to declare a variable with two fields.
Actually, we don't even need to define before we use it.
We don't need any declarators like `struct`.

But we can also explicitly give them an alias.

```go
type Pos
[
    x: int
    y: int
]


point:Pos=(0,0)
```

Pretty much legible than before.
Note that types use brackets, not the braces.


## functions

Functions can be defined with familiar characters.

```go
myfunc1()={}

myfunc2(msg:str)int=
{
    print("myfunction: %s", msg);
}
```

**Functions** itself are the data. Like strings surrounded by `""`.
Instead of double quotes, functions are surrounded by `{}`.

But the name of the function works just like a **variable** that never be changed.
That's why functions and variable expressions are slightly similar.

Then there should be some functions that can be changed.

```go
myinterface:(msg:str)int

myinterface = myfunc2
myinterface = myfunc1 //error
```

`myinterface` can conatain `myfunc2` since they have the same signiture.
But in case of `myfunc1` signitures are different. There should be error.

As you can see, the signiture of a function works with just the same rule of variable and its type.

## const

To make constants, replace `:` in variable with `()`

```go
zero()int=0
five()int=5
```

No one can ever change `zero` and `five`.

This expression is strong when you trying to make a getter

```rust
getString()str="hello world"
```

You can simply implement string getter without `return`.
