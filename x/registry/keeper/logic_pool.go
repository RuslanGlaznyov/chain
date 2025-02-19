package keeper

import (
	"math"
	"math/rand"
	"sort"

	"github.com/KYVENetwork/chain/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func containsElement(array []string, element string) bool {
	for _, v := range array {
		if v == element {
			return true
		}
	}
	return false
}

// updateLowestFunder is an internal function that updates the lowest funder entry in a given pool.
func (k Keeper) handleNonVoters(ctx sdk.Context, pool *types.Pool) {
	nonVoters := make([]string, 0)

	for _, staker := range pool.Stakers {
		if staker == pool.BundleProposal.Uploader {
			continue
		}

		valid := containsElement(pool.BundleProposal.VotersValid, staker)
		invalid := containsElement(pool.BundleProposal.VotersInvalid, staker)
		abstain := containsElement(pool.BundleProposal.VotersAbstain, staker)

		if !valid && !invalid && !abstain {
			nonVoters = append(nonVoters, staker)
		}
	}

	for _, voter := range nonVoters {
		staker, foundStaker := k.GetStaker(ctx, voter, pool.Id)

		// skip timeout slash if staker is not found
		if foundStaker {

			if staker.Points < k.MaxPoints(ctx) {
				// Increase points
				staker.Points += 1
				k.SetStaker(ctx, staker)
			} else {
				// slash nonVoter for not voting in time
				slashAmount := k.slashStaker(ctx, pool, staker.Account, k.TimeoutSlash(ctx))

				// emit slashing event
				ctx.EventManager().EmitTypedEvent(&types.EventSlash{
					PoolId:    pool.Id,
					Address:   staker.Account,
					Amount:    slashAmount,
					SlashType: types.SLASH_TYPE_TIMEOUT,
				})

				// Check if staker is still in stakers list and remove staker.
				staker, foundStaker = k.GetStaker(ctx, voter, pool.Id)

				// check if next uploader is still there or already removed
				if foundStaker {
					// Transfer remaining stake to account.
					k.TransferToAddress(ctx, staker.Account, staker.Amount)

					// remove current next_uploader
					k.removeStaker(ctx, pool, &staker)

					// emit unstake event
					ctx.EventManager().EmitTypedEvent(&types.EventUnstakePool{
						PoolId:  pool.Id,
						Address: staker.Account,
						Amount:  staker.Amount,
					})
				}

				// Update current lowest staker
				k.updateLowestStaker(ctx, pool)
			}
		}
	}
}

// updateLowestFunder is an internal function that updates the lowest funder entry in a given pool.
func (k Keeper) updateLowestFunder(ctx sdk.Context, pool *types.Pool) {
	minAmount := uint64(math.Inf(0))
	minFunder := ""

	for _, account := range pool.Funders {
		funder, _ := k.GetFunder(ctx, account, pool.Id)

		if funder.Amount <= minAmount {
			minAmount = funder.Amount
			minFunder = funder.Account
		}
	}

	pool.LowestFunder = minFunder
}

// updateLowestStaker is an internal function that updates the lowest staker entry in a given pool.
func (k Keeper) updateLowestStaker(ctx sdk.Context, pool *types.Pool) {
	minAmount := uint64(math.Inf(0))
	minStaker := ""

	for _, account := range pool.Stakers {
		staker, _ := k.GetStaker(ctx, account, pool.Id)

		if staker.Amount <= minAmount {
			minAmount = staker.Amount
			minStaker = staker.Account
		}
	}

	pool.LowestStaker = minStaker
}

// removeFunder is an internal function that removes a funder from a given pool.
func (k Keeper) removeFunder(ctx sdk.Context, pool *types.Pool, funder *types.Funder) {
	// Find the index of the given funder.
	var funderIndex = -1

	for i, v := range pool.Funders {
		if v == funder.Account {
			funderIndex = i
			break
		}
	}

	// Return if the funder wasn't found.
	if funderIndex < 0 {
		return
	}

	// Remove funder from list of funders (replace with last entry and then slice).
	pool.Funders[funderIndex] = pool.Funders[len(pool.Funders)-1]
	pool.Funders = pool.Funders[:len(pool.Funders)-1]

	k.RemoveFunder(ctx, funder.Account, funder.PoolId)

	// Decrease the pool's total funds.
	pool.TotalFunds -= funder.Amount
}

// removeStaker is an internal function that removes a staker from a given pool.
func (k Keeper) removeStaker(ctx sdk.Context, pool *types.Pool, staker *types.Staker) {
	// Find the index of the given staker.
	var stakerIndex = -1

	for i, v := range pool.Stakers {
		if v == staker.Account {
			stakerIndex = i
			break
		}
	}

	// Return is the staker wasn't found.
	if stakerIndex < 0 {
		return
	}

	// Remove staker from list of stakers (replace with last entry and then slice).
	pool.Stakers[stakerIndex] = pool.Stakers[len(pool.Stakers)-1]
	pool.Stakers = pool.Stakers[:len(pool.Stakers)-1]

	k.RemoveStaker(ctx, staker.Account, staker.PoolId)

	// Decrease the pool's total stake.
	pool.TotalStake -= staker.Amount
}

// RandomChoiceCandidate ...
type RandomChoiceCandidate struct {
	Account string
	Amount  uint64
}

