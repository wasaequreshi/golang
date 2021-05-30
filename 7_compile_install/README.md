# Learning Go Part 7

## Compiling and Installing

In order to run Go programs, you need to setup your environment.

```sh
export PATH=$PATH:/path/to/your/install/directory
```

This is the go bin directory that is setup in your home directory. It should look like the following:

```sh
export PATH=$PATH:/Users/<your username>/go/bin/
```

Once you do this, you need to run the following command on the `go` file which contains the `main` package:

```sh
go install
```

This will compile and install the program in the path you exported earlier. You should now be able to run the executable program from anywhere:

```sh
➜  hello git:(master) ✗ hello
map[Bilal:Hello Bilal. Nice to meet you Wasae:Hi, Wasae. Welcome!]
```