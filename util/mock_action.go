package util

import (
	"fmt"
	"math/rand"
	"strconv"
)


func RandomINT32String(min, max int) string {

	var i int32 = int32(rand.Intn(max - min) + min)

	return fmt.Sprintf("%d", i)
}
func RandomNameAction() string {
	return RandomString(10)
}

func RandomIsinAction() string {
	return RandomString(10)
}

func RandomIDAction() int32 {	
	convID, err := strconv.ParseInt(RandomINT32String(1, 10), 10, 64)
	if err != nil {
		panic(err)
	}
	return  int32(convID)
}

func RandomWknAction() string {
	return RandomINT32String(1, 10) + RandomString(10)
}

func RandomCurrentValueAction() string {
	//format to value equal to 1.00

	convCurrentValue, err := strconv.ParseInt(RandomINT32String(1, 100), 10, 64)
	if err != nil {
		panic(err)
	}

	var currentValue string = fmt.Sprintf("%d", convCurrentValue) + ".00"

	return currentValue 	
}

func RandomBidAction() string {
	return RandomINT32String(1, 10)
}

func RandomAskAction() string {
	return RandomINT32String(1, 10)
}

func RandomSpreadAction(bid, ask int64) string {
	return fmt.Sprintf("%d", ask - bid)
}

func RandomChangePercentageAction(currentValue string, lastValue string) string {

	convCurrentValue, err := strconv.ParseInt(currentValue, 10, 64)
	if err != nil {
		panic(err)
	}

	convLastValue, err := strconv.ParseInt(lastValue, 10, 64)
		if err != nil {
		panic(err)
	}

	var changePercentage int64 = 100 * (convCurrentValue - convLastValue) / convLastValue

	return fmt.Sprintf("%d", changePercentage)
}

func RandomChangeAbsoluteAction(currentValue string, lastValue string) string {

	convCurrentValue, err := strconv.ParseInt(currentValue, 10, 64)
	if err != nil {
		panic(err)
	}

	convLastValue, err := strconv.ParseInt(lastValue, 10, 64)
		if err != nil {
		panic(err)
	}

	var absoluteChange int64 = convCurrentValue - convLastValue

	return fmt.Sprintf("%d", absoluteChange)
}

func RandomPeak24hAction() string {
	return RandomINT32String(1, 10)
}

func RandomPeak7dAction() string {
	return RandomINT32String(1, 10)
}

func RandomPeak30dAction() string {
	return RandomINT32String(1, 10)
}

func RandomLow24hAction() string {
	return RandomINT32String(1, 10)
}

func RandomLow7dAction() string {
	return RandomINT32String(1, 10)
}

func RandomLow30dAction() string {
	return RandomINT32String(1, 10)
}







