package tickets

import (
	"strings"
	"time"
)


// ticket represents every flight ticket found in the external
// csv file that match with the struct model attributes order.
//  
// Example:
// 
// good csv file row content order should be: 
// 1,Steve Musk,stevemusk@etsy.com,Colombia,20:44,550
type ticket struct {
	// id is the flight ticket id -> [0].
	id string

	// name is the ticket passanger name -> [1].
	name string

	// email is the ticket passanger email -> [2].
	email string 

	// destination is the flight ticket destination -> [3].
	destination string

	// flightTime is the time of the flight -> [4].
	flightTime time.Time

	// price is the ticket price -> [5].
	price float64
}

// Tickets represents an slice containing all ticket structs.
// This is use to manipulate and query tickets struct data.
type Tickets []ticket

// GetTicketsAmountByDestination calculates the amount of flight tickets
// going to a specfic destination.
func (t Tickets) GetTicketsAmountByDestination(destination string) int {
	var amount int	
	for _, ticket := range t {
		if strings.ToLower(destination) == strings.ToLower(ticket.destination) {
			amount++
		}
	}
	return amount
}

func (t Tickets) GetTicketsAmountByTimeRange(startTime time.Time, endTime time.Time) int {
	var amount int
	for _, ticket := range t {
		if ticket.flightTime.After(startTime) && ticket.flightTime.Before(endTime) {
			amount++
		}
	}
	return amount
}

func (t Tickets) GetTicketsAverageByDestination(destination string) float64 {
	var amount float64
	for _, ticket := range t {
		if strings.ToLower(destination) == strings.ToLower(ticket.destination) {
			amount++
		}
	}
	return amount / float64(len(t))
}
