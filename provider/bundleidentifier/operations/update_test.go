package operations

import (
	"sort"
	"testing"
)

func TestDiffCapabilities(t *testing.T) {
	tests := []struct {
		name      string
		old, next []string
		wantAdd   []string
		wantRm    []string
	}{
		{"add one", []string{"APPLE_ID_AUTH"}, []string{"APPLE_ID_AUTH", "PUSH_NOTIFICATIONS"}, []string{"PUSH_NOTIFICATIONS"}, nil},
		{"remove one", []string{"APPLE_ID_AUTH", "PUSH_NOTIFICATIONS"}, []string{"APPLE_ID_AUTH"}, nil, []string{"PUSH_NOTIFICATIONS"}},
		{"add and remove", []string{"APPLE_ID_AUTH"}, []string{"PUSH_NOTIFICATIONS"}, []string{"PUSH_NOTIFICATIONS"}, []string{"APPLE_ID_AUTH"}},
		{"no change", []string{"APPLE_ID_AUTH"}, []string{"APPLE_ID_AUTH"}, nil, nil},
		{"from empty", nil, []string{"PUSH_NOTIFICATIONS"}, []string{"PUSH_NOTIFICATIONS"}, nil},
		{"to empty", []string{"PUSH_NOTIFICATIONS"}, nil, nil, []string{"PUSH_NOTIFICATIONS"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, remove := diffCapabilities(tt.old, tt.next)
			assertSameSet(t, tt.wantAdd, add, "add")
			assertSameSet(t, tt.wantRm, remove, "remove")
		})
	}
}

func TestToStringList(t *testing.T) {
	got := toStringList([]any{"APPLE_ID_AUTH", "PUSH_NOTIFICATIONS"})
	assertSameSet(t, []string{"APPLE_ID_AUTH", "PUSH_NOTIFICATIONS"}, got, "toStringList")
	if len(toStringList([]any{})) != 0 {
		t.Fatal("expected empty slice for empty input")
	}
}

func assertSameSet(t *testing.T, want, got []string, label string) {
	t.Helper()
	want = append([]string(nil), want...)
	got = append([]string(nil), got...)
	sort.Strings(want)
	sort.Strings(got)
	if len(want) != len(got) {
		t.Fatalf("%s: want %v, got %v", label, want, got)
	}
	for i := range want {
		if want[i] != got[i] {
			t.Fatalf("%s: want %v, got %v", label, want, got)
		}
	}
}
