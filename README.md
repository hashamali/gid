## gid

[![CircleCI](https://circleci.com/gh/hashamali/gid/tree/master.svg?style=svg)](https://circleci.com/gh/hashamali/gid/tree/master)

Helper that converts UUIDs to base62 short IDs (which are 22 characters long), or strings without dashes. And vice versa.

#### Methods

* `GetShortID`: Converts a UUID into a base62 encoded 22 character string.
* `GetUUIDFromShortID`: Converts a valid base62 encoded 22 character string into a UUID, returns an error if invalid.
* `GetCollapsedID`: Converts a UUID into a string with no dashes.
* `GetUUIDFromShortID`: Converts a valid 32 character string into a UUID, returns an error if invalid.

#### Testing

`make test`
