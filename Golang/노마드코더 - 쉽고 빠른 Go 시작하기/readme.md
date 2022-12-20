# 노마드코더 - 쉽고 빠른 Go 시작하기

[강의 링크](https://nomadcoders.co/go-for-beginners)

![1671546491620](image/readme/1671546491620.png)

### 목적

Go로 서버 열고, 병렬화 배워보기 전에 Go 기초 다지기

### 학습 상태

- [X] #0 INTRODUCTION
- [X] #1 THEORY
- [ ] #2 BANK & DICTIONARY PROJECTS
- [ ] #3 URL CHECKER & GO ROUTINES
- [ ] #4 JOB SCRAPPER
- [ ] #5 WEB SERVER WITH ECHO

### 학습 메모

학습 도중 새롭게 알게된 사실에 대해서 메모합니다.

**1. THEORY**

- 함수명을 대문자로 시작해야만 다른 파일에서 호출할 수 있다.
- 함수명을 소문자로 시작하면 다른 파일에서 호출하지 못한다. (Python의 \_ 로 시작하는 함수와 비슷)
- 상수 선언법 : `const 변수명 변수타입 = 값` (ex, `const name string = "Junah"`)
- 변수 선언법 : `var 변수명 변수타입 = 값` (ex, `var name string = "Junah"`)
- 축약형 변수 선언법 : `변수명 := 값` (ex, `name := "Junah"`)
- 축약형 변수는 함수 밖에서 선언 불가능
- 함수에서는 매개변수에도 타입을 지정해주어야하고, 리턴값에도 타입을 지정해줘야한다.
- 만약 매개변수가 모두 같은 타입이면 `func funcname(a int, b int)` ->`func funcname(a, b int)`으로 축약할 수 있다.
- 리턴값이 여러개이면 리턴형을 `int` -> `(int, int)` 이런 식으로 쓴다.
- 파이썬에서와 같이 리턴값을 받지 않을려면 변수명을 `_`로 하면 된다.
- 하나의 매개변수에 여러가지 값을 받으려면 변수타입 앞에 `...`을 붙이면 된다. (ex, `func repeatMe(words ...string)`)
- "naked" return : `return` 키워드 뒤에 값을 적지 않고도, 리턴형을 지정하는 부분에서 타입 앞에 리턴할 변수명을 적어주고, 함수 내부에서 해당 변수를 선언하면 자동으로 그 변수가 리턴된다.
- "defer" : `defer` 키워드 뒤에 적은 코드는 해당 함수가 끝나면 자동으로 실행된다.
- if 조건문 내에서 변수 선언이 가능하다. (ex, `if KoreanAge := age + 2; koreanAge < 18 {실행문}`)
- switch 내에서 변수 선언이 가능하다. (ex, `switch KoreanAge := age + 2; koreanAge { case문 }`)
- 마치 C처럼 변수 앞에 `&`를 붙이면 메모리 주소 확인할 수 있다.
- 메모리 주소 변수 앞에 `*`를 붙이면 해당 메모리에 값을 확인할 수 있다.
- 정적 배열 선언 : `변수명 := [크기]타입 {값들}` (ex, `names:= [5]string {"Junah", "Merry", "dal", "japari", "flynn"}`)
- 동적 배열 선언 : `변수명 := []타입 {값들}` (ex, `names:= []string {"Junah", "Merry", "dal", "japari", "flynn"}`)
- 동적 배열에 값 추가 : `append(배열, 값)` (주의 : 값을 직접 수정하지 않고 수정된 값을 리턴함)
- map 변수 선언법 : `변수명 := map[키타입]값타입{키1 : 값1, 키2 : 값2}`
- map 변수 for문 사용법 : `for key, value := range 변수명 {실행문}`
- struct 선언법 : `type 구조체명 struct {내부변수명1 타입1}`
- struct 사용법 : `변수명 := 구조체명{내부변수명1 : 내부변수값1, 내부변수명2 : 내부변수값2}`
- struct 내부 요소 접근 : `변수명.필드명`
