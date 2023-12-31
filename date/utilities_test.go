package date

import (
	"testing"
	"time"
)

const desc = "Expected %v but it got %v"

func TestToSecond(t *testing.T) {
	second := 1
	expected := time.Duration(second) * time.Second
	actual := ToSecond(second)
	if expected != actual {
		t.Errorf(desc, expected, actual)
	}
}

func TestToMinute(t *testing.T) {
	minute := 1
	expected := time.Duration(minute) * time.Minute
	actual := ToMinute(minute)
	if expected != actual {
		t.Errorf(desc, expected, actual)
	}
}

func TestToHour(t *testing.T) {
	hour := 1
	expected := time.Duration(hour) * time.Hour
	actual := ToHour(hour)
	if expected != actual {
		t.Errorf(desc, expected, actual)
	}
}

func TestSleepSecond(t *testing.T) {
	nowSecond := time.Now().Second()
	second := 1
	SleepSecond(second)
	if nowSecond == time.Now().Second() {
		t.Errorf("Expected %v but it got %v", nowSecond+second, time.Now().Second())
	}
}

// skip this test because it will take a long time
// func TestSleepMinute(t *testing.T) {
// 	nowMinute := time.Now().Minute()
// 	minute := 1

// 	SleepMinute(minute)
// 	if nowMinute == time.Now().Minute() {
// 		t.Errorf("Expected %v but it got %v", nowMinute+minute, time.Now().Minute())
// 	}
// }

// skip this test because it will take a long time
// func TestSleepHour(t *testing.T) {
// 	nowHour := time.Now().Hour()
// 	hour := 1
// 	SleepHour(hour)
// 	if nowHour == time.Now().Hour() {
// 		t.Errorf("Expected %v but it got %v", nowHour+hour, time.Now().Hour())
// 	}
// }
