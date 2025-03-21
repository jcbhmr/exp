# exp

[![PkgGoDev](https://pkg.go.dev/badge/github.com/jcbhmr/exp)](https://pkg.go.dev/github.com/jcbhmr/exp)

This subrepository holds experimental and deprecated (in the `old`
directory) packages.

The idea for this subrepository originated as the `pkg/exp` directory
of the main repository, but its presence there made it unavailable
to users of the binary downloads of the Go installation. The
subrepository has therefore been created to make it possible to `go
get` these packages.

**Warning:** Packages here are experimental and unreliable. Some may
one day be promoted to the main repository or other subrepository,
or they may be modified arbitrarily or even disappear altogether.

In short, code in this subrepository is not subject to the Go 1
compatibility promise. (No subrepo is, but the promise is even more
likely to be violated by go.exp than the others.)

Caveat emptor.
