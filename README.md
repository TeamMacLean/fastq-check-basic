# FastQ Check (basic)

This tool tests every 4 lines of a file and checks that 1/4 starts with an "@" and that 3/4 starts with an "+".
If it fails to match this pattern it will tell you where it failed and on which line the last good read ended.

## Build
```
go build fq_test.go
chmod +x ./fq_test
```

## Run
```
fq_test path/to/file.fq
```

