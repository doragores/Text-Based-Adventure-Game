package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Room struct {
	Name        string
	Description string
	IsLocked    bool
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func navigationMenu() {
	fmt.Println("1. Look left")
	fmt.Println("2. Look right")
	fmt.Println("3. Look forward")
	fmt.Println("4. Exit the room")
}
func main() {

	clearScreen()
	var backpack []string

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
			" or to the " + "Disco Ball meeting room",
	}

	hallway6th := Room{
		Name:        "Hallway 6th Floor",
		Description: "This is the hallway. From here you can either go to back to the lift, or the terrace",
	}

	bankificationBox := Room{
		Name:        "Bankification Box meeting room",
		Description: "This is the Bankification box meeting Room. From here you can go only back to the hallway or have a look around.",
	}

	discoBall := Room{
		Name:        "Disco Ball meeting room",
		Description: "This is the Disco ball meeting room.",
		IsLocked:    true,
	}

	terrace := Room{
		Name:        "Terrace",
		Description: "This is the Terrace. From here you can only go back to the hallway or have a look around.",
	}

	fmt.Println("Hey 👋 \nThis is a text game where you need to find a golden laptop in order to become a software engineer at Monzo," +
		" you can navigate between rooms and look for clues 👀\n" +
		"Are you ready to become a software engineer? 🚀" +
		"[Yes/No]")

	var userInput string
	fmt.Scanln(&userInput)
	userInput = strings.ToLower(userInput)
	clearScreen()

	switch userInput {
	case "yes":
		fmt.Println("Here we go! 🚀yes To exit the game at any point, please type `Exit`.\n")
	case "no":
		fmt.Println("Maybe next time. Goodbye 👋")
	default:
		fmt.Println("Invalid input. Please enter 'Yes' or 'No'.")
	}

	var currentPlace Room
	currentPlace = groundFloor

	for {
		if userInput == "backpack" {
			fmt.Printf("Here's what you have in your backpack 🎒 %v\n", backpack)
		}
		if userInput == "Exit" {
			fmt.Println("Maybe next time. Goodbye 👋")
			break
		}

		fmt.Println("You are currently in:", currentPlace.Name)
		fmt.Println("Description:", currentPlace.Description+"\n")
		if currentPlace != discoBall {
			fmt.Println("Where do you want to go! \nYou currently have a backpack with you and if you want to look at what you have in it, just type `backpack` at any point.")
		}

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
			clearScreen()

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
			clearScreen()

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
			clearScreen()

		case "Bankification Box meeting room":
			fmt.Println("1. Go back to the hallway\n2. Look around")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = hallway4th
			} else {
				for {
					navigationMenu()
					fmt.Scanln(&userInput)
					clearScreen()
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
			clearScreen()
		case "Disco Ball meeting room":
			if contains(backpack, "key") {
				discoBall.IsLocked = false
				fmt.Println("You've unlocked the room with the key in your backpack!🔓Do you wanna take a look around? [Yes/No]")
				fmt.Scanln(&userInput)
				if userInput == "Yes" {
					navigationMenu()
					fmt.Scanln(&userInput)
					clearScreen()
				}
			} else {
				clearScreen()
				fmt.Println("This room is locked 🔐\n")
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
			clearScreen()
		case "Terrace":
			fmt.Println("1. Go back to the hallway\n2. Look around")
			fmt.Scanln(&userInput)
			if userInput == "1" {
				currentPlace = hallway6th
			} else {
				for {
					navigationMenu()
					fmt.Scanln(&userInput)
					if userInput == "1" {
						fmt.Println("\nYou look left and see the sunset 🌇.It's beautiful isn't it...\n")
					} else if userInput == "2" {
						fmt.Println("\nYou look right, but there's nothing important to see, just a few coworkers chatting and having drinks.\nYou realise you could actually use " +
							"a drink, so you grab a can of coke and put into your backpack 🎒")
						item := "coke"
						backpack = append(backpack, item)
					} else if userInput == "3" {
						fmt.Println("\nYou look forward and spot a shinny object... As you look closer you notice it's a key! 🔑 " +
							"Do you want to put it in your backpack? [Yes/No]\n")
						fmt.Scanln(&userInput)
						clearScreen()
						if userInput == "Yes" {
							item := "key"
							backpack = append(backpack, item)
							fmt.Println("The key has been added to your backpack 🎒")
						}
						clearScreen()
					} else if userInput == "4" {
						currentPlace = hallway6th
						break
						clearScreen()
					}
				}
			}
		}
	}
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}
