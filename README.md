mockery
=======

mockery (v2) is used to generate mock implementations of Go interfaces.


## Installation

Run `go get github.com/shoenig/mockery`

## Arguments

##### version

Use `-version=true` to print the mockery version string and exit.

##### interface

Use `-interface=Interface` to specify which interface a mock should be generated for.

##### package

Use `-package=name` to specify the name of the output package name containing the generated mocks.

##### stdout

Use `-stdout=true` to have mockery print generated mocks to standard out instead of writing to disk.

##### comment

Use `-comment="some text"` to have mockery inject a comment into the prologue of each generated file.

## Environment

##### custom import prefix

Set `MOCKERY_IMPORT_PREFIX=internal.net/packages/` to hack mockery prefix generated imports with custom location.

##### verification of no changes

Set `MOCKERY_CHECK_NOCHANGE=1` to have mockery return non-zero exit code if any files are modified.

## Best Practices

Typically, a project should always use `mockery` via the `go generate` command so that future developers
can see exactly how generated code was created. Each interface to be mocked should have its own generate
line. Libraries which export interfaces should always provide mocks to those interfaces, in a subpackage.
For example if you have package "foo", it should have generate lines that put the generated mocks into a
sub-package called "footest". That way, clients can easily consume those mocks.

## Examples

```go
package foo

//go:generate mockery -interface=Bar -package=footest

type Bar interface {
	String() string
}

//go:generate mockery -interface=Bazzer -package=footest

type Bazzer interface {
	Baz() error
}
```

## Private Repositories

Some teams use private repositories for all 3rd-party code, including the `github.com/stretchr/testify`
packages, references to which are used in generated mocks. To support internally hosted copies of these
packages, set `MOCKERY_IMPORT_PREFIX=some/internal/prefix` to automatically mangle the import string to
reference the internally mirrored packages.

## Continuous Integration

In an effort to prevent the checking-in of artisinal generated content that is difficult to reproduce
without explicit instructions, `mockery` supports a verification mode that returns a non-zero exit code
if any files were modified during execution. The idea is that this can be used in CI environments, where
running `mockery` should not make changes to generated mocks already checked into the repository. To enable
this mode, set `MOCKERY_CHECK_NOCHANGE=1`.
