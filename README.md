# GitMoji CLI in Go

Inspired by: https://github.com/carloscuesta/gitmoji-cli

A simple CLI util written in Go to eliminate searching manually through
the list of codepoints, shortcodes, and rationale messages on
https://gitmoji.dev/

![Example](https://raw.githubusercontent.com/vald-phoenix/gmj/master/media/example.svg)

In order to install util invoke:

```sh
$ go install github.com/vald-phoenix/gmj@v0.1.0
```

Or download binaries from release page corresponding to your OS.

You may do something like this when committing as well:

```sh
$ git commit -m "$(gmj -c memo) Update README"
$ git commit -m "$(gmj -c memo) $(gmj -u memo)"
```
