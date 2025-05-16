package service

import (
	"log/slog"
	"os"
	"packs/internal/domain"
	"reflect"
	"testing"
)

func testLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func TestNewPackCalculator(t *testing.T) {
	logger := testLogger()

	tests := []struct {
		name string
		args []int
	}{
		{
			name: "Valid calculator creation",
			args: []int{250, 500, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := NewPackCalculator(tt.args, logger)
			if calc == nil {
				t.Errorf("Expected non-nil PackCalculator")
			}
		})
	}
}

func Test_packCalculator_CalculatePacks(t *testing.T) {
	logger := testLogger()
	packs := []int{250, 500, 1000, 2000, 5000}

	tests := []struct {
		name    string
		amount  int
		want    *domain.CalculationResult
		wantErr bool
	}{
		// Edge cases
		{
			name:    "Invalid zero amount",
			amount:  0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid negative amount",
			amount:  -100,
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Smallest possible amount (1)",
			amount: 1,
			want: &domain.CalculationResult{
				Packs: map[int]int{250: 1},
			},
			wantErr: false,
		},

		// Exact matches
		{
			name:   "Exact match with 1x250",
			amount: 250,
			want: &domain.CalculationResult{
				Packs: map[int]int{250: 1},
			},
			wantErr: false,
		},
		{
			name:   "Exact match with 1x500",
			amount: 500,
			want: &domain.CalculationResult{
				Packs: map[int]int{500: 1},
			},
			wantErr: false,
		},
		{
			name:   "Exact match with 1x1000",
			amount: 1000,
			want: &domain.CalculationResult{
				Packs: map[int]int{1000: 1},
			},
			wantErr: false,
		},
		{
			name:   "Exact match with multiple packs (750)",
			amount: 750,
			want: &domain.CalculationResult{
				Packs: map[int]int{500: 1, 250: 1},
			},
			wantErr: false,
		},

		// Optimization cases
		{
			name:   "Choose larger single pack over multiple smaller (251)",
			amount: 251,
			want: &domain.CalculationResult{
				Packs: map[int]int{500: 1},
			},
			wantErr: false,
		},
		{
			name:   "Choose larger single pack over multiple smaller (499)",
			amount: 499,
			want: &domain.CalculationResult{
				Packs: map[int]int{500: 1},
			},
			wantErr: false,
		},
		{
			name:   "Mixed pack case (501)",
			amount: 501,
			want: &domain.CalculationResult{
				Packs: map[int]int{500: 1, 250: 1},
			},
			wantErr: false,
		},
		{
			name:   "Mixed pack case (751)",
			amount: 751,
			want: &domain.CalculationResult{
				Packs: map[int]int{1000: 1},
			},
			wantErr: false,
		},
		{
			name:   "Mixed pack case (1001)",
			amount: 1001,
			want: &domain.CalculationResult{
				Packs: map[int]int{1000: 1, 250: 1},
			},
			wantErr: false,
		},

		// Large numbers
		{
			name:   "Large number (12001)",
			amount: 12001,
			want: &domain.CalculationResult{
				Packs: map[int]int{5000: 2, 2000: 1, 250: 1},
			},
			wantErr: false,
		},
		{
			name:   "Large number (25001)",
			amount: 25001,
			want: &domain.CalculationResult{
				Packs: map[int]int{5000: 5, 250: 1},
			},
			wantErr: false,
		},
		{
			name:   "Very large number (100000)",
			amount: 100000,
			want: &domain.CalculationResult{
				Packs: map[int]int{5000: 20},
			},
			wantErr: false,
		},
		{
			name:   "Very large number with remainder (100001)",
			amount: 100001,
			want: &domain.CalculationResult{
				Packs: map[int]int{5000: 20, 250: 1},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &packCalculator{
				packSizes: packs,
				logger:    logger,
			}
			got, err := calc.CalculatePacks(tt.amount)

			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatePacks() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != nil {
				if !reflect.DeepEqual(got.Packs, tt.want.Packs) {
					t.Errorf("CalculatePacks() got = %v, want = %v", got.Packs, tt.want.Packs)
				}
			}
		})
	}
}

func Test_packCalculator_findOptimalCombination(t *testing.T) {
	logger := testLogger()
	packs := []int{250, 500, 1000, 2000, 5000}

	tests := []struct {
		name   string
		amount int
		want   map[int]int
	}{
		// Edge cases
		{
			name:   "Tiny amount",
			amount: 1,
			want:   map[int]int{250: 1},
		},
		{
			name:   "Smallest pack size",
			amount: 250,
			want:   map[int]int{250: 1},
		},

		// Exact matches
		{
			name:   "Exact match with single pack",
			amount: 500,
			want:   map[int]int{500: 1},
		},
		{
			name:   "Exact match with multiple packs",
			amount: 750,
			want:   map[int]int{500: 1, 250: 1},
		},
		{
			name:   "Exact match with big packs",
			amount: 12000,
			want:   map[int]int{5000: 2, 2000: 1},
		},

		// Optimization cases
		{
			name:   "Prefer fewer packs (251)",
			amount: 251,
			want:   map[int]int{500: 1},
		},
		{
			name:   "Prefer fewer packs (751)",
			amount: 751,
			want:   map[int]int{1000: 1},
		},
		{
			name:   "Overage from smallest possible",
			amount: 12001,
			want:   map[int]int{5000: 2, 2000: 1, 250: 1},
		},

		// Large numbers
		{
			name:   "Large number with exact match",
			amount: 50000,
			want:   map[int]int{5000: 10},
		},
		{
			name:   "Large number with remainder",
			amount: 50001,
			want:   map[int]int{5000: 10, 250: 1},
		},
		{
			name:   "Very large number",
			amount: 250000,
			want:   map[int]int{5000: 50},
		},
		{
			name:   "Complex large number",
			amount: 123456,
			want:   map[int]int{5000: 24, 2000: 1, 1000: 1, 500: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &packCalculator{
				packSizes: packs,
				logger:    logger,
			}
			got := calc.findOptimalCombination(tt.amount)
			if !reflect.DeepEqual(got.Packs, tt.want) {
				t.Errorf("findOptimalCombination() = %v, want %v", got.Packs, tt.want)
			}
		})
	}
}
