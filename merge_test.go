package merge

import (
	"reflect"
	"testing"
)

func TestMergeReturnsAscendingValuesFromAllCollections(t *testing.T) {
	tests := []struct {
		name        string
		collection1 []int
		collection2 []int
		collection3 []int
		want        []int
	}{
		{
			name:        "mixed inputs from all collections",
			collection1: []int{9, 7, 5, 1},
			collection2: []int{2, 4, 6, 8},
			collection3: []int{0, 3, 10},
			want:        []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:        "duplicates and negative values",
			collection1: []int{5, 3, 3, -2},
			collection2: []int{-4, 3, 6},
			collection3: []int{-4, 0, 3, 7},
			want:        []int{-4, -4, -2, 0, 3, 3, 3, 3, 5, 6, 7},
		},
		{
			name:        "empty collections",
			collection1: []int{},
			collection2: []int{},
			collection3: []int{},
			want:        []int{},
		},
		{
			name:        "one collection already exhausted",
			collection1: []int{},
			collection2: []int{1, 2, 3},
			collection3: []int{4, 5, 6},
			want:        []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:        "two collections already exhausted",
			collection1: []int{3, 2, 1},
			collection2: []int{},
			collection3: []int{},
			want:        []int{1, 2, 3},
		},
		{
			name:        "nil slices behave like empty slices",
			collection1: nil,
			collection2: []int{1, 4},
			collection3: nil,
			want:        []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.collection1, tt.collection2, tt.collection3)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Merge(%v, %v, %v) = %v, want %v", tt.collection1, tt.collection2, tt.collection3, got, tt.want)
			}
		})
	}
}

func TestMergeReturnsNewSlice(t *testing.T) {
	collection1 := []int{3, 2, 1}
	collection2 := []int{4, 5}
	collection3 := []int{6, 7}

	got := Merge(collection1, collection2, collection3)
	if len(got) == 0 {
		t.Fatal("Merge returned an empty result")
	}

	got[0] = 99

	if !reflect.DeepEqual(collection1, []int{3, 2, 1}) {
		t.Fatalf("collection1 was mutated or aliased: %v", collection1)
	}
	if !reflect.DeepEqual(collection2, []int{4, 5}) {
		t.Fatalf("collection2 was mutated or aliased: %v", collection2)
	}
	if !reflect.DeepEqual(collection3, []int{6, 7}) {
		t.Fatalf("collection3 was mutated or aliased: %v", collection3)
	}
}
