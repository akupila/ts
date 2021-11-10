# ts

A tiny no-fuss utility to get the current timestamp in various formats. On
macOS, `date +%s` provides a UNIX timestamp in seconds, i needed an easy way to
get microseconds.
Another option is to install `coreutils` (`brew install coreutils`) and use
`gdate +%s%N` to get nanoseconds, but this is too much to remember.

```
$ ts
1636628860
$ ts ms
1636628867045
$ ts us
1636628893786386
$ ts "2006-01-02 15:04:05.000"
2021-11-11 12:08:23.710
```
