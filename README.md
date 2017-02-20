# s

Usage:
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
