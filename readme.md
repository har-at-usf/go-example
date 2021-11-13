# Go-Example

Working example of the write-up posted on my blog.

# Installation

1.  Set GOPATH to this (root) folder. On Unix-like systems:

        export GOPATH=.

    In Powershell, you can crudely do this in the root directory with:

        $env:GOPATH=Get-Location

    Test it with:

        go env GOPATH

2.  Build the interface module.

        go build geometry

3.  Install the main module.

        go install demo

4.  Run the program:

        ./bin/main

# References

https://gobyexample.com/interfaces

https://go.dev/blog/using-go-modules

https://ganeshramr.github.io/normal/2015/03/29/goerror1.html

https://golang.org/doc/gopath_code#GOPATH