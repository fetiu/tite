# `diccle`

The Diccle is a programming language inspired by [compound literals](https://en.cppreference.com/w/c/language/compound_literal) of C99.

```rust
main:(){
    print("hello, world")
}
```

Diccle is a recursive acronym for "**D**iccle **i**sn't **C** **c**ompound **l**iteral **e**xpression". 

At the same time, Diccle (디끌) is a southeastern dialect of Digeut (ㄷ), the 3rd letter of the Korean alphabet.

## Key Features of Diccle

- No keywords to remember: everything is done with some special characters, yet intuitive.
- Familiar syntax: you'd already know most parts of this language (if you're a developer).
- Multi-Paradigm: object-oriented or functional programming also was considered in syntax.
- Performance: as fast as C. Diccle first compiles to human-readable C.


## Examples

A few basic rules to help you understand:

- variables are declared when `:typename` is designated. 
- `:=` will triger type inference based on rvalue, like golang.
- both array and struct use `{}` with [designated inits](https://gcc.gnu.org/onlinedocs/gcc/Designated-Inits.html), like [carbon-lang](https://github.com/carbon-language/carbon-lang)

### primitive types

```rust
a := 5
b:int = 5
c:char = '5'
d:float = 5.5
e:string = "5"
```

### array

```rust
a := {5, 5}
b:int[] = {5, 5} // length inference
c:int[2] = {5, 5}
d:int[2] = {[0]=5, 5}
e:int[2] = {[0]=5, [1]=5}
```

### unnamed struct
```rust
a := {.x:=5, .y:=5}
b := {.x:int=5, .y:int=5}
c:{.x:int, .y:int} = {5, 5}
d:{.x:int=0, .y:int=0} = {5, 5} // 5 overrides the default 0
e:{.x:int, .y:int} = {.x=5, .y=5}
f:{.x:int=0, .y:int=0} = {.x=5, .y=5}
```

### immutable struct
```rust
a := (.x:=5, .y:=5)
b := (.x:int=5, .y:int=5)
c:(.x:int, .y:int) = (5, 5)
e:(.x:int, .y:int) = (.x=5, .y=5)
```

### unnamed tuple
```rust
a := (5, 5)
b := ([0]=5, [1]=5)
c := ([0]:int=5, [1]:float=5)
d:(int, float) = (5, 5)
e:(int, float) = ([0]=5, [1]=5)
```

### unnamed function

```rust
(){} // empty function
(){print("hello")}
(x:int, y:int) => x * y // like JS arrow function
(x:int, y:int){=> x * y} // now `=>` replaces `return`
((x:int, y:int) => x * y)() // like JS IIFE to execute
```

### named function

```rust
mul:(a:int, b:int) => a * b
c := mul(5, 5)

div:(a:int, b:int){
    d:float = a
    // arrow goes inside for multi-line body
    => (d / b)
}
e := div(5,3)
```

### named struct (class)

```rust
point:{.x:int, .y:int}() // typedef
a:point{5, 5}

Point:{x:int=0, y:int=0}( // typedef with a method
    .move:(dx:int,dy:int){
        x += dx
        y += dy
    }
)
b:Point{} // {0,0} by default
b.move(5, 5)
```

### class and access modifiers

Now you might have noticed that some identifiers have a `.` in front, some don't.
In fact, these dots act as access modifiers for encapsulation.

```rust
// Note no dots on x & y, but .move has it
Point:{x:int=0, y:int=0}(
    .move:(dx:int,dy:int){
        x += dx
        y += dy
    }
)
a:Point{x=5,y=5} // x and y are hidden after this point
a.x // compile error
a.y // compile error
a.move // valid access

// In order to fix this, we need `.` in front of desired identifiers.
point:{.x:int, .y:int}()
b:point{.x=5, .y=5}
b.x // valid access
b.y // valid access
```
