编一个TCP并发服务器，可同时支持N个客户端访问。服务器接收客户端发送内容，将内容按单词逆置，回发给客户端。

如：客户端发送：this is a socket test        服务器回复：test socket a is this