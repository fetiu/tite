# Ideation

## Types and (const) variables

```python
int: {-2147483648 ... 2147483647}

a: 1
a: int
a: int=1
a := 1 // why exists? type inference? for non-consts like stack variables? both are already "a : 1"
// This might essentially mean "a = a : 1" as "a += 1" do, but assigning to a undefined symbol `a` shouldn't generate a valid notion. We might just have to discourage this go compatible expression


string: /".*"/

s: "hello"
s: string
s: string ="hello"


point: {x:int, y:int}

p: point // defaults to 0 fields inherited from 'int'
p: point={0,0}
p: point={x:0, y:0}
// or tentatively like go, for the simplicity of function body definition
p: point{0,0}
p: point{x:0, y:0}
```

## Functions and methods

```python
f: (s: string) => console.log(s)
f: string => console.log(string)
// let unnamed parameter can be referred if it's unique in the context
// this aligns with go's unnamed struct field (e.g. struct{string, int})
// this will also let function notations like "f: X => Y" makes sense
```

As a consequence, we can now derive a principle to disallow reassign to parameter, even if it's a local stack copy persists within the function scope. Because this implies that modifying type wouldn't seem right from the first place. This will also eliminate the misconception that the parameter may be changed and persisted outside the function scope.

Only the stack variables may be modified (tentative: or the mutator called from declared context)

```python
Person:{
  name: string
  setName:(string){
    name = string
  }
}

main:(){
  p: Person
  p.setName("Jake") // allowed
  changeName(p) // error, implicitly changes parameter
  p = mutateName(p)
}

changeName:(p: Person){
  p.setName("David") // not allowed.. because f does not have ownership to `p`
}

mutateName:(p: Person) Person {
  q: Person = p // this is the correct way, explicitly creating a copy.
  q.setName("David")
  q // return
}
```

- [ ] or probably, we can consider an "enter" as implicit union substitute of "|", the only downside is that we need to use explicit comma for every declarations inside a type.
- [ ] this somewhat makes sense, because "|" actually used for concatentation of process, regardless of the success of the previous action.


## Enums

```python
cardsuit: {spades, diamonds, hearts, clubs}
cardsuit: {spades | diamonds | hearts | clubs} // how do we designate the value of identifier
cardsuit: {spades:0, diamonds:1, hearts:2, clubs:3}  // how can we tell whether member or element
cardsuit: {spades:0 | diamonds:1 | hearts:2 | clubs:3} // how can we tell if "0|x" isn't union type -> maybe we can just consider "|" as a chainable operator (like "?"), which has lower priority than ":". So that means ":" also becomes a operator that can be resolved very early stage (but not higher than =, since it's not an "operator"... or just = becomes the highest operator).
// how do we define a nesting type for enums?

cardsuit: uint8_t & {spades:0 | diamonds:1 | hearts:2 | clubs:3} 
cardsuit: uint8_t = {spades:0 | diamonds:1 | hearts:2 | clubs:3} // or a direct realization?
cardsuit:= {spades:0 | diamonds:1 | hearts:2 | clubs:3} // in that case, type inference?

cardsuit: uint8_t = {spades:0, diamonds:1, hearts:2, clubs:3} // or assigning into integral type implicitly becomes union type?
cardsuit: uint8_t = {spades | diamonds | hearts | clubs}

c: cardsuit
c: cardsuit=spades
c: cardsuit=0
```

## Operator priority

```python
a?: 1

a? asdf | b? fdsa

// discern between condition and declaration. 

a? asdf : b? fdsa

// this appears to be the most friendly one.

a? :
// the problem is that ":" here means that it can omit the first operand ....

// making ? as a operator that does not require second operand doesn't make sense...
// if ":" next to "?" considered a non binary operator, this becomes clean, but using :value only makes it work, which doesn't also seem to be a good way to approach this
// so, let's make "?" as a postfix operator that always require a second operand that is whether starting with ":", or any expressions.
```
