package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadTar(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Open the tar archive for reading.
	tr := tar.NewReader(file)

	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)

		// Get current position of our file reader
		pos, err := file.Seek(0, 1)
		fmt.Printf("Start byte: %d\nsize: %d\nrange: %d", pos, hdr.Size, pos+hdr.Size)

		// if _, err := io.Copy(os.Stdout, tr); err != nil {
		//   log.Fatalln(err)
		// }
		fmt.Println()
	}

}

func main() {
	ReadTar("artifacts.tar")

	// Try dumping an individual file
	file, err := os.Open("artifacts.tar")
	if err != nil {
		panic(err)
	}

	var offset int64 = 2784768
	var size int64 = 318291

	_, err = file.Seek(offset, 0)
	if err != nil {
		panic(err)
	}

	outfile, err := os.Create("out.jpg")
	if err != nil {
		panic(err)
	}

	_, err = io.CopyN(outfile, file, size)
	if err != nil {
		panic(err)
	}
}
