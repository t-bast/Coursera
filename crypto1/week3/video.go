package week3

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const (
	// BlockSize is the number of bytes in a block.
	// The last block might be smaller.
	BlockSize = 1024
)

// Block is a file block with a recursive hash
// to provide integrity.
type Block struct {
	blockBytes   []byte
	blockRecHash []byte
}

// Hash computes the recursive hash of this block.
func (b *Block) Hash() []byte {
	toHash := append(b.blockBytes, b.blockRecHash...)
	h := sha256.Sum256(toHash)
	return h[:]
}

// File represents an authenticated file.
// It is split into authenticated blocks.
type File struct {
	Blocks []*Block
}

// NewFile creates a File from its path.
func NewFile(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	info, err := f.Stat()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	blockCount, lastBlockLength := fileStat(info)
	blocks := make([]*Block, blockCount)

	// The last block won't have a recursive hash and can have a smaller size,
	// so we treat it differently.
	lastBlock, err := readBlockBytes(f, blockCount-1, lastBlockLength)
	if err != nil {
		return nil, err
	}
	blocks[blockCount-1] = lastBlock

	for i := blockCount - 2; i >= 0; i-- {
		block, err := readBlockBytes(f, i, BlockSize)
		if err != nil {
			return nil, err
		}

		block.blockRecHash = blocks[i+1].Hash()
		blocks[i] = block
	}

	return &File{Blocks: blocks}, nil
}

// Hash returns the recursive hash of the whole file.
func (f *File) Hash() []byte {
	return f.Blocks[0].Hash()
}

func fileStat(info os.FileInfo) (int64, int64) {
	fileLen := info.Size()
	blockCount := fileLen / BlockSize
	lastBlockLength := int64(BlockSize)
	if fileLen%BlockSize != 0 {
		blockCount++
		lastBlockLength = fileLen % BlockSize
	}

	return blockCount, lastBlockLength
}

func readBlockBytes(f *os.File, blockIndex, blockSize int64) (*Block, error) {
	blockBytes := make([]byte, blockSize)
	n, err := f.ReadAt(blockBytes, BlockSize*blockIndex)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if int64(n) != blockSize {
		return nil, fmt.Errorf("read %d bytes, expected %d bytes", n, blockSize)
	}

	return &Block{blockBytes: blockBytes}, nil
}
