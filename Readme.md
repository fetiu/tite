# Project Dejava

```python
import print from std

def main( ) {
    print("Hello world!");
}
```

Dejava는 `C`의 기능적 단순함을 유지하면서, `Python`의 간결함과, `TypeScript`의 정적 타이핑, `Scala`와 같은 함수형 언어의 특징을 추구하는 언어이다. (rust나 go와도 다소 비슷할 수 있다. 최신 언어 트렌드를 반영하다보니)

> 즉, 인간과 컴퓨터 모두 읽기 쉬운 언어를 목표로 한다.

```ruby
for item: 1~10{
    print("Looped for %d times".fmt(item))
}
```

10번의 루프는 이렇게 사용한다.



가장 중요한 변수의 선언은, 다음과 같이 한다.

```go
item1: 5
item2: (int) 5
```

직관적으로 알 수 있듯이, 변수의 자료형은 선택사항이다. 

변수명에 자료형을 적어주는 것이 아니라, **데이터의 앞에 자료형을 적어준다.**

그리고, 항상 선언시에는 **변수명이 자료형이나 데이터보다 왼쪽에 온다**



시작하기 전에, Dejava의 기본 정신 하나를 소개하려한다.

> 특수 문자를 마음껏 사용해라! 단, 직관적으로.

때때로, 새로운 키워드를 하나 만들어내는 것보다 특수문자 하나를 잘 사용하는 것이 더 직관성을 늘려줄 수 있는걸 보여주고 싶다.



변수명, 함수명 등은 `심볼`, 또는 `라벨`이라 통칭한다.



## Primitive data syntax & types

몇몇 원시타입이 
표현 방식은 다소 변경될 수 있다.
```rust
//number
int 1
/*
int, uint
float, double
...most c types compatible
*/

//character
char 'a'

//string
str "hello"

//instant (unnamed) function
(a:int) => (int){ return a } //should be having parameter name.

//struct
["hello",5,"world"]
{ str, int, str } //굳이 괄호가 다를 필요있을까? 같아도 의미통하는데.. 
//혹시라도 자료형과 혼동가능한 변수를 넣을수도 있으니까.

//instant struct with member name
{x:int, y:int} [x=5, y=6]
{x:int, y:int} [5, 6]


//array, list
[1,2,3,4] //메모리의 형태를 모사한다고 생각?
int[] //[]로 묶인 하나의 원소만 있는 배열이랑 혼동되지않을까? 
int array
int*N
(int,int,int,int) //const array?
//or vector? list type needed
[int]

//dictionary (tentative), hash
{"hello": }

//enum
{A;B;C;D}
```

괄호 닫는것도 자료형도 데이터하고 똑같이 표현해주면 될듯.

{}를 구조체와 배열에 사용하는건 혼동을 불러일으키려나?

그리고 컨벤션에 조금 안맞음. 일반적으로 복합구조체들은 {}로 묶는 경향이 있음.



## Condition

조건문의 시작은 `if`또는 `?`로 시작된다.

```c
item1 == 5 ?
	print("item is 5!")
else
    print("item is not 5");
```

삼항 연산을 사용해 본사람이라면 꽤나 쉽게 이해할 수 있으리라 생각된다.

조건 식을 나열하기 위해 사용했던 :가 else로 대체 되었을 뿐이다.



```
if (item1 = 5 || 6) {
    print("item is 5 or 6");
} else if ( item = 0 ){
	print("item is 0");
    break;
}
```

if는 조건을 만들어주는 명령어이다.  인자로 제공되는 표현식이 참임을 검증하여 bool값을 반환해준다.

>  TODO: if "item1 = 5 || 6"도 고민해보기. 만든 문자열 조건식을 if에 바로 변수로 제공해도 좋을듯
>
> if "{} = 'hello' ".fmt(word) {} 이렇게. 문자열 조건식.
>
> 필요성은? 문자열 비교를 위해서? 바로 변수명 $이걸로 써주는건 어때



## Loop

`for` 과`while`두가지가 있다.

### for

```python
name_list: ["john","mike","andrew","sam"]

for name: name_list{
    printf("hello %s".fmt(name));
}
```

 먼저 리스트의 원소를 담기위한 새로운 인스턴스 변수를 선언하고,

가장 마지막에 기입되는 리스트형의 구조에서 원소를 하나씩 불러들여서 변수에 넣는다.

리스트의 마지막 원소에 도달할때까지 반복한다.

