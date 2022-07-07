package ziphelper

import "testing"

func Test_passwordZip(t *testing.T) {
	outPath, err := CreatePasswordProtectedZip("Sample", FileEntry{
		SrcFile:  "testdata/foo.txt",
		Filename: "foo.txt",
	}, FileEntry{
		SrcFile:  "testdata/bar.txt",
		Filename: "bar.txt",
	})
	if err != nil {
		t.Fatalf("Create zip error: %s", err)
	}
	if len(outPath) == 0 {
		t.Fatalf("File path of created zip file is empty.")
	}
}
