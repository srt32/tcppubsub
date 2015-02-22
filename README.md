A sample TCP server and client in Go

Open a terminal window and run (this process is our server):
`cd tcpsub; go run tcpsub.go`

Open another terminal window and run (this process is our client):
`cd tcppub; go run tcppub.go`

From the second process type a word, such as "chicken". You should see `Sending
response for 8 bytes` in the server window.

Next steps would be to manage multiple clients and create a small chat service.


References:

- [http://golang.org/pkg/net/](http://golang.org/pkg/net/)
- [https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go](https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go)
- [http://www.badgerr.co.uk/2011/06/20/golang-away-tcp-chat-server/](http://www.badgerr.co.uk/2011/06/20/golang-away-tcp-chat-server/)
