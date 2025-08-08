package strategy

type StrategyType int

const (
	Flat StrategyType = iota
	Percent
	PercentWithCap
)

// it is factory which gives the discount staretgy

func GetStrategy(t StrategyType, percent, maxAmount float64) DiscountStrategy {
	switch t {
	case Flat:
		return &FlatDiscount{
			Amount: maxAmount,
		}
	case Percent:
		return &PercentageDiscount{
			Percent: percent,
		}
	case PercentWithCap:
		return &PercentageWithCap{
			Percent: percent,
			Cap:     maxAmount,
		}
	default:
		return nil
	}
}
