package ziphelper

import "testing"

func Test_zip(t *testing.T) {
	outPath, err := CreateZip(FileEntry{
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

func Test_protectedZip(t *testing.T) {
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
