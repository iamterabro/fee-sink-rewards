package feesinkreward

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/mroth/weightedrand/v2"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type ChoiceType string

const ChoiceSchizo ChoiceType = "schizo_holder"
const ChoiceBurnBRO ChoiceType = "burn_bro"
const ChoiceNodeCost ChoiceType = "node_running_cost"

type Choice struct {
	ChoiceType ChoiceType
	Weight     int
}

var rewardChoices = []Choice{
	{ChoiceType: ChoiceSchizo, Weight: 40},
	{ChoiceType: ChoiceBurnBRO, Weight: 40},
	{ChoiceType: ChoiceNodeCost, Weight: 20},
}

func DetermineChoiceType(randSource *rand.Rand) Choice {
	var choices []weightedrand.Choice[Choice, int]
	for _, trait := range rewardChoices {
		choices = append(choices, weightedrand.Choice[Choice, int]{
			Item:   trait,
			Weight: trait.Weight,
		})
	}
	chooser, err := weightedrand.NewChooser[Choice, int](choices...)
	if err != nil {
		panic(err)
	}
	return chooser.PickSource(randSource)
}

func CreateRandomSourceFromAlgorandBlockHash(round uint64) *rand.Rand {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://mainnet-api.algonode.cloud/v2/blocks/%d/hash", round))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var resBody map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		panic(err)
	}

	blockHash := resBody["blockHash"]

	// Compute SHA256 hash of the block
	hash := sha256.Sum256([]byte(blockHash))

	// Convert the first 8 bytes of the hash to an int64 for seeding
	seed := int64(binary.BigEndian.Uint64(hash[:8]))

	// Seed the random number generator
	return rand.New(rand.NewSource(int64(math.Abs(float64(seed)))))
}

func PickRandomSchizo(randSource *rand.Rand) int {
	return randSource.Intn(999) + 1
}
