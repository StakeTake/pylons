package types

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WeightedParam is to make structs which is using weight to be based on
type OutputsList struct {
	Result []int
	Weight int // TODO should upgrade to string and execute program to calculate weight
}

func (ol OutputsList) String() string {
	return fmt.Sprintf("OutputsList{Result: %v, Weight: %v}", ol.Result, ol.Weight)
}

func (ol OutputsList) GetWeight() int {
	// TODO should handle < 0 weight when upgrading Weight as Program
	if ol.Weight < 0 {
		return 0
	}
	return ol.Weight
}

// WeightedOutputsList is a struct to keep items which can be generated by weight;
// ItemOutput and CoinOutput is possible in current stage
type WeightedOutputsList []OutputsList

func (wpl WeightedOutputsList) String() string {
	itm := "WeightedOutputsList{"

	for _, output := range wpl {
		itm += output.String() + ",\n"
	}

	itm += "}"
	return itm
}

func (wol WeightedOutputsList) Actualize() ([]int, sdk.Error) {
	lastWeight := 0
	var weights []int
	for _, wp := range wol {
		lastWeight += wp.GetWeight()
		weights = append(weights, lastWeight)
	}
	if len(wol) == 0 {
		return nil, nil
	}

	if lastWeight == 0 {
		return nil, sdk.ErrInternal("total weight of weighted param list shouldn't be zero")
	}
	randWeight := rand.Intn(lastWeight)

	first := 0
	chosenIndex := -1
	for i, weight := range weights {
		if randWeight >= first && randWeight <= weight {
			chosenIndex = i
			break
		}
		first = weight
	}
	return wol[chosenIndex].Result, nil
}
