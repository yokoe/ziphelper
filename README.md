# ziphelper

## How to use
### Create password protected zip file

```
outPath, err := ziphelper.CreatePasswordProtectedZip("Sample", FileEntry{
  SrcFile:  "testdata/foo.txt",
  Filename: "foo.txt",
}, FileEntry{
  SrcFile:  "testdata/bar.txt",
  Filename: "bar.txt",
})
```
