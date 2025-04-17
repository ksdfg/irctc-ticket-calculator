package ticketpricecalculator

// TicketPriceCalculator is used to calculate the price of a train ticket
type TicketPriceCalculator interface {
	// Calculate will calculate the total ticket price for n passengers travelling in a given coach from one stop to another
	Calculate(coach int, startStop string, endStop string, numberOfPassengers int) (price float64, err error)
}
