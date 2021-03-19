package hotel_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uniplaces/carbon"

	"go-tw/service/hotel"
)

func TestShouldReturnParqueDasFlores(t *testing.T) {
	firstDay, _ := carbon.CreateFromDate(2020, time.March, 16, "America/Sao_Paulo")
	secondDay, _ := carbon.CreateFromDate(2020, time.March, 17, "America/Sao_Paulo")
	thirdDay, _ := carbon.CreateFromDate(2020, time.March, 18, "America/Sao_Paulo")
	reservationDates := []time.Time{firstDay.Time,secondDay.Time,  thirdDay.Time}
	hotelCheapest := hotel.GetCheapestHotel(reservationDates, true)
	assert.Equal(t, hotel.PARQUE_DAS_FLORES, hotelCheapest)
}

func TestShouldReturnJardimBotanico(t *testing.T) {
	firstDay, _ := carbon.CreateFromDate(2020, time.March, 20, "America/Sao_Paulo")
	secondDay, _ := carbon.CreateFromDate(2020, time.March, 21, "America/Sao_Paulo")
	thirdDay, _ := carbon.CreateFromDate(2020, time.March, 22, "America/Sao_Paulo")
	reservationDates := []time.Time{firstDay.Time,secondDay.Time,  thirdDay.Time}
	hotelCheapest := hotel.GetCheapestHotel(reservationDates, true)

	assert.Equal(t, hotel.JARDIM_BOTANICO, hotelCheapest)
}

func TestShouldReturnMarAtlantico(t *testing.T) {
	firstDay, _ := carbon.CreateFromDate(2020, time.March, 26, "America/Sao_Paulo")
	secondDay, _ := carbon.CreateFromDate(2020, time.March, 27, "America/Sao_Paulo")
	thirdDay, _ := carbon.CreateFromDate(2020, time.March, 28, "America/Sao_Paulo")
	reservationDates := []time.Time{firstDay.Time,secondDay.Time,  thirdDay.Time}
	hotelCheapest := hotel.GetCheapestHotel(reservationDates, false)

	assert.Equal(t, hotel.MAR_ATLANTICO, hotelCheapest)
}