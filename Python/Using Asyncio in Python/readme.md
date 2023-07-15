# Using Asyncio in Python

[책 링크](https://product.kyobobook.co.kr/detail/S000001810381)

![image](https://github.com/junah201/self-study/assets/75025529/c79baf3b-51f3-43cb-b3de-74320b137ca0)

### 목적

파이썬에서 비동기에 대해서 디테일하게 모르고 넘어갔던 부분에 대해서 학습하기

### 학습 상태

- [x] CHAPTER 1 Asyncio 소개
- [ ] CHAPTER 2 스레드에 관한 진실
- [ ] CHAPTER 3 Asyncio 공략
- [ ] CHAPTER 4 여러분이 사용하지 않는 Asyncio 라이브러리 20개
- [ ] CHAPTER 5 마치며
- [ ] APPENDIX A 파이썬의 비동기 지원에 대한 역사
- [ ] APPENDIX B 보충 자료

### 학습 메모

학습 도중 새롭게 알게된 사실에 대해서 메모합니다.

**CHAPTER 1 Asyncio 소개**

- CPU는 작업을 완료한 후 네트워크 I/O 작업을 기다린다. CPU 작업은 네트워크 작업보다 10만 배 이상 빠르다.
- I/O 위주 작업에 스레드 기반 병행 처리보다 비동기 기반 병행처리를 적용해야 하는 이유
  - Asyncio는 스레드를 사용하는 선점형 멀티태스킹보다 오류, 경합조건, 혹은 비결정론적 위험 요소(nondeterministic danger)에 더욱 안전하다.
  - Asyncio를 통해 **동시에** 수천 개의 소켓 연결(WebSocket, MQTT)을 간단히 처리할 수 있다.
- Asyncio를 통해 GIL를 해결할 수 있는 것이 아니라, GIL의 문제는 멀티스레드를 사용할 때 멀티코어 병렬화를 막는 것이지만, Asyncio는 명목상 단일 스레드이기 때문에 GIL의 영향을 받지 않는 것이다.
