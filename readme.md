<div align="center">

# Tite

The **Tite** is a programming language and key-value data exchange format. 

`main: () => print("hello, world")`

</div>

## Key Features of Tite

- Zero keywords to remember: everything is done with some special characters, yet intuitive.
- Familiar syntax: you'd already know most parts of this language if you're a developer.
- Multi-Paradigm: object-oriented or functional programming also was considered in syntax.
- Performance: as fast as C. Tite first compiles to human-readable C.


## Examples

A few basic rules to help you understand:

- variables are declared when `:type=expression` is designated after its name.
- walrus operator `:=` will trigger type inference based on rvalue, like golang.
- identifiers are immutable and private by default unless there is a preceding tag (e.g., `#`, `$`, or `.`).

### primitive types

```rust
a := 5
b: int = 5
c: char = '5'
d: float = 5.5
e: string = "5"
```

### array

```rust
b: int[] = [5, 5] // length inference
c: int[2] = [5, 5]
d: int[2] = [0: 5, 5]
e: int[2] = [0: 5, 1: 5]

print(a[0], a[1]) // prints out "5 5" 
```

*TODO: use `int[string]` type for hashmap?*


### anonymous struct

```rust
a := {.x: 5, .y: 5}
b := {.x:int=5, .y:int=5}
c: {.x:int, .y:int} = {5, 5}
d: {.x:int=0, .y:int=0} = {5, 5} // 5 overrides the default 0
e: {.x:int, .y:int} = {x:5, y:5}
f: {.x:int=0, .y:int=0} = {x:5, y:5}

print(a.x, a.y) // prints out "5 5" 
```

### anonymous function

```rust
() => {} // empty function literal
() => {print("hello")}
(x:int, y:int) => x * y // like JS arrow function
((x:int, y:int) => x * y)() // like JS IIFE to execute
```

### named function

```rust
mul: (a:int, b:int) => a * b
c := mul(5, 5)

div: float(a:int, b:int) => {
    d:float = a
    (d / b) // the last expression becomes a return value
}
e := div(5,3)
```

### named struct (class)

```rust
#point := {.x:int, .y:int} // typedef is done with `#` tag
a: point = {5, 5}

#Point := {    // typedef with a method
    $x:int=0   // `$` tag adds mutability to a variable
    $y:int=0
    .move: (dx:int,dy:int) => {
        x += dx
        y += dy
    }
}
b: Point // {0,0} by default
b.move(5, 5)
```

### immutable class

```rust

#Point := {
    x:int    // methods won't modify fields
    y:int
    .dist: () => {
        // using pythagoras theorem
        (x**2 + y**2)**(1/2)
    }
}
b: Point = {3,4}
print(b.dist()) // prints out "5"
```

### class and access modifiers

You might have noticed that there are `.` tags in front of some identifiers, but some don't.
In fact, this tag act as access modifiers for encapsulation, controlling the visibility.
It means `public` when dots are specified, and `private` when not specified.

```rust
// Note no dots on x & y, while .move & .dist have it
#Point := {
    x:int=0
    y:int=0
    .move:(dx:int, dy:int) => {
        {x + dx, y + dy}
    }
    .dist:() => {
        (x**2 + y**2)**(1/2)
    }
}
a: Point = {x:5, y:5} // x and y are hidden after this point
a.x // compile error
a.y // compile error
a.move // valid access
a.dist // valid access

// In order to fix this, we need `.` in front of desired identifiers.
#point := {.x:int, .y:int}
b: point = {x: 5, y: 5}
b.x // valid access
b.y // valid access
```

*TODO: file scope access modifiers*
