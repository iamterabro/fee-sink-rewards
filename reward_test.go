package feesinkreward_test

import (
	feesinkreward "github.com/iamterabro/fee-sink-rewards"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDetermineChoiceType(t *testing.T) {
	testCases := map[string]struct {
		seed           int64
		expectedChoice feesinkreward.ChoiceType
	}{
		"Pick Schizo":  {seed: 55, expectedChoice: feesinkreward.ChoiceSchizo},
		"Pick BurnBRO": {seed: 11, expectedChoice: feesinkreward.ChoiceBurnBRO},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			randSource := fixedRandSource(tc.seed)
			choiceType := feesinkreward.DetermineChoiceType(randSource)
			assert.Equal(t, tc.expectedChoice, choiceType.ChoiceType)
		})
	}
}

func fixedRandSource(seed int64) *rand.Rand {
	source := rand.NewSource(seed)
	return rand.New(source)
}
