package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// House holds all info about one specifice home
type House struct {
	Address       string
	City          string
	NumberOfRooms int
	Price         int
}

func main() {
	var houseArr []House

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number of houses: ")
	houseNum, _ := reader.ReadString('\n')
	houseNum = houseNum[:len(houseNum)-1]
	houseInt, _ := strconv.Atoi(houseNum)

	for houseInt > 0 {
		var house House

		// address
		fmt.Printf("Enter the address for house %v : ", houseInt)
		address, _ := reader.ReadString('\n')
		house.Address = address[:len(address)-1]

		// city
		fmt.Printf("Enter the city for house %v : ", houseInt)
		city, _ := reader.ReadString('\n')
		house.City = city[:len(city)-1]

		// numberOfRooms
		fmt.Printf("Enter the numberOfRooms for house %v : ", houseInt)
		numberOfRooms, _ := reader.ReadString('\n')
		numberOfRooms = numberOfRooms[:len(numberOfRooms)-1]
		numberOfRoomsInt, _ := strconv.Atoi(numberOfRooms)
		house.NumberOfRooms = numberOfRoomsInt

		// price
		fmt.Printf("Enter the price for house %v : ", houseInt)
		price, _ := reader.ReadString('\n')
		price = price[:len(price)-1]
		priceInt, _ := strconv.Atoi(price)
		house.Price = priceInt

		// add house to array
		houseArr = append(houseArr, house)
		houseInt--
	}

	for _, house := range houseArr {
		fmt.Printf("%s\t %s\t %d\t %d\n",
			house.Address,
			house.City,
			house.NumberOfRooms,
			house.Price)
	}
}
