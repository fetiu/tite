# Project Dejava

```python
import print from std

def main( ) {
    print("Hello world!");
}
```

Dejava는 `C`의 기능적 단순함을 유지하면서, `Python`, `TypeScript`, `Scala`와 같은 함수형 언어의 간결함을 추구하는 언어이다. (rust나 go와도 다소 비슷할 수 있다. 최신 언어 트렌드를 반영하다보니)

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

심볼에 자료형을 적어주는 것이 아니라, **데이터의 앞에 자료형을 적어준다.**

그리고, 항상 선언시에는 **심볼이 자료형이나 데이터보다 왼쪽에 온다**



시작하기 전에, Dejava의 기본 정신 하나를 소개하려한다.

> 특수 문자를 마음껏 사용해라! 단, 직관적으로.

때때로, 새로운 키워드를 하나 만들어내는 것보다 특수문자 하나를 잘 사용하는 것이 더 직관성을 늘려줄 수 있는걸 보여주고 싶다.



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
    
}
else if
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

 ```python
for i: 'a'~'z' {
    name: name_list.dict(i); #딕셔너리 키로 원소 불러오기
    printf("[%c] %s".fmt(i, name) )
}
 ```



### while

타 언어의 while문과 동일하다.



## Function

함수는 독특한 자료형을 갖는 변수이다.

동시에, 기능을 정의해주어야하는 구조이기도 하다.

```typescript
print(input:str) => (bool) {
    result: system.out.println(input);
    return result;
}
```

최초엔 string형의 구조를 갖고, 최종적으론 bool형의 구조를 갖는다.

함수의 심볼에는 최초에 선언한 시그니처와 일치하지 않는다면, 어떤 값도 대입할 수 없다.



함수는 다음과 같이  프로토타입과 정의를 분리해서 정의할 수도 있다.

```typescript
print: (str) => (bool);

def print(input) {
    result: system.out.println(input);
    return result;
}//print에 input을 넣는상황의 기능을 정의한다.

def print() {
    system.out.println("Helloworld") 
}//아무런 인자도 필요없는 함수의 정의를 내릴수도있다.
```



```

print(input:str) => (bool){
	result: system.out.println(input);
    return result;
} //이보다는 위의 표현이 좋지 않을까? 음... 
----------------------

print(str) => (bool); //이건 어떨까. 함수는 :없이 무조건 =를 사용하는걸
//문제점은, 선언같지 않다는점. print:와 같이 심볼에 대한 선언을 명시해주는 방식이 더 깔끔하지 않을까 싶음.

def print(input) => (bool) {
    
}

=>는 또 다른 대입 연산이니까...

print(input): (str) =>(bool) //input이 string이라는게 명시적이지 않음.
```

TODO: input과 str의 1대1대응이 명시적이지 않음. 표현식 바꿔보기. return형에도 심볼 붙일수있도록.

//헷갈리는건 괜찮다. 무조건 :의 왼쪽에 오는게 심볼이라는 것만 기억한다면



호출형태는 기존 언어와 같다.



## Definition

`def`는 새로운 기능 또는 구조를 정의할 때 사용한다.

```
def pointer (
    data: byte,
    size: uint32_t
)
```

+ 위와 같이 괄호없는 자료형 지정은 def에서만 가능한 문법이다. 



## String format

```python
print("hello world %s".format(username));

```





---

