# Using Asyncio in Python

[책 링크](https://product.kyobobook.co.kr/detail/S000001810381)

![image](https://github.com/junah201/self-study/assets/75025529/c79baf3b-51f3-43cb-b3de-74320b137ca0)

### 목적

파이썬에서 비동기에 대해서 디테일하게 모르고 넘어갔던 부분에 대해서 학습하기

### 학습 상태

- [x] CHAPTER 1 Asyncio 소개
- [x] CHAPTER 2 스레드에 관한 진실
- [x] CHAPTER 3 Asyncio 공략
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

**CHAPTER 2 스레드에 관한 진실**

- 스레딩의 장점

  - 읽기 쉬운 코드 (코드가 top-down 방식이며, 함수 내에서 다른 코드 혹은 함수의 동시 실행에 대해 고려할 필요가 없다.)
  - 공유 메모리를 통한 병렬처리 (스레드 간 공유 메모리를 통해 통신하면서 코드에서 여러 CPU를 이용할 수 있다.)
    - 하지만 파이썬에서는 GIL로 인해서 병렬성이 제한됨
  - 많은 모범 사례 및 노하우

- 스레딩의 단점

  - 스레딩은 어렵다. (스레드 관련 오류나 경합 조건은 가장 고치기 어렵다.)
  - 스레드는 더 많은 OS 자원을 사용한다.
  - 스레딩은 처리량에 영향을 줄 수 있다. (매우 높은 병행 수준(5000개 이상의 스레드)에서는 콘텍스트 전환 비용으로 인해 처리량에 영향이 있을 수 있다.)
  - 스레딩은 유연하지 않다. (ex, 일부 스레드는 소켓으로부터 데이터가 도달하기를 기다리고 있을 수도 있다. 하지만 OS 스케줄러는 데이터가 도달하여 해당 스레드의 실행 재개가 필요하기 전까지 수천 번에 거처 콘텍스트 전환을 의밓 없이 수행할 것이다.)

- 경합 조건 오류의 예시
  - 스레딩 연산에서 `+=` 연산에 경우 내부적으로 여러 단계로 구성되어 있다.
    1. 원래 변수 값을 읽어 임시 저장소에 저장한다.
    2. 임시 저장소 내의 값에 더할 값을 더한다.
    3. 원래 변수의 값에 임시 저장소 내의 값을 복제하여 저장한다.
  - 내부적으로 `+=`이 진행되는 중간에 스레드 간 콘텍스트 전환이 발생하면, 임시 저장소에 저장된 값이 유실되는 경합 조건 오류가 발생할 수 있다.
  - 이러한 경합 조건 오류는 소스 코드만 확인해서는 찾아내기 매우 힘들다. (OS는 거의 모든 곳에서 스레드 간의 콘텍스트 전환을 수행할 수 있다.)

**CHAPTER 3 Asyncio 공략**

- Asyncio를 사용한 `Hello World`

```py
import asyncio, time

async def main():
  print(f"{time.ctime()} Hello!")
  await asyncio.sleep(1.0)
  print(f"{time.ctime()} Goodbye!")

asyncio.run(main())
```

- 일반적으로는 asyncio 모듈에서 제공하는 run() 함수로 asyncio 함수를 실행한다.

- run() 함수 내부의 모든 동작에 대한 거의 완벽한 예제

```py
import asyncio, time

async def main():
  print(f"{time.ctime()} Hello!")
  await asyncio.sleep(1.0)
  print(f"{time.ctime()} Goodbye!")

loop = asyncio.get_event_loop()  # 1
task = loop.create_task(main())  # 2
loop.run_until_complete(task)    # 3
pending = asyncio.all_tasks(loop=loop)
for task in pending:
  task.cancel()
group = asyncio.gather(*pending, return_exceptions=True)  # 4
loop.run_until_complete(group)  # 3
loop.close()  # 5
```

- 1 : `loop = asyncio.get_event_loop()` 코루틴을 실행하기 위한 루프 인스턴스를 얻는 문법
- 2 : `task = loop.create_task(main())`를 하기 전까지 코루틴 함수를 실행되지 않는다. 반환받은 task 객체를 통해 작업의 상태를 모니터링 할 수 있고, 코루틴 완료 후 반환 값을 얻을 수도 있으며, 중간에 `task.cancel()`을 이용해서 작업을 취소할 수도 있다.
- 3 : `loop.run_until_complete(task)` 을 통해 현재 스레드를 **블로킹** 할 수 있다.'
- 4 : process signal이나 loop.stop()의 호출로 인한 루프 중지 등으로 `main` 블로킹 상태가 풀린 후 (위 예제에서는 `await asyncio.sleep(1.0)`이 반환된 후) run_until_complete() 이후의 코드가 실행된다. 이후 코드의 절차는 아직 실행되는 태스크를 취합하고, 모든 태스크에게 취소 요청을 한 후 loop.run_until_complete(group)를 호출하여 태스크들이 모두 종료 상태가 될 때까지 대기한다.
- 5 : `loop.close()`는 보통 최종 동작으로, 루트의 모든 대기열을 비우고, executor를 종료시킨다.

- 위 모든 과정은 `asyncio.run()`으로 추상화되어 있다.

- `코루틴` 이란

  - `async def`로 선언된 함수의 type은 코루틴이 아니라 함수이고, 정확하게 말하면 코루틴 함수이다.
  - 코루틴이란 완료되지 않는 채 일시 정지했던 함수를 재개할 수 있는 기능을 가진 객체이다.
  - 즉 `async def`로 선언된 함수 자체가 코루틴이 아니라, `async def`의 반환 값이 코루틴이다. (반환 값이 함수의 retrun 값이 아니라, 함수가 실행되기전 함수가 반환하는 값임)

- awaitable 이란

  - 코루틴
  - `__await__()`라는 특별 메서드를 구현한 모든 객체 (일상적인 프로그래밍에서 사용할 일이 거의 없음)

- 코루틴 내부에 send()와 StopIteration 사용하기

```py
async def f():
  return 123

core = f()
try:
  core.send(None) # 1
except StopIteration as e:
  print('The answer was :', e.value) # 2
```

- 1: 코루틴에 None을 send하여 초기화한다.
- 2: 코루틴이 **반환**될 때 StopIteration라는 특별한 예외가 발생한다.

- 코루틴에 직접 예외를 주입해보기

```py
>>> core = f() #1
>>> core.send(None)
>>> core.throw(Exception, 'blsh') # 2
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "<stdin>", line 2, in f
Exception: blah
blash
```

- 1: 코루틴 함수 f()를 이용해서 코루틴 객체를 생성함
- 2: core.throw()를 호출하여 예외 클래스와 값을 전달한다.

- async with : 네트워크 연결과 같은 네트워크 자원의 생명주기를 적절히 정의한 범위 내에서 관리하고자 할때 async with은 매우 편리하다.

```py
class Connection:
  def __init__(self, host, port):
    self.host = host
    self.port = port
  async def __aenter__(self):
    self.conn = await get_conn(self.host, self.port)
    return self.conn
  async def __aexit(self, exc_type, exc, tb):
    await self.conn.close()

async with Connection('localhost', 9001) as conn:
  # doing something using conn
```

- 동기 콘텍스트 관리자를 위한 특별 메서드인 **enter**() 대신, **aenter**() 라는 새로운 특별 메서드를 사용함.
- 동기 콘텍스트 관리자를 위한 특별 메서드인 **exit**() 대신, **aexit**() 라는 새로운 특별 메서드를 사용함.

- 전통적인 이터레이터

```py
class A:
  def __iter__(self):
    self.x = 0
    return self
  def __next(self):
    if self.x > 2:
      raise StopIteraion
    self.x += 1
    return self.x

for i in A():
  print(i)
```

- 이터레이터는 **iter**()라는 특별 메서드를 구현하고, 이 메서드는 iterable을 반환해야 한다.
- iterable은 **next**() 특별 메서드를 구현한 객체이다.
- **next**() 는 반복에 모든 단계에서 StopIteration이 발생할 때 까지 계속 호출되며, 각 단계마다 반환값을 반환단다.

- 비동기 이터레이터 (async for)

  - def **aiter**()를 구현해야한다. (async def가 아님)
  - **aiter**()는 asycn def \_\_anext()를 구현한 객체를 반환해야함.
  - **anext()**는 반복의 각 단계에 대한 값을 반환하고, 반복이 끝나면 StopAsyncIteration을 발생시켜야함.

```py
import asyncio
from aioredis import create_redis

async def main():
  redis = await create_redis(('localhost', 6379))
  keys = ['Americas', 'Africas', 'Europe', 'Asia']

  async for value in OneAtATime(redis, keys):
    await do_something_with(value)

class OneAtTime:
  def __init__(self, redis, keys):
    self.redis = redis
    self.keys = keys
  def __aiter__(self):
    self.ikeys = iter(self.keys)
    return self
  async def __anext__(self):
    try:
      k = next(slef.ikeys)
    except StopIteration:
      raise StopAsyncIteration

    value = await self.redis.get(k)
    return value

  asyncio.run(main())
```

- async for를 사용한다. 중요한 점은 반복 중에 다음 데이터를 얻기 전까지 반복 자체를 일시 정지 할 수 있음
- 키에 연관된 데이터를 레디스에서 가져오는 동안 await 하여, 네트워크 I/O 동작이 완료되길 기다리는 동안 이벤트 루프에서 다른 동작이 수행되게 할 수 있다.
- async for를 통해 for 루프의 편의성은 유지하면서, 데이터를 가져오는 I/O를 반복 수행하는 동작은 비동기로 처리할 수 있다. 하나의 이벤트 루프만으로 엄청난 양의 데이터를 처리할 수 있다.

- 비동기 제너레이터를 사용한 더 간단한 코드

```py
import asyncio
from aioredis import create_redis

async def main():
  redis = await create_redis(('localhost', 6379))
  keys = ['Americas', 'Africas', 'Europe', 'Asia']redis = await create_redis(())

  async for value in one_at_a_time(redis, keys):
    await do_something_with(value)

async def one_at_a_time(redis, keys):
  for k in keys:
    value = await redis.get(k)
    yield value

asyncio.run(main())
```

- 비동기 이터레이터를 이용한 코드에서 필요했던 self.ikeys 부분이 필요가 없다. 키를 직접 전달하고 값을 받는다.
