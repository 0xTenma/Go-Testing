# Go-Testing
Contains basic smoke test, unit tests, fuzzing in Golang


### Important Points For running tests:

1. Unit testing
`$ mkdir unit-test`
`$ cd unit-test`

For initialize the Go module:
`$ go mod init unit-test`

We need to import testing library
`import ("testing")`

`testing.T` provides the struct with helpful error logging

```go
	t.Log("Similar to fmt.Println()")
	t.Fail() // Will show a test cas has failed in the results
	t.FailNow() // t.Fail() + safely exit without continuing
	t.Error() // t.Log() + t.Fail();
	t.Fatal() // t.Log() + t.FailNow();
```


For testing:
`$ go test`
Testing with verbose:
`$ go test -v`

To run specific test: 
`$ go test -v -run=Test_SayGoodBye`

To check test coverage: 
`$ go test -cover`

To export cover profile to a file:
`$ go test -coverprofile=cover_out`

### Benchmarking with Go