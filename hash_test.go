package datastruct

import "testing"

func TestHashTable(t *testing.T) {
	for i := uint(0); i < 3; i++ {
		testHashTable(t, i)
	}
}

func testHashTable(t *testing.T, cap uint) {
	ht := NewHashTable[string, int](cap)
	if ht.Has("a") {
		t.Fatalf("expected=empty; got=notEmpty")
	}
	tt := []struct {
		key    string
		val    int
		updVal int
		del    bool
		reAdd  bool
	}{
		{"abc", 1, 2, true, false},
		{"abc", 2, 2, true, false},
		{"acb", 3, 3, false, false},
		{"bac", 4, 4, false, false},
		{"xa", 5, 5, true, true},
		{"ax", 6, 6, true, false},
	}
	for _, tc := range tt {
		ht.Set(tc.key, tc.val)
	}
	for _, tc := range tt {
		if !ht.Has(tc.key) {
			t.Fatalf("expected=%d; got=empty", tc.updVal)
		}
		val, exists := ht.Get(tc.key)
		if !exists {
			t.Fatalf("expected=%d; got=empty", tc.updVal)
		}
		if val != tc.updVal {
			t.Fatalf("expected=%d; got=%d", tc.updVal, val)
		}
	}
	for _, tc := range tt {
		if tc.del {
			ht.Del(tc.key)
		}
	}
	for _, tc := range tt {
		if tc.reAdd {
			ht.Set(tc.key, tc.updVal)
		}
	}
	for _, tc := range tt {
		shouldBe := !tc.del || tc.reAdd
		exists := ht.Has(tc.key)
		if shouldBe != exists {
			t.Fatalf("expected=%v; got=%v", shouldBe, exists)
		}
		if !exists {
			continue
		}
		val, exists := ht.Get(tc.key)
		if !exists {
			t.Fatalf("expected=%d; got=empty", tc.updVal)
		}
		if val != tc.updVal {
			t.Fatalf("expected=%d; got=%d", tc.updVal, val)
		}
	}
}
