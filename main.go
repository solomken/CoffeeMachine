package main

import "fmt"

var availableWater int = 400
var availableMilk int = 540
var availableCoffeeBeans = 120
var disposableCupsAmount = 9
var moneyAmount int = 550

func makeCoffee(coffeeOption string, availableWater *int, availableMilk *int, availableCoffeeBeans *int, disposableCupsAmount *int, moneyAmount *int) {
	switch coffeeOption {
	case "espresso":
		if *availableWater-250 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if *availableCoffeeBeans-16 < 0 {
			fmt.Println("Sorry, not enough coffee beans!")
		} else if *disposableCupsAmount == 0 {
			fmt.Println("Sorry, not enough disposable cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*availableWater -= 250
			*availableCoffeeBeans -= 16
			*disposableCupsAmount -= 1
			*moneyAmount += 4
		}
	case "latte":
		if *availableWater-350 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if *availableMilk-75 < 0 {
			fmt.Println("Sorry, not enough milk!")
		} else if *availableCoffeeBeans-20 < 0 {
			fmt.Println("Sorry, not enough coffee beans!")
		} else if *disposableCupsAmount == 0 {
			fmt.Println("Sorry, not enough disposable cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*availableWater -= 350
			*availableMilk -= 75
			*availableCoffeeBeans -= 20
			*disposableCupsAmount -= 1
			*moneyAmount += 7
		}
	case "cappuccino":
		if *availableWater-200 < 0 {
			fmt.Println("Sorry, not enough water!")
		} else if *availableMilk-100 < 0 {
			fmt.Println("Sorry, not enough milk!")
		} else if *availableCoffeeBeans-12 < 0 {
			fmt.Println("Sorry, not enough coffee beans!")
		} else if *disposableCupsAmount == 0 {
			fmt.Println("Sorry, not enough disposable cups!")
		} else {
			fmt.Println("I have enough resources, making you a coffee!")
			*availableWater -= 200
			*availableMilk -= 100
			*availableCoffeeBeans -= 12
			*disposableCupsAmount -= 1
			*moneyAmount += 6
		}
	}
}

func selectCoffeeOption(coffeeOption string, availableWater *int, availableMilk *int, availableCoffeeBeans *int, disposableCupsAmount *int, moneyAmount *int) {
	switch coffeeOption {
	case "espresso":
		makeCoffee("espresso", availableWater, availableMilk, availableCoffeeBeans, disposableCupsAmount, moneyAmount)
	case "latte":
		makeCoffee("latte", availableWater, availableMilk, availableCoffeeBeans, disposableCupsAmount, moneyAmount)
	case "cappuccino":
		makeCoffee("cappuccino", availableWater, availableMilk, availableCoffeeBeans, disposableCupsAmount, moneyAmount)
	}
}

func doRefill(resourceOption string, desiredAmount int, availableWater *int, availableMilk *int, availableCoffeeBeans *int, disposableCupsAmount *int) {
	switch resourceOption {
	case "water":
		*availableWater += desiredAmount
	case "milk":
		*availableMilk += desiredAmount
	case "coffee beans":
		*availableCoffeeBeans += desiredAmount
	case "disposable cups":
		*disposableCupsAmount += desiredAmount
	default:
		fmt.Println("invalid option")
	}
}

func doWithdrawal(moneyAmount *int) {
	fmt.Printf("I gave you $%d\n", *moneyAmount)
	*moneyAmount = 0
}

func getUserChoice() string {
	var choice string
	fmt.Scan(&choice)
	return choice
}

func showResources() {
	fmt.Printf(`The coffee machine has:
	%d ml of water
	%d ml of milk
	%d g of coffee beans
	%d disposable cups
	$%d of money
`, availableWater, availableMilk, availableCoffeeBeans, disposableCupsAmount, moneyAmount)
}

func doAction(action string) {
	var desiredAmount int

	switch action {
	case "buy":
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
		coffeeOption := getUserChoice()
		switch coffeeOption {
		case "1":
			selectCoffeeOption("espresso", &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount, &moneyAmount)
			fmt.Println()
			startProgram()
		case "2":
			selectCoffeeOption("latte", &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount, &moneyAmount)
			fmt.Println()
			startProgram()
		case "3":
			selectCoffeeOption("cappuccino", &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount, &moneyAmount)
			fmt.Println()
			startProgram()
		case "back":
			fmt.Println()
			startProgram()
		default:
			fmt.Println("invalid option")
			fmt.Println()
			startProgram()
		}
	case "fill":
		fmt.Println("Write how many ml of water you want to add:")
		fmt.Scan(&desiredAmount)
		doRefill("water", desiredAmount, &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount)
		fmt.Println("Write how many ml of milk you want to add:")
		fmt.Scan(&desiredAmount)
		doRefill("milk", desiredAmount, &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount)
		fmt.Println("Write how many grams of coffee beans you want to add:")
		fmt.Scan(&desiredAmount)
		doRefill("coffee beans", desiredAmount, &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount)
		fmt.Println("Write how many disposable cups you want to add:")
		fmt.Scan(&desiredAmount)
		doRefill("disposable cups", desiredAmount, &availableWater, &availableMilk, &availableCoffeeBeans, &disposableCupsAmount)
		fmt.Println()
		startProgram()
	case "take":
		doWithdrawal(&moneyAmount)
		fmt.Println()
		startProgram()
	case "remaining":
		showResources()
		fmt.Println()
		startProgram()
	case "exit":
		return
	case "42":
		fmt.Println()
		showEasterEgg()
	default:
		fmt.Println("invalid option")
		fmt.Println()
		startProgram()
	}
}

func showEasterEgg() {
	fmt.Println("what's the name of Golang mascot?")
	action := getUserChoice()
	if action == "gopher" || action == "Gopher" {
		fmt.Println("free coffee for you bro! <3")
		fmt.Println()
		startProgram()
	} else {
		fmt.Println("banned forever!!!1")
		return
	}
}

func startProgram() {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	action := getUserChoice()
	fmt.Println()
	doAction(action)
	fmt.Println()
}
func main() {
	startProgram()
}
