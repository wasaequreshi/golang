# Learning Go Part 6

## Testing

In this part, we created `greetings_test.go`. Test function name require `Test<name>` in order for Go to pick up the test. In addition the function parameter takes a pointer to `testing.T` type. To fail a function, you need to use Fatalf of `testing.T`.

To execute the tests, run the following command:

```sh
    go test
```

You can pass the -v to get verbose output.