# `tejava`

```go
io:stdio

main()=
{
    io.print("Hello world!")
}
```

`tejava` is a simple programming language.

## basics


Sometimes using a proper symbol is more intuitive than a new keyword.

```go
a:=0
b:int=5
c:int[5]=(1,2,3,4,5)
```

This is how you declare a variable. I believe you can figure it empirically.

## loop

What if we can generate a list just by putting `~` between two numbers?

```go
for i (1~5) {
    print(i)
}
```

even characters?

```go
for a ('a'~'z') {
    print(a)
}
```

loops made easy.

## types 

To define a new type, just create a file.

> point.tj

```go
x:int
y:int
```

`point.tj` defines a type with two fields.

Now, type `point` is available in other modules.

```go
p:point=(0,0)
```

As you can see, file name is regarded as a type name.

The act of creating a file itself is making a module and a type at the same time.

We don't need any declarators like `struct`, `class`.


## functions

Functions can be defined with familiar characters.

```go
myfunc1()={}

myfunc2(msg:str)int=
{
    print("myfunc2: %s", msg);
}
```

**Function** itself is a data. Like strings surrounded by `""`.
Instead of double quotes, function uses `{}`.

The name of the function works just like a **variable** that never be changed.
That's why function and variable expression both use `=` to assign value.

Then there should be some functions that **can** be changed.

```go
myinterface:(msg:str)int

myinterface = myfunc2
myinterface = myfunc1 //error
```

`myinterface` can conatain `myfunc2` since they have the same signiture.
But in case of `myfunc1` signitures are different. There should be error.

As you can see, the signiture of a function works with just the same rule of variable and its type.

## const

Constants are read-only data that ​​that can't be changed at runtime.

Constants can be only existed in global scope.

```go
zero=0
five=5
```

Assigning value into `zero` and `five`is prohibited henceforth.

Note that defining a constant inside a function is not allowed.

```go
func()=
{
    four=4 //error: expected ':' before assignemnt.
}
```

Every part inside braces `{}` should be have "local" scope. 

It does not make sense that a constant has a lifetime, hence it is deprecated.

## objects

OOP is possible in tejava.

Every tejava file is a module, type, class. In order to use components of a module, we should make its instance/object.

> point.tj

```go
x:int=0
y:int=0

move(dx:int,dy:int)=
{
    x += dx
    y += dy
}
```

> main.tj

```go
io:stdio

main()=
{
    p1:point
    p1.move(2,3)

    p2:point
    p2.move(5,5)

    io.print(p1)
    io.print(p2)
}
```

Example above shows how to define a class `point` and to use it.

Defining a variable with `point` type, does both inclusion and allocation.
But the point is, each instance stored into variables do not share any memory.

Although there is no reserved word like `this` or `self`, but the file itself works as a scope designator.

```go
.x:int=0
.y:int=0

move(dx:int,dy:int)=
{
    .x += dx
    .y += dy
}
```
to make x and y private, add `.` front.
This will help distinguishing members to local variables or parameters.

It is recommended to put all members of the class in one side. Methods should be placed below.


> // todo: write some example, and verify if it is plausible.


if there's no value assignment in module definition, it is logically an "interface" of class. Consider how to "implement" it.
do we even need a instance to accomplish it?

maybe we should include interface in the other module as "super" or "prototype" and assign implementation to it?
or maybe we really should make new syntax to accomlish interface-implementation mechanism.
like, real #include.

or implmentation could be done in a function that returns interface type.
so, that means implementation does not need to be a separate file. 

if an imple become a separate file, it also require another instance to use.. that is quite a redundancy.
