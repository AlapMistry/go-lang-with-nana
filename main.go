package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"

/*
Alternative syntax for variable
conferenceName := "Go Conference"
*/

var remainingTickets uint = 50

// var bookings = []string{}
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	//for {
	firstName, lastName, email, userTickets := getUserInput()

	/*
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	*/

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		// sendTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		// fmt.Printf("These are all our bookings: %v\n", bookings)
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		// var noTicketsRemaining bool = remainingTickets == 0
		// noTicketsRemaining := remainingTickets == 0
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}

		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}

		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	/*
			else {
				fmt.Println("Your input data is invalid, try again")
			}
			else {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
		}
	*/

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	/*
		fmt.Println("Welcome to", conferenceName, "booking application!")
		fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
		fmt.Println("Get your tickets here to attend.")
	*/
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		/*
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		*/

		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// var bookings = [50]string{}
	// var bookings [50]string
	// bookings := [50]string{}
	// var bookings = []string{}
	// var bookings []string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	/*
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	// bookings[0] = firstName + " " + lastName
	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	/*
		fmt.Printf("The whole array/slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array/slice type: %T\n", bookings)
		fmt.Printf("Array/slice length: %v\n", len(bookings))
	*/

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
