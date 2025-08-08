package strategy

type DiscountStrategy interface {
	Calculate(baseAmount float64) float64
}

type FlatDiscount struct {
	Amount float64
}

func (f *FlatDiscount) Calculate(baseAmount float64) float64 {
	if f.Amount > baseAmount {
		return baseAmount
	}
	return f.Amount
}

type PercentageDiscount struct {
	Percent float64
}

func (p *PercentageDiscount) Calculate(baseAmount float64) float64 {
	return (p.Percent / 100.00) * baseAmount
}

type PercentageWithCap struct {
	Percent float64
	Cap     float64
}

func (pc *PercentageWithCap) Calculate(baseAmount float64) float64 {
	disc := (pc.Percent / 100.0) * baseAmount
	if disc > pc.Cap {
		disc = pc.Cap
	}
	return disc
}
