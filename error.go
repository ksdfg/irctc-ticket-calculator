package ticketpricecalculator

import "errors"

var (
	ErrNoGeneralCoach        = errors.New("train doesn't have general coach")
	ErrNoSleeperCoach        = errors.New("train doesn't have sleeper coach")
	ErrInvalidCoachType      = errors.New("invalid coach type")
	ErrStartNotFound         = errors.New("train does not go to the given start stop")
	ErrEndNotFound           = errors.New("train does not go to the given end stop")
	ErrEndBeforeStart        = errors.New("end stop cannot come before start stop")
	ErrPassengersLessThanOne = errors.New("number of passengers cannot be less than 1")
)
