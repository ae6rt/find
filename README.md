## Find files

The find package works like the Unix find command, but with fewer options.  It supports
finding regular files by regex in a given root directory.  Results can be tuned to only return
results at a certain depth, and a set of given subdirectories directories can be ignored altogether.

```
files, err := Find("testdata", regexp.MustCompile(`cert\.pem`), 2, []string{".dot"})
```

will find in directory _testdata_ files named _cert.pem_ at depth 2.  Any subdirectory of testdata
named _.dot_ will be ignored, which is useful for ignoring .git directories when traversing git clones.