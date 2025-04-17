package ticketpricecalculator

import "slices"

// Service is used to calculate the price of a train ticket for a particular train
type Service struct {
	train Train
}

// NewService will initialize a new Service instance with the details of the given train
func NewService(train Train) *Service {
	service := new(Service)
	service.train = train
	return service
}

// Calculate will calculate the total ticket price for n passengers travelling in a given coach from one stop to another
func (s Service) Calculate(coach int, startStop string, endStop string, numberOfPassengers int) (price float64, err error) {
	// Check if number of passengers is valid
	if numberOfPassengers < 1 {
		err = ErrPassengersLessThanOne
		return
	}

	// Get base price for the ticket on the basis of the coach
	var basePrice float64
	switch coach {
	case GeneralCoach:
		if !s.train.HasGeneralCoach {
			err = ErrNoGeneralCoach
			return
		}
		basePrice = GeneralCoachBasePrice

	case SleeperCoach:
		if !s.train.HasSleeperCoach {
			err = ErrNoSleeperCoach
			return
		}
		basePrice = SleeperCoachBasePrice

	default:
		err = ErrInvalidCoachType
		return
	}

	// Get indices of the start and end stops
	startIndex := slices.Index(s.train.Route, startStop)
	endIndex := slices.Index(s.train.Route, endStop)
	if startIndex < 0 {
		err = ErrStartNotFound
		return
	} else if endIndex < 0 {
		err = ErrEndNotFound
		return
	} else if endIndex <= startIndex {
		err = ErrEndBeforeStart
		return
	}

	// Calculate number of stops the passengers have to make to get to their destination
	numberOfStops := endIndex - startIndex

	// Calculate final ticket price
	price = basePrice * float64(numberOfStops) * float64(numberOfPassengers)
	return
}
