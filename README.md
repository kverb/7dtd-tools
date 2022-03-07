# 7dtd-tools
small programs to assist with running a 7 Days To Die dedicated serve


## Usage

Current code as of Mar 07 2022 just launches a go http server.
A listening port can be specified as the first argument to the program.
The api is simple. There is just the root, and 3 possible query parameters:
  host - the hostname or IP address of the gameserver that you want to query
  port - the (TCP) port that the gameserver listens on
  filter - [optional] you can exclude all other info parameters except one specified here.


### examples

Run the http server:

```
$ go run main.go
```
And make a request to the server:

```
$ curl http://localhost:8787/?host=my7dtdserver.com&port=26900
{
    ... <JSON blob truncated>
}

$ curl http://localhost:8787/?host=my7dtdserver.com&port=26900&filter=DayCount
{
  "DayCount": "3"
}
```
