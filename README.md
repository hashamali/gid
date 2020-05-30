# gid
[![tests](https://img.shields.io/github/workflow/status/hashamali/gid/tests?label=tests&style=flat-square)](https://github.com/hashamali/gid/actions?query=workflow%3Atests)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gid/security?label=security&style=flat-square)](https://github.com/hashamali/gid/actions?query=workflow%3Asecurity)
[![coverage](https://img.shields.io/codecov/c/github/hashamali/gid)](https://codecov.io/gh/hashamali/gid)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gid)](https://goreportcard.com/report/github.com/hashamali/gid)
[![license](https://badgen.net/github/license/hashamali/gid)](https://opensource.org/licenses/MIT)

Helper that converts UUIDs to base62 short IDs (which are 22 characters long), or strings without dashes. And vice versa.

#### Methods

* `GetShortID`: Converts a UUID into a base62 encoded 22 character string.
* `GetUUIDFromShortID`: Converts a valid base62 encoded 22 character string into a UUID, returns an error if invalid.
* `GetCollapsedID`: Converts a UUID into a string with no dashes.
* `GetUUIDFromShortID`: Converts a valid 32 character string into a UUID, returns an error if invalid.

#### Testing

`make test`
