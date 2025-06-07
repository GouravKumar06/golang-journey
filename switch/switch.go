package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	var day int = 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
		case time.Saturday,time.Sunday: {
			fmt.Println("Weekend")
		}
		default: {
			fmt.Println("Weekday")
		}
    }

	//type switch
	whoAmI := func(i any){
        switch i.(type){
			case bool: {
				fmt.Println("bool")
			}
			case int: {
				fmt.Println("int")
			}
			case string: {
				fmt.Println("string")
			}
			default: {
				fmt.Println("unknown")
			}
		}
	}

	whoAmI(true)
	whoAmI(1)
	whoAmI("gourav")

}