package ticketpricecalculator

// Enum constants for the different types of coaches
const (
	GeneralCoach = iota
	SleeperCoach
)

// GeneralCoachBasePrice is the price a passenger has to pay for travelling one stop in the general coach
const GeneralCoachBasePrice float64 = 20

// SleeperCoachBasePrice is the price a passenger has to pay for travelling one stop in the sleeper coach
const SleeperCoachBasePrice float64 = 40

// Train defines all the properties of a train
type Train struct {
	// Route specifies the list of stations that the train travels through, in the order in which they are visited
	Route           []string
	HasGeneralCoach bool
	HasSleeperCoach bool
}
