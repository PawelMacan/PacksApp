package service

import (
	"errors"
	"math"
	"packs/internal/domain"
	"sort"

	"log/slog"
)

type packCalculator struct {
	packSizes []int
	logger    *slog.Logger
}

func NewPackCalculator(packs []int, logger *slog.Logger) domain.PackCalculatorService {
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	return &packCalculator{
		packSizes: packs,
		logger:    logger,
	}
}

func (p *packCalculator) CalculatePacks(amount int) (*domain.CalculationResult, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	result := p.findOptimalCombination(amount)

	p.logger.Info("calculated optimal packs", "requested", amount, "result", result)

	return result, nil
}

func (p *packCalculator) findOptimalCombination(amount int) *domain.CalculationResult {

	// For smaller numbers, use a recursive approach with pruning
	minOverage := math.MaxInt
	minPacks := math.MaxInt
	var best map[int]int
	var bestTotal int
	var solvedExactly bool

	var try func(idx int, current map[int]int, total int, packs int)
	try = func(idx int, current map[int]int, total int, packs int) {
		// Early termination if we've found an exact match
		if solvedExactly {
			return
		}

		// If we've exceeded or met the target amount
		if total >= amount {
			over := total - amount
			// Update if this is a better solution
			if over < minOverage || (over == minOverage && packs < minPacks) {
				best = make(map[int]int)
				for k, v := range current {
					best[k] = v
				}
				minOverage = over
				minPacks = packs
				bestTotal = total

				// If exact match, we're done
				if over == 0 {
					solvedExactly = true
					return
				}
			}
			return
		}

		// If we've tried all pack sizes
		if idx >= len(p.packSizes) {
			return
		}

		size := p.packSizes[idx]
		// Calculate maximum number of this pack size we could use
		maxCount := (amount - total) / size

		// Try using 0 to maxCount+2 of this pack size
		// The +2 allows for some overage to find optimal solutions
		for i := 0; i <= maxCount+2; i++ {
			if i > 0 {
				current[size] = i
			}
			try(idx+1, current, total+i*size, packs+i)
			if i > 0 {
				delete(current, size)
			}

			// Early termination if we've found an exact match
			if solvedExactly {
				break
			}
		}
	}

	try(0, map[int]int{}, 0, 0)

	return &domain.CalculationResult{
		Packs:           best,
		TotalItems:      bestTotal,
		RequestedAmount: amount,
		Overage:         minOverage,
		TotalPacks:      minPacks,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
