# 노마드코더 - 바닐라 JS로 크롬 앱 만들기

[강의 링크](https://nomadcoders.co/javascript-for-beginners)

![image](https://user-images.githubusercontent.com/75025529/208254925-920c7a32-395c-4603-9d35-ca6094e24c80.png)

### 목적

html css 부분에서 단순히 구글링(야매)로 해결하고 넘어갔던 부분를 제대로 익히기 (기초다지기)

### 학습 상태

- [x] #1 [2021 Update] Introduction
- [x] #2 [2021 Update] Welcome to Javascript
- [x] #3 [2021 Update] Javascript on the Browser
- [x] #4 [2021 Update] Login
- [x] #5 [2021 Update] Clock
- [x] #6 [2021 Update] Quotes and Background
- [x] #7 [2021 Update] To Do List
- [x] #8 [2021 Update] Weather

### 학습 메모

학습 도중 새롭게 알게된 사실에 대해서 메모합니다.

**3. Javascript on the Browser**

- element.classList.toggle("className") : 해당 클래스가 있으면 빼고, 없으면 추가

**4. Login**

- event.preventDefault() : 브라우저의 기본 동작 막기
- localStorage.setItem("key", "value")
- localStorage.getItem("key")
- localStorage.removeItem("key")

**5. Clock**

- setInterval(함수, ms) : 해당 함수를 지정한 ms마다 실행합니다.
- setTimeout(함수, ms) : 해당 함수를 지정한 ms 이후에 한번만 실행합니다.
- string.padStart(N, "K") : 해당 string의 길이가 N보다 작으면 string에 앞에 "K"를 붙여 길이를 N개로 만듭니다.
- string.padEnd(N, "K") : 해당 string의 길이가 N보다 작으면 string에 뒤에 "K"를 붙여 길이를 N개로 만듭니다.

**6. Quotes and Background**

- Math.random() : 0 에서 1 사이의 랜덤 값
- Math.round(숫자) : 해당 숫자를 반올림합니다.
- Math.ceil(숫자) : 해당 숫자를 올림합니다.
- Math.floor(숫자) : 해당 숫자를 내림합니다.
- document.createElement("tag") : 해당 tag의 Element를 생성합니다.

**7. To Do List**

- event.event.target.parentElement : target의 부모 요소
- element.remove() : 해당 Element를 삭제합니다.
- JSON.stringify(anyObject) : 해당 오브젝트를 문자열화 합니다.
- JSON.parse("string") : 해당 문자열을 오브젝트화 합니다.
- parseInt("string") : 해당 문자열을 int로 변환합니다.
- List.filter(anyFunction) : 해당 List에서 anyFunction에 값을 하나 씩 넣고 이후 리턴 값이 true 인 값만으로 모은 List를 반환합니다.

**8. Weather**

- navigator.geolocation.getCurrentPosition(successCallback, errorCallback) :
- fetch(url) : url에 get 요청
