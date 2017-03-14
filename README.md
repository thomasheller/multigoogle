# multigoogle

multigoogle starts multiple Google searches at once.

Its syntax is very similar to shell brace expansion: Simply use
regular braces instead of curly braces. Furthermore, multigoogle
treats single elements inside braces as optional.

## Install

```sh
$ go get github.com/thomasheller/multigoogle
$ alias mgo='noglob multigoogle'
```

## Examples

```sh
$ mgo San Francisco,Berlin
```
Searches for "San Francisco" and "Berlin".

```sh
$ mgo Go (San Francisco,Berlin)
```
Searches for "Go San Francisco" and "Go Berlin".

```sh
$ mgo (Go,Golang) (San Francisco,Berlin)
```
Searches for "Go San Francisco", "Golang San Francisco", "Go Berlin" and "Golang Berlin".

```sh
$ mgo San Francisco (Golang)
```
Searches for "San Francisco" and "San Francisco Golang".

```sh
$ mgo (San Francisco,Berlin) (Golang)
```
Searches for "San Francisco" and "San Francisco Golang", "Berlin" and "Berlin Golang".
