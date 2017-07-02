package week3

import (
	"crypto/sha256"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// FileAuth returns the Authentication Code of given file generated as
/**
Now, a browser downloads the file F one block at a time, where each block includes the appended hash value from the diagram above. When the first block (B0 ∥∥ h1) is received the browser checks that H(B0 ∥∥ h1) is equal to h0 and if so it begins playing the first video block. When the second block (B1 ∥∥ h2) is received the browser checks that H(B1 ∥ h2) is equal to h1 and if so it plays this second block. This process continues until the very last block. This way each block is authenticated and played as it is received and there is no need to wait until the entire file is downloaded.
*/
func FileAuth(path string) [32]byte {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	const ChunkSize int64 = 1024

	info, err := f.Stat()
	check(err)

	fileSize := info.Size()

	// early return for empty file
	if fileSize == 0 {
		return sha256.Sum256([]byte(""))
	}

	var chunkNumber int64 = fileSize / ChunkSize
	lastChunkSize := fileSize % ChunkSize
	if lastChunkSize == 0 {
		lastChunkSize = ChunkSize
		chunkNumber--
	}

	buf := make([]byte, ChunkSize)
	_, err = f.Seek(-lastChunkSize, 2)
	check(err)

	_, err = f.Read(buf)
	if err != io.EOF {
		check(err)
	}
	var bufArr = buf[:lastChunkSize]
	hashBuf := sha256.Sum256(bufArr)

	var offset int64 = -lastChunkSize - ChunkSize
	for ; offset >= -fileSize; offset -= ChunkSize {
		_, err = f.Seek(offset, 2)
		check(err)
		_, err = f.Read(buf)
		check(err)
		var bufArr = buf[:]
		hashBuf = sha256.Sum256(append(bufArr, hashBuf[:]...))
	}

	return hashBuf

}
