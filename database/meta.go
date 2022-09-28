package database

import (
	"encoding/binary"
)

const (
	metaPageNumber = 0
)

type Meta struct {
	FreePage uint64
}

func NewMeta() *Meta {
	return &Meta{}
}

func (m *Meta) Serialize(buffer []byte) {
	pos := 0

	binary.LittleEndian.PutUint64(buffer[:], uint64(m.FreePage))
	pos += pageNumberSize
}

func (m *Meta) Deserialize(buffer []byte) {
	pos := 0
	m.FreePage = uint64(binary.LittleEndian.Uint64(buffer[pos:]))
	pos += pageNumberSize
}
