package map2pairs_test

import (
	"testing"

	mp "github.com/takanoriyanagitani/go-asn1-map2pairs"
)

func TestMapToPairs(t *testing.T) {
	t.Parallel()

	t.Run("RawMap", func(t *testing.T) {
		t.Parallel()

		t.Run("ToDerBytes", func(t *testing.T) {
			t.Parallel()

			t.Run("empty", func(t *testing.T) {
				t.Parallel()

				var empty mp.RawMap = map[string]string{}

				der, e := empty.ToDerBytes()
				if nil != e {
					t.Fatalf("unexpected error: %v", e)
				}

				if 0 == len(der) {
					t.Fatal("empty der bytes got")
				}
			})

			t.Run("single", func(t *testing.T) {
				t.Parallel()

				var single mp.RawMap = map[string]string{
					"Helo": "Wrld",
				}

				der, e := single.ToDerBytes()
				if nil != e {
					t.Fatalf("unexpected error: %v", e)
				}

				if 0 == len(der) {
					t.Fatal("empty der bytes got")
				}
			})
		})
	})
}
