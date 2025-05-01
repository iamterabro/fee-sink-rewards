package main

import (
	feesinkreward "github.com/iamterabro/fee-sink-rewards"
	"github.com/rs/zerolog/log"
)

const round = uint64(49530251)

func main() {
	log.Info().Msgf("Proposer payout for round: %d", round)

	blockHashRandSource := feesinkreward.CreateRandomSourceFromAlgorandBlockHash(round)

	choiceType := feesinkreward.DetermineChoiceType(blockHashRandSource)

	log.Info().Msgf("Choice type: %s", choiceType.ChoiceType)

	if choiceType.ChoiceType == feesinkreward.ChoiceSchizo {
		randomSchizo := feesinkreward.PickRandomSchizo(blockHashRandSource)

		log.Info().Msgf("Random SCHIZO: #%d", randomSchizo)
	}
}
