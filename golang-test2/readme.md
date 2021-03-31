# Mission 2 - Please write a microservice to search movies from http://www.omdbapi.com/
## About

The microservice should be able to handle two transports : REST JSON HTTP and GRPC Access credentials :
OMDBKey : "faf7e5bb&s"
URL : http://www.omdbapi.com/
* Example url call to search is --> GET http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2 Functions to be implemented are :
- Search with pagination --> 2 parameters : "pagination" and "searchword"
- Get single detail of the movie
- Log each search calls to a dummy DB eg. let's just say we have a MySQL DB table for this. Important aspects :
- Readability of code
- Good display on the knowledge of "Separation of Concerns for Codes"
- Write unit tests on some of the important files. (For Bigger plus points see below)
- Good use of asynchronousy with Go-Routine
Plus points:
- Implementation of Clean Architecture is a BIG plus
- Complete Unit tests on all codes


## Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/) [mac]:

```Shell
sudo brew install go
```

with [apt](http://packages.qa.debian.org/a/apt.html)-get [linux]:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)

## Compile
that you can start go applications via 
```Shell
go run main.go
```
but also compile it to an executable with 
```Shell
go build main.go
```

## License
ISC
