package ticketpricecalculator

import (
	"errors"
	"testing"
)

func TestService(t *testing.T) {
	type testCaseInput struct {
		coach              int
		startStop          string
		endStop            string
		numberOfPassengers int
	}

	type testCaseOutput struct {
		price float64
		err   error
	}

	type testCase struct {
		train Train
		input testCaseInput
		want  testCaseOutput
	}

	stops := make([]string, 5)
	stops = append(stops, "Mumbai", "Karjat", "Lonavala", "Chinchwad", "Pune")

	testCases := map[string]testCase{
		"happy case single passenger": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: GeneralCoach, startStop: "Karjat", endStop: "Chinchwad", numberOfPassengers: 1},
			want:  testCaseOutput{price: 40, err: nil},
		},
		"happy case multiple passengers": {
			train: Train{Route: stops, HasSleeperCoach: true},
			input: testCaseInput{coach: SleeperCoach, startStop: "Mumbai", endStop: "Pune", numberOfPassengers: 3},
			want:  testCaseOutput{price: 480, err: nil},
		},
		"invalid coach type": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: 2, startStop: "Karjat", endStop: "Chinchwad", numberOfPassengers: 1},
			want:  testCaseOutput{price: 0, err: ErrInvalidCoachType},
		},
		"train has no sleeper coach": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: SleeperCoach, startStop: "Karjat", endStop: "Chinchwad", numberOfPassengers: 1},
			want:  testCaseOutput{price: 0, err: ErrNoSleeperCoach},
		},
		"train has no general coach": {
			train: Train{Route: stops, HasSleeperCoach: true},
			input: testCaseInput{coach: GeneralCoach, startStop: "Karjat", endStop: "Chinchwad", numberOfPassengers: 1},
			want:  testCaseOutput{price: 0, err: ErrNoGeneralCoach},
		},
		"start stop missing": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: GeneralCoach, startStop: "Panvel", endStop: "Chinchwad", numberOfPassengers: 1},
			want:  testCaseOutput{price: 0, err: ErrStartNotFound},
		},
		"end stop missing": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: GeneralCoach, startStop: "Karjat", endStop: "Panvel", numberOfPassengers: 1},
			want:  testCaseOutput{price: 0, err: ErrEndNotFound},
		},
		"end stop before start stop": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: GeneralCoach, startStop: "Chinchwad", endStop: "Karjat", numberOfPassengers: 1},
			want:  testCaseOutput{price: 0, err: ErrEndBeforeStart},
		},
		"number of passengers less than one": {
			train: Train{Route: stops, HasGeneralCoach: true},
			input: testCaseInput{coach: GeneralCoach, startStop: "Karjat", endStop: "Chinchwad", numberOfPassengers: 0},
			want:  testCaseOutput{price: 0, err: ErrPassengersLessThanOne},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			service := NewService(tc.train)

			price, err := service.Calculate(tc.input.coach, tc.input.startStop, tc.input.endStop, tc.input.numberOfPassengers)

			if !errors.Is(err, tc.want.err) {
				t.Fatalf("incorrect error! expected: %v, got: %v", tc.want.err, err)
			}

			if price != tc.want.price {
				t.Fatalf("incorrect price! expected: %v, got: %v", tc.want.price, price)
			}
		})
	}
}