// getWeightedRandomChoice is an internal function that returns a random selection out of a list of candidates.
func (k Keeper) getWeightedRandomChoice(candidates []RandomChoiceCandidate, seed uint64) string {
	type WeightedRandomChoice struct {
		Elements    []string
		Weights     []uint64
		TotalWeight uint64
	}

	wrc := WeightedRandomChoice{}

	for _, candidate := range candidates {
		i := sort.Search(len(wrc.Weights), func(i int) bool { return wrc.Weights[i] > candidate.Amount })
		wrc.Weights = append(wrc.Weights, 0)
		wrc.Elements = append(wrc.Elements, "")
		copy(wrc.Weights[i+1:], wrc.Weights[i:])
		copy(wrc.Elements[i+1:], wrc.Elements[i:])
		wrc.Weights[i] = candidate.Amount
		wrc.Elements[i] = candidate.Account
		wrc.TotalWeight += candidate.Amount
	}

	rand.Seed(int64(seed))
	value := uint64(math.Floor(rand.Float64() * float64(wrc.TotalWeight)))

	for key, weight := range wrc.Weights {
		if weight > value {
			return wrc.Elements[key]
		}

		value -= weight
	}

	return ""
}

func (k Keeper) GetUploadProbability(ctx sdk.Context, stakerAddress string, poolId uint64) sdk.Dec {

	pool, poolFound := k.GetPool(ctx, poolId)
	if !poolFound {
		return sdk.NewDec(0)
	}

	totalWeight := uint64(0)
	userWeight := uint64(0)

	for _, s := range pool.Stakers {
		staker, _ := k.GetStaker(ctx, s, pool.Id)
		delegation, _ := k.GetDelegationPoolData(ctx, pool.Id, s)

		totalWeight += staker.Amount + getDelegationWeight(delegation.TotalDelegation)
		if staker.Account == stakerAddress {
			userWeight = staker.Amount + getDelegationWeight(delegation.TotalDelegation)
		}
	}

	return sdk.NewDec(int64(userWeight)).Quo(sdk.NewDec(int64(totalWeight)))
}

// Calculate Delegation weight to influnce the upload probability
// formula:
// A = 10000, dec = 10**9
// weight = dec * (sqrt(A * (A + x/dec)) - A)
func getDelegationWeight(delegation uint64) uint64 {

	const A uint64 = 10000

	number := A * (A + (delegation / 1_000_000_000))

	// Deterministic sqrt using only int
	// Uses the babylon recursive formula:
	// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method
	var x uint64 = 14142 // expected value for 10000 $KYVE as input
	var xn uint64
	var epsilon uint64 = 100
	for epsilon > 2 {

		xn = (x + number/x) / 2

		if xn > x {
			epsilon = xn - x
		} else {
			epsilon = x - xn
		}
		x = xn
	}

	return (x - A) * 1_000_000_000
}

// getNextUploaderByRandom is an internal function that randomly selects the next uploader for a given pool.
func (k Keeper) getNextUploaderByRandom(ctx sdk.Context, pool *types.Pool, candidates []string) (nextUploader string) {
	var _candidates []RandomChoiceCandidate

	if len(candidates) == 0 {
		return ""
	}

	for _, s := range candidates {
		staker, foundStaker := k.GetStaker(ctx, s, pool.Id)
		delegation, foundDelegation := k.GetDelegationPoolData(ctx, pool.Id, s)

		if foundStaker {
			if foundDelegation {
				_candidates = append(_candidates, RandomChoiceCandidate{
					Account: s,
					Amount:  staker.Amount + getDelegationWeight(delegation.TotalDelegation),
				})
			} else {
				_candidates = append(_candidates, RandomChoiceCandidate{
					Account: s,
					Amount:  staker.Amount,
				})
			}
		}
	}

	return k.getWeightedRandomChoice(_candidates, uint64(ctx.BlockHeight()+ctx.BlockTime().Unix()))
}

// slashStaker is an internal function that slashes a staker in a given pool by a certain percentage.
func (k Keeper) slashStaker(
	ctx sdk.Context, pool *types.Pool, stakerAddress string, slashAmountRatioDecimalString string,
) (slash uint64) {
	staker, found := k.GetStaker(ctx, stakerAddress, pool.Id)

	if found {
		// Parse the provided slash percentage and panic on any errors.
		slashAmountRatio, err := sdk.NewDecFromStr(slashAmountRatioDecimalString)
		if err != nil {
			k.PanicHalt(ctx, "Invalid value for params: "+slashAmountRatioDecimalString+" error: "+err.Error())
		}

		// Compute how much we're going to slash the staker.
		slash = uint64(sdk.NewDec(int64(staker.Amount)).Mul(slashAmountRatio).RoundInt64())

		if staker.Amount == slash {
			// If we are slashing the entire staking amount, remove the staker.
			k.removeStaker(ctx, pool, &staker)
		} else {
			// Subtract slashing amount from staking amount, and update the pool's total stake.
			staker.Amount = staker.Amount - slash
			k.SetStaker(ctx, staker)

			pool.TotalStake -= slash
		}

		// Transfer the slashed amount to the treasury.
		err = k.transferToTreasury(ctx, slash)
		if err != nil {
			k.PanicHalt(ctx, err.Error())
		}
	}

	return slash
}
