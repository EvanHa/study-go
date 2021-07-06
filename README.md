# study-golang

## testing

### Types of Tests

- Unit Tests - These are the lowest level tests that validate that at a base level, a function can take in an input and produce the correct output.
- Integration Tests - These are slightly more involved tests that cover a wider aspect of your code. They cover a particular use-case within your system, but they don’t talk to upstream or downstream systems. These dependencies are often mocked to simulate how our system would react given mocked responses from other systems.
- End-2-End Tests - These are the most involved tests that require the most effort to setup and tear down.

### Running Tests in Go
```bash
$ go test ./...
```
