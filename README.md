# Simple Golang URL Shortener
This is simple example Golang URL Shortener with restful api server only with **gorilla/mux**.  

## Install and Run
```shell
$ go get github.com/atrasnajwan/Simple-Golang-URL-Shortener

$ cd $GOPATH/src/github.com/atrasnajwan/Simple-Golang-URL-Shortener
$ go build
$ ./Simple-Golang-URL-Shortener
```

## API Endpoint
- http://localhost:8000/create
    - `POST`: create shorten URL

## Data Structure
```json
{"longurl":"https://github.com/atrasnajwan"}
```

## Data Output
```json
{
	"id":"iCMRAjWw",
	"shorturl":"localhost:8000/iCMRAjWw",
	"longurl":"https://github.com/atrasnajwan"
}
```