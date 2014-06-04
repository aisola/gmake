GMake (1.0)
===========

A very lightweight build utility written in Go.

### Install

    $ git clone https://github.com/aisola/gmake.git
    $ go build

### Getting Started

In any project, create a file `GMakefile` and use the following syntax to create rules:

    target {
        command;
    }

Notice that the commands end in semicolons. This is required. A more realistic example 
is shown below compiling a go program:

    all {
        go build -o hello main.go;
    }
    
    fmt {
        go fmt main.go;
    }
    
    clean {
        rm hello;
    }

Enjoy!
ACI    
