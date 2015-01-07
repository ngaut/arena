package arena

type ArenaAllocator interface {
	AllocBytes(n int) []byte
	Reset()
}

type SimpleArenaAllocator struct {
	arena []byte
	off   int
}

func NewArenaAllocator(capability int) *SimpleArenaAllocator {
	return &SimpleArenaAllocator{arena: make([]byte, 0, capability)}
}

func (s *SimpleArenaAllocator) AllocBytes(n int) []byte {
	if s.off+n < cap(s.arena) {
		slice := s.arena[s.off:s.off : s.off+n]
		s.off += n
		return slice
	}

	return make([]byte, 0, n)
}

func (s *SimpleArenaAllocator) Reset() {
	s.off = 0
}
