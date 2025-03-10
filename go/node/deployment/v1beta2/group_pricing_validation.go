package v1beta2

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func validateGroupPricing(gspec GroupSpec) error {
	var price sdk.DecCoin

	mem := sdk.NewInt(0)
	denom := ""

	for idx, resource := range gspec.Resources {

		if err := validateUnitPricing(resource); err != nil {
			return fmt.Errorf("group %v: %w", gspec.GetName(), err)
		}

		// all must be same denomination
		if denom == "" {
			denom = resource.FullPrice().Denom
		} else {
			if resource.FullPrice().Denom != denom {
				return fmt.Errorf("%w: denomination must be %q", ErrInvalidDeployment, denom)
			}
		}

		if idx == 0 {
			price = resource.FullPrice()
		} else {
			rprice := resource.FullPrice()
			if rprice.Denom != price.Denom {
				return fmt.Errorf("multi-denonimation group: (%v == %v fails)", rprice.Denom, price.Denom)
			}
			price = price.Add(rprice)
		}

		memCount := sdk.NewInt(0)
		if u := resource.Resources.Memory; u != nil {
			memCount.Add(sdk.NewIntFromUint64(u.Quantity.Value()))
		}

		mem = mem.Add(memCount.Mul(sdk.NewIntFromUint64(uint64(resource.Count))))
	}

	return nil
}

func validateUnitPricing(rg Resource) error {
	if !rg.GetPrice().IsValid() {
		return fmt.Errorf("error: invalid price object")
	}

	if rg.Price.Amount.GT(sdk.NewDecFromInt(sdk.NewIntFromUint64(validationConfig.MaxUnitPrice))) {
		return fmt.Errorf("error: invalid unit price (%v > %v fails)", validationConfig.MaxUnitPrice, rg.Price)
	}

	return nil
}

func validateOrderBidDuration(rg GroupSpec) error {
	return nil
}
