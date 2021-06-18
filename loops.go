package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// House holds all info about one specifice home
type House struct {
	address       string
	city          string
	numberOfRooms int
	price         int
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of houses: ")
	houseNum, _ := reader.ReadString('\n')
	houseNum = houseNum[:len(houseNum)-1]
	houseInt, _ := strconv.Atoi(houseNum)

	houseArr := []House{}

	for houseInt > 0 {
		var ahouse House

		// address
		fmt.Printf("Enter the address for house %v : ", houseInt)
		ahouse.address, _ = reader.ReadString('\n')
		ahouse.address = ahouse.address[:len(ahouse.address)-1]

		// city
		fmt.Printf("Enter the city for house %v : ", houseInt)
		ahouse.city, _ = reader.ReadString('\n')
		ahouse.city = ahouse.city[:len(ahouse.city)-1]

		// numberOfRooms
		fmt.Printf("Enter the numberOfRooms for house %v : ", houseInt)
		numberOfRooms, _ := reader.ReadString('\n')
		numberOfRooms = numberOfRooms[:len(numberOfRooms)-1]
		numberOfRoomsInt, _ := strconv.Atoi(numberOfRooms)

		ahouse.numberOfRooms = numberOfRoomsInt

		// price
		fmt.Printf("Enter the price for house %v : ", houseInt)
		price, _ := reader.ReadString('\n')
		price = price[:len(price)-1]
		priceInt, _ := strconv.Atoi(price)

		ahouse.numberOfRooms = priceInt

		houseArr = append(houseArr, ahouse)

		houseInt--
	}

	fmt.Println(houseArr)
}
