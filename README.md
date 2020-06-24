# Coding Task

Write an command line tool for receiving dummy post data from the api endpoint JSONPlaceholder (https://jsonplaceholder.typicode.com)

## Build

Clone repository and change current directory to go-coding-task/src. Run `go build -o ../bin/go-cli`. This creates a binary file in the folder `go-coding-task/bin`.

## Usage 

`./bin/go-clie -userId :userId [-filter :filter]`
* userId: must be of type int
* filter: filters comments from the post data
