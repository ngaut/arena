package arena

import (
	"testing"
)

func TestSimpleArenaAllocator(t *testing.T) {
	arena := NewArenaAllocator(1000)
	slice := arena.AllocBytes(10)
	if arena.off != 10 {
		t.Error("off not match, expect 10 bug got", arena.off)
	}

	if len(slice) != 0 {
		t.Error("slice len not match, expect 0 bug got", len(slice))
	}

	if cap(slice) != 10 {
		t.Error("slice len not match, expect 0 bug got", cap(slice))
	}

	slice = arena.AllocBytes(20)
	if arena.off != 30 {
		t.Error("off not match, expect 30 bug got", arena.off)
	}

	if len(slice) != 0 {
		t.Error("slice len not match, expect 0 bug got", len(slice))
	}

	if cap(slice) != 20 {
		t.Error("slice len not match, expect 0 bug got", cap(slice))
	}

	slice = arena.AllocBytes(1024)

	if arena.off != 30 {
		t.Error("off not match, expect 30 bug got", arena.off)
	}

	if len(slice) != 0 {
		t.Error("slice len not match, expect 0 bug got", len(slice))
	}

	if cap(slice) != 1024 {
		t.Error("slice len not match, expect 0 bug got", cap(slice))
	}

	arena.Reset()
	if arena.off != 0 || cap(arena.arena) != 1000 {
		t.Error("off should be 0")
	}

	if cap(arena.arena) != 1000 {
		t.Error("off should be 0")
	}
}
