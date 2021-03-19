package date_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uniplaces/carbon"

	dateUtils "go-tw/utils/date"
)

func TestShouldVerifyWeekDayIfPassADateOfWeek (t *testing.T) {
	firstDay, _ := carbon.CreateFromDate(2021, time.March, 19, "America/Sao_Paulo")
	isWeekendDay := dateUtils.GetIsWeekendOrWeekeDay(&firstDay.Time)
	assert.Equal(t, false, isWeekendDay)
}


func TestShouldVerifyWeekDayIfPassADateOfWeekend (t *testing.T) {
	firstDay, _ := carbon.CreateFromDate(2021, time.March, 14, "America/Sao_Paulo")
	isWeekendDay := dateUtils.GetIsWeekendOrWeekeDay(&firstDay.Time)
	assert.Equal(t, true, isWeekendDay)
}

