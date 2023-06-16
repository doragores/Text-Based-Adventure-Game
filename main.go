package main

import (
	"fmt"
)

type Room struct {
	Name        string
	Description string
}

func navigationMenu() {
	fmt.Println("1. Look left")
	fmt.Println("2. Look right")
	fmt.Println("3. Look forward")
	fmt.Println("4. Exit the room")
}
func main() {
	groundFloor := Room{
		Name:        "Ground Floor",
		Description: "From here you can exit the building to end the game, or get the lift to continue your journey.",
	}

	lift := Room{
		Name:        "Lift",
		Description: "You entered the lift. From here you can exit the lift and go back to the ground floor, got to the 4th floor or go to the 6th floor.",
	}

	hallway4th := Room{
		Name: "Hallway 4th Floor",
		Description: "This is the hallway. From here you can either go back to the lift, to the Bankification box meeting room," +
			" or to the  " + "Disco Ball meeting room",
	}

	hallway6th := Room{
		Name:        "Hallway 6th Floor",
		Description: "This is the hallway. From here you can either go to back to the lift, or the terrace",
	}

	bankificationBox := Room{
		Name:        "Bankification Box meeting room",
		Description: "This is the Bankification box meeting Room. From here you can go only back to the hallway.",
	}

	discoBall := Room{
		Name:        "Disco Ball meeting room",
		Description: "This is the Disco ball meeting room. From here you can only go back to the hallway.",
	}

	terrace := Room{
		Name:        "Terrace",
		Description: "This is the Terrace. From here you can only go back to the hallway.",
	}

	fmt.Println("Hey 👋 \nThis is a text game where you need to find a golden laptop in order to become a software engineer at Monzo," +
		" you can navigate between rooms and look for clues 👀\n" +
		"Are you ready to become a software engineer? 🚀" +
		"[Yes/No]")

	var userInput string

	fmt.Scanln(&userInput)

	if userInput == "Yes" {
		fmt.Println("Here we go! To exit the game at any point, please type `Exit`.\n")
	} else {
		fmt.Println("Maybe next time. Goodbye 👋")
		return
	}
	var currentPlace Room
	currentPlace = groundFloor

	for {

		if userInput == "Exit" {
			fmt.Println("Maybe next time. Goodbye 👋")
			break
		}

		fmt.Println("You are currently in:", currentPlace.Name)
		fmt.Println("Description:", currentPlace.Description+"\n")
		fmt.Println("Where do you want to go?")

		switch currentPlace.Name {

		case "Ground Floor":
			fmt.Println("1. Exit\n2. Lift")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				fmt.Println("Maybe next time. Goodbye 👋")
				return
			} else if userInput == "2" {
				currentPlace = lift
			}
		case "Lift":
			fmt.Println("1. Ground Floor\n2. 4th Floor\n3. 6th Floor")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = groundFloor
			} else if userInput == "2" {
				currentPlace = hallway4th
			} else {
				currentPlace = hallway6th
			}
		case "Hallway 4th Floor":
			fmt.Println("1. Lift\n2. The Bankification Box Meeting Room\n3. The Disco Ball Meeting Room")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = lift
			} else if userInput == "2" {
				currentPlace = bankificationBox
			} else {
				currentPlace = discoBall
			}
		case "Bankification Box meeting room":
			fmt.Println("1. Go back to the hallway\n2. Look around")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = hallway4th
			} else {
				for {
					navigationMenu()
					fmt.Scanln(&userInput)
					if userInput == "1" {
						fmt.Println("\nYou look left and spot a note 📝 The note says:\n`There's a secret in this room...`\n")
					} else if userInput == "2" {
						fmt.Println("\nYou look right, but there's nothing important to see.\n")
					} else if userInput == "3" {
						fmt.Println("\nYou look forward and see another note 📝:\n`To find the secret go to the terrace 👀`\n")
					} else if userInput == "4" {
						currentPlace = hallway4th
						break
					}
				}
			}
		case "Disco Ball meeting room":
			fmt.Println("1. Hallway")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = hallway4th
			}
		case "Hallway 6th Floor":
			fmt.Println("1. Lift\n2. Terrace")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = lift
			} else {
				currentPlace = terrace
			}
		case "Terrace":
			fmt.Println("1. Hallway")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = hallway6th
			}
		}
	}
}
