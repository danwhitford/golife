package board

import (
	"hash/crc32"
	"math/rand"
	"strings"
	"time"

	"github.com/bits-and-blooms/bitset"
)

type coord struct {
	x, y uint
}

type Board struct {
	Cells         *bitset.BitSet
	Height, Width uint
}

func NewRandom(width, height uint) *Board {
	size := height * width
	bitset := bitset.New(size)
	var i uint
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i = 0; i < size; i++ {
		if r.Intn(3) == 0 {
			bitset.Set(i)
		}
	}
	return &Board{bitset, height, width}
}

func (board Board) String() string {
	var buf strings.Builder
	buf.WriteString("\033[H\033[3J")

	var i uint
	for i = 0; i < board.Cells.Len(); i++ {
		if i > 0 && i%board.Width == 0 {
			buf.WriteRune('\n')
		}
		if board.Cells.Test(i) {
			buf.WriteString("\033[42m")
		} else {
			buf.WriteString("\033[40m")
		}
		buf.WriteRune(' ')
	}

	return buf.String()
}

func (board Board) GetAt(x, y uint) bool {
	if x >= board.Width || y >= board.Height {
		return false
	}
	i := (y * board.Width) + x
	return board.Cells.Test(i)
}

func (board Board) getI(x, y uint) uint {
	return (y * board.Width) + x
}

func (board Board) Hash() uint32 {
	b, _ := board.Cells.MarshalBinary()
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum(b, crc32q)
}

func (board Board) Copy() *Board {
	cpy := bitset.New(board.Cells.Len())
	board.Cells.CopyFull(cpy)
	return &Board{Cells: cpy, Height: board.Height, Width: board.Width}
}

func (board *Board) SetAt(x, y uint) {
	i := board.getI(x, y)
	board.Cells.Set(i)
}

func (board *Board) SetAtTo(x, y uint, to bool) {
	i := board.getI(x, y)
	board.Cells.SetTo(i, to)
}

func (board Board) CountNeighbours(x, y uint) int {
	count := 0
	for _, v := range []coord{
		{x: x - 1, y: y - 1},
		{x: x, y: y - 1},
		{x: x + 1, y: y - 1},
		{x: x - 1, y: y},
		{x: x + 1, y: y},
		{x: x - 1, y: y + 1},
		{x: x, y: y + 1},
		{x: x + 1, y: y + 1},
	} {
		if board.GetAt(v.x, v.y) {
			count++
		}
	}
	return count
}
