# Testing

Quick intro to testing.

Go doesn't haven't an "assert" but actually just uses the if.

`t.Errorf` assumes a failure and outputs to console, so basically use an if to assert, and if that assertion fails, error out with the message.