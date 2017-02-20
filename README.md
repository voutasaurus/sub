# s

A recursive find and replace tool.

# Install
Install go: https://golang.org/dl/

Then:
```
$ go get -u github.com/voutasaurus/s
```

# Usage:
```
$ s old new
```

This will (starting in the current directory) recursively replace the 
string `old` by the string `new` in all files at or below this directory.

Files and directories beginning with `.` and `_` are ignored. 
Files and directories named `vendor` are also ignored.

```
$ s -p old new
```

This will do the same except it prompts for each file before doing the replacement.

# why did you do this?

Because using `sed` to recursively replace one word with another word is genuinely harder to figure out how to do than writing a go program to do it.
