package hotel

import (
	"time"
)

const (
	PARQUE_DAS_FLORES = "Parque das flores"
	JARDIM_BOTANICO = "Jardim Botânico"
	MAR_ATLANTICO = "Mar Atlântico"
)

type Hotel struct {
	name 				 string
	regularPriceWeek 	 float64
	fidelityPriceWeek 	 float64
	regularPriceWeekend  float64
	fidelityPriceWeekend float64
	stars				 int
}

type HotelPricesResumes struct {
	name 	   string
	totalPrice float64
	stars int
}

var HotelList []Hotel = []Hotel{
	{PARQUE_DAS_FLORES, 110, 80, 90, 80, 3},
	{JARDIM_BOTANICO, 160, 110, 60, 50, 4},
	{MAR_ATLANTICO, 220, 100, 150, 40, 5},
}

func GetCheapestHotel(reservationDates []time.Time, isRegular bool) string  {
	var hotelCheapest HotelPricesResumes

	for i, hotel := range HotelList {
		price := getHotelPrice(reservationDates, hotel.regularPriceWeek, hotel.regularPriceWeekend, hotel.fidelityPriceWeek, hotel.fidelityPriceWeekend, isRegular)
		hotelResume := HotelPricesResumes{hotel.name, price, hotel.stars}

		if i == 0 {
			hotelCheapest = hotelResume
			continue
		}
		if hotelResume.totalPrice < hotelCheapest.totalPrice {
			hotelCheapest = hotelResume
			continue
		} 
		if hotelResume.totalPrice == hotelCheapest.totalPrice && hotel.stars > hotelCheapest.stars {
			hotelCheapest = hotelResume	
		}
	}
	return hotelCheapest.name
}


func getHotelPrice(reservationDates []time.Time, regularPriceWeek, regularPriceWeekend, fidelityPriceWeek, fidelityPriceWeekend float64, isRegular bool) float64 {
	var totalPrice float64 = 0
	for _, date := range reservationDates {
		dayOfWeek := int(date.Weekday())
		isWeekendPrice := getIsWeekendPrice(dayOfWeek);
		price := getWeekPriceOrWeekendPrice(regularPriceWeek, regularPriceWeekend, fidelityPriceWeek, fidelityPriceWeekend, isRegular, isWeekendPrice)
		totalPrice = price + totalPrice		
	}
	return totalPrice
}


func getWeekPriceOrWeekendPrice(regularPriceWeek, regularPriceWeekend, fidelityPriceWeek, fidelityPriceWeekend float64, isRegular, isWeekendPrice bool) float64 {
	if isRegular && isWeekendPrice {
		return regularPriceWeekend
	}
	if isRegular && !isWeekendPrice {
		return regularPriceWeek
	}
	if isWeekendPrice {
		return fidelityPriceWeekend
	}
	return fidelityPriceWeek

}

func getIsWeekendPrice(dayOfWeek int) bool {
	if dayOfWeek == 0 || dayOfWeek == 6 {
		return true 
	}
	return false
}