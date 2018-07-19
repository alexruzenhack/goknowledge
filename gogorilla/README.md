# Building better Go Web Apps with Gorilla Toolkit

This course shows all the toolkits of Gorilla project. With then you can construct complete web apps.

## Agenda

* Gorilla/Mux
* Gorilla/Reverse
* Gorilla/Context
* Gorilla/Schema
* Gorilla/Securecookie
* Gorilla/RPC
* Gorilla/Sessions
* Gorilla/WebSocket

## Gorilla/Mux

### Parameterized Routes

* Schemes
* PathPrefix
* Path
* Method
* Queries

### Gorilla/Reverse

* Matchers
* Reverse Regexp
* Route Templates

Use reverse to make one source of truth between controller path and template link path.

### Gorilla/Context

Allow to pass data throughout the application layers.

Use router filters to centralize the setting and clearing of a request context.

### Gorilla/Schema

Handle the form binding to Go struct.

You can teach schema's decode to parse other types than just string.

### Gorilla/SecureCookies

Protect information stored in the cookies that are transmitted to and from the client.

### Gorilla/rpc

Types of Web Services

| REST | RPC |
| ---- | --- |
| Tied to HTTP | Protocol Agnostic |
| Resource based | Action based |
| Loosely Coupled | Tightly Coupled |

RPC in GO

| net/rpc | gorilla/rpc |
| ------- | ----------- |
| Single codec per server | Multiple codecs per server |
| Methods take two parameters | Methods take three parameters |
| Persistent connection | One connect per call |

### Gorilla/websocket