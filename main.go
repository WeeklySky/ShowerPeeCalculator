package main


import (
	"flag"
	"fmt"
)

func main()  {
	showerFlowMinute := flag.Float64("shower", 9.5, "Shower water flow in liters per minute.")

	flushFlow := flag.Float64("flush", 4, "Water flow in liters in a single flush.")

	flag.Parse()

	showerFlowSeconds := *showerFlowMinute / 60;

	fmt.Print("Enter how long did your pee last (in seconds): ")
	var peeTime float64
	fmt.Scanln(&peeTime)

	waterSpent := peeTime * showerFlowSeconds;

	if waterSpent < *flushFlow {
		fmt.Printf("You saved %.2f liters of water by peeing in the shower. Good job I guess.\n", *flushFlow - waterSpent)
	}

	if waterSpent >= *flushFlow {
		fmt.Println("You are wasting water!")
		fmt.Printf("By peeing in the shower for longer than %.0f seconds, contrary to popular belief, you are actually wasting water.\n", *flushFlow / showerFlowSeconds)
	}
}