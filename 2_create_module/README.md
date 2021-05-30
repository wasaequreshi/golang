# Learning Go Part 2

## Creating A Module

`greetings` is a custom module created that can be used by others. A couple of new commands. One was from the previous section but I missed it. It was not required then but in this section it was needed.

### Dependency Tracker

```sh
    go mod init example.com/<module name>
```

This is needed to track the dependencies for the go module. This was ran in the `greetings` and `hello` directory. See `go.mod` file in each respective directory.

### Adding Location Of Dependency (If Not Public)

```sh
    go mod edit -replace=example.com/<module name>=../<path to module name>
```

This was ran in the `hello` directory to let the go compiler know that the `greetings` module was local. You can see in the `go.mod` file in the `hello` directory that it includes a `replace` to let the compiler know what replace during runtime to locate the `greetings` module appropriately.

### Adding All Dependencies For Go Module

```sh
    go mod tidy
```

This was ran in the `hello` directory to include the dependencies used in the `hello.go` file. This is different from the previous step. The previous step is just letting the compiler know the location. That doesn't mean it's being used or not in the `hello.go` file. One may be frustrated at this point but having gone through `c` programming, I can bare with this :) 

### Alternative Way Of Executing

```sh
    go run .
```

Similar to 

```sh
    go run hello.go
```