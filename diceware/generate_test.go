package diceware

import (
	"log"
	"testing"
)

const (
	N = 10000
)

func testUnique(t testing.TB, list []string) {
	seen := make(map[string]struct{}, len(list))
	for _, v := range list {
		if _, ok := seen[v]; ok {
			t.Errorf("found duplicate: %q", list)
		}
		seen[v] = struct{}{}
	}
}

func TestGenerator_Generate(t *testing.T) {
	t.Parallel()

	gen, err := NewGenerator(nil)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < N; i++ {
		list, err := gen.Generate(16)
		if err != nil {
			t.Fatal(err)
		}
		testUnique(t, list)
	}
}

func TestGenerateWordList(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		wordList WordList
	}{
		{
			"eff_large",
			WordListEffLarge(),
		},
		{
			"eff_small",
			WordListEffSmall(),
		},
		{
			"original",
			WordListOriginal(),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for i := 0; i < N; i++ {
				list, err := GenerateWithWordList(16, tc.wordList)
				if err != nil {
					t.Fatal(err)
				}
				testUnique(t, list)
			}
		})
	}
}

func TestRollDie(t *testing.T) {
	t.Parallel()

	for i := 0; i < N; i++ {
		r, err := RollDie()
		if err != nil {
			t.Fatal(err)
		}

		if r < 1 || r > 6 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}

func TestRollWord(t *testing.T) {
	t.Parallel()

	for i := 0; i < N; i++ {
		r, err := RollWord(5)
		if err != nil {
			t.Fatal(err)
		}

		if r < 11111 || r > 66666 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}

func ExampleGenerate() {
	words, err := Generate(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", words)
}

func ExampleMustGenerate() {
	words := MustGenerate(6)
	log.Printf("%q", words)
}

func ExampleGenerator_Generate() {
	gen, err := NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	words, err := gen.Generate(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", words)
}

func ExampleNewGenerator_nil() {
	// This is the same as calling Generate directly at the package level, but you
	// can share the generator across functions.
	gen, err := NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = gen // gen.Generate(...)
}

func ExampleNewGenerator_custom() {
	gen, err := NewGenerator(&GeneratorInput{
		WordList: WordListOriginal(),
	})
	if err != nil {
		log.Fatal(err)
	}

	_ = gen // gen.Generate(...)
}