이 원칙 때문에, 상수 번의 루프 순회를 위해서 새로운 리스트를 만들어야하는 부담감이 있을것이다.

이때 유용한 문법이 `리스트 제너레이터`이다. 연산 가능한 범위의 primitive type의 value두개의 사이에

`~`를 집어넣으면, 새로운 리스트 인스턴스가 생성된다.

 ```python
for i: 'a'~'z' {
    name: name_list.dict(i); #딕셔너리 키로 원소 불러오기
    printf("[%c] %s".fmt(i, name) )
}
 ```



### while

타 언어의 while문과 동일하다.



## Definition

`def`는 새로운 기능 또는 구조를 정의할 때 사용한다.

```
def pointer {
    data: byte,
    size: uint32_t
}

def myfuction ()=>() {
    print("myfunction");
}

```

- 위와 같이 괄호없는 자료형 지정은 def에서만 가능한 문법이다.  (현재 변수명도 괄호없이 자료형 지정 고려중. )



## Function

함수는 변수도, 자료형도 아니다.

마치 자료형 처럼 입력할 데이터가 정해져있지 않다는 특징이 있지만, 정의된 일련의 과정은 마치 내용이 담긴 변수와 같은 느낌을준다.

함수는 고유의 독특한 구조를 가지고있다.

```typescript
//익명 함수 선언
( input: str ) => (bool) {
    result: system.out.println(input);
    return result;
}
```

위의 함수는 Dejava함수의 기본 데이터 구조이자, 익명함수의 선언이다.

최초엔 string형의 구조를 받아서, 최종적으론 bool형의 구조를 출력으로 내놓는다.

함수의 모든 매개변수와 반환값은 ()와 =>를 통해 관계가 명시되어야한다.



위의 익명 함수 선언은 시그니처와 기능은 정의되었지만 아직 호출하기 위한 함수명이 없기 때문에 함수의 역할을 하지 못한다.

때문에 위 같은 함수에는 `심볼`을 지정하여 추후에 부를 수 있도록 해야한다.

총 두가지 방법이 있다.

```scala
//익명함수에 define을 통해서 명칭을 부여하기. -> 정적함수
def print1 (input:str) => (bool) {
    result: system.out.println(input);
    return result;
}

//익명함수를 담는 (함수포인터) 변수 선언하기. -> 동적함수
print2: (input:str) => (bool) {
    result: system.out.println(input);
    return result;
}

```

위 두가지 표현 모두 호출 가능한 함수를 만들어 낼 수 있다.

차이점은, 전자는 변경 불가능한 정적 함수 선언이고,

후자는 시그니처가 일치하는 다른 함수로 대체할 수 있는 동적 함수 선언이다.

 

 함수는 다음과 같이 바로 기능을 정의하지 않고, 프로토타입과 정의를 분리해서 정의할 수도 있다.

```scala
/* 정적 함수 */
def print1 (str) => (bool); //프로토타입

def print1 (input) { //위에서 정의가 내려진 시그니처를 생략하고 작성 가능하다.
    return system.out.println(input);
}//print에 input을 넣는상황의 기능을 정의한다.


/* 동적 함수 */
print: (str) => (bool); //프로토타입

//동적함수의 프로토타입에 정의를 집어넣을 때는 수작업이 필요하다.
print = print1;
print("hello"); //it will print "hello"


def print2() {
    system.out.println("Helloworld") 
}//아무런 인자도 필요없는 함수의 정의

print = print2;//시그니처가 맞지 않아 실패한다.


print = (message:str) => (bool)  {
    return system.out.println("Hello world %s".fmt(message));
} //새로운 익명함수를 대입한다.

print("bro"); //it will print "Hello world bro"

```

위와 같은 방법을 사용하면, 가독성과 확장성 면에서 이점을 얻을 수 있다.

동적함수는 특히 클래스나 모듈같은 복합구조 안에 메소드를 집어넣을 때 효과적이다.

- [ ] return 매개변수명을 지정할 수 있도록 할것인가?



함수에서는 매개변수 명이 중요한 역할을 한다. 함수를 정의하는 구간에서는 분명하게 그 매개변수의 이름과 자료형을 알 수 있어야한다. 함수의 프로토타입이 아니라, {}로 묶인 정의부의 앞에 매개변수명을 적어줘야 함에 유의하자.

함수를 호출하는 구간에서는 매개변수 명은 상관없다. 호출형태는 기존 언어와 같다.



## String format

```python
print("hello world %s".fmt(username));

```





---

