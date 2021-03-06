package arena

type ArenaAllocator interface {
	AllocBytes(n int) []byte
	AllocBytesWithLen(length int, capability int) []byte
	Reset()
}

type SimpleArenaAllocator struct {
	arena []byte
	off   int
}

type stdAllocator struct {
}

func (a *stdAllocator) AllocBytes(n int) []byte {
	return make([]byte, 0, n)
}

func (a *stdAllocator) AllocBytesWithLen(length int, capability int) []byte {
	return make([]byte, length, capability)
}

func (a *stdAllocator) Reset() {
}

var _ ArenaAllocator = &stdAllocator{}

var StdAllocator = &stdAllocator{}

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

func (s *SimpleArenaAllocator) AllocBytesWithLen(length int, capability int) []byte {
	slice := s.AllocBytes(capability)
	return slice[:length:capability]
}

func (s *SimpleArenaAllocator) Reset() {
	s.off = 0
}
