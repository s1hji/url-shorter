package memory

import (
	"fmt"
	"sync"
	"testing"
)

func TestSave(t *testing.T) {
	memory_storage := NewMemoryStorage()
	err := memory_storage.Save("https://web.telegram.org", "tg")
	if err != nil {
		t.Error(err)
	}
}

func Test_duplicate(t *testing.T) {
	memory_storage := NewMemoryStorage()
	memory_storage.Save("https://web.telegram.org", "tg")

	err := memory_storage.Save("https://web.telegram.org", "tg")
	if err == nil {
		t.Error(err)
	}
}

func TestGetOriginLink(t *testing.T) {
	memory_storage := NewMemoryStorage()
	memory_storage.Save("https://web.telegram.org", "tg")

	origin, err := memory_storage.GetOriginLink("tg")
	if err != nil {
		t.Error(err)
	}

	if origin != "https://web.telegram.org" {
		t.Errorf("ожидали https://web.telegram.org, а пришло %s", origin)
	}

}

func TestGetShortLink(t *testing.T) {
	memory_storage := NewMemoryStorage()
	memory_storage.Save("https://web.telegram.org", "tg")

	short, err := memory_storage.GetShortLink("https://web.telegram.org")
	if err != nil {
		t.Error(err)
	}

	if short != "tg" {
		t.Errorf("ожидали увидеть tg, а пришло %s", short)
	}
}

func TestConcurrency(t *testing.T) {
	memory_storage := NewMemoryStorage()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			memory_storage.Save(fmt.Sprintf("telegram.org %d", n), fmt.Sprintf("tg %d", n))
		}(i)
	}

	wg.Wait()

}
