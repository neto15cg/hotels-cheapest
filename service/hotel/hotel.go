package hotel

import (
	dateUtils "go-tw/utils/date"
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

	for index, hotel := range HotelList {
		price := GetTotalHotelPrice(reservationDates, &hotel, isRegular)
		hotelResume := HotelPricesResumes{hotel.name, price, hotel.stars}

		if index == 0 {
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


func GetTotalHotelPrice(reservationDates []time.Time, hotel *Hotel, isRegular bool) float64 {
	var totalPrice float64 = 0
	for _, date := range reservationDates {
		isWeekendPrice := dateUtils.GetIsWeekendOrWeekeDay(&date);
		price := GetHotelWeekOrWeekendPrice(hotel, isRegular, isWeekendPrice)
		totalPrice = price + totalPrice		
	}
	return totalPrice
}


func GetHotelWeekOrWeekendPrice(hotel *Hotel, isRegular, isWeekendPrice bool) float64 {
	if isRegular && isWeekendPrice {
		return hotel.regularPriceWeekend
	}
	if isRegular && !isWeekendPrice {
		return hotel.regularPriceWeek
	}
	if isWeekendPrice {
		return hotel.fidelityPriceWeekend
	}
	return hotel.fidelityPriceWeek
}