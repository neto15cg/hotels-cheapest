package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"

	"go-tw/service/hotel"
)

func main() {
	firstDay, _ := carbon.CreateFromDate(2020, time.March, 16, "America/Sao_Paulo")
	secondDay, _ := carbon.CreateFromDate(2020, time.March, 17, "America/Sao_Paulo")
	thirdDay, _ := carbon.CreateFromDate(2020, time.March, 18, "America/Sao_Paulo")
	reservationDates := []time.Time{firstDay.Time,secondDay.Time,  thirdDay.Time}
	hotelCheapest := hotel.GetMoreCheapHotel(reservationDates, true)

	fmt.Println(hotelCheapest)
}