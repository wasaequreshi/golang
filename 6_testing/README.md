# Learning Go Part 6

## Testing

In this part, we created `greetings_test.go` in the `greetings` directory. A test function name requires `Test` at the prefix in order for Go to pick up the test:

```sh 
func TestHelloWorld(...)
```

In addition the function parameter takes a pointer to `testing.T` type:

```sh
func TestHelloWorld(t *testing.T) {
    ...
}
```

To fail a function, you need to use `Fatalf` of `testing.T`:

```sh
func TestHelloWorld(t *testing.T) {
    if <test condition fails> {
        t.Fatalf(...)
    }
}
```

To execute the tests, run the following command:

```sh
go test
```

You can pass the -v to get verbose output.