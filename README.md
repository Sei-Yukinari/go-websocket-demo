# go-websocket-demo
Websocket & Redis Pubsub Golang untuk Chatting

## Libraries used
1. Redis
2. Gin Framework
3. Gorilla Websocket

## Usage

```bash
$ make start
```

http://localhost:13333/

```bash
$ make stop
```


## Flow
```mermaid
graph TD
    A[client-1] ---|1.Websocket<br>join| B{Server}
    A --> |3.send message|B
    B --> |2.Subscribe|C{Redis}
    B --- |1.websocket<br>join|D[client-2]
    B --> |6.recieve message|D
    B --> |4.Publish|C
    C --> |5.Message|B
    B --> |6.recieve message|A
```


### reference
https://medium.com/@thomi.algh/websocket-redis-pubsub-golang-untuk-chat-80513b880550