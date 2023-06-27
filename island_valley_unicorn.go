package main

import "fmt"

// A Brighter Place
func main() {
	// 1. Print a Welcome Message
	fmt.Println("Welcome to A Brighter Place!")

	// 2. Print a Description of the Place
	fmt.Println("A Brighter Place is a place of hope and light.")

	// 3. Variables to Store Hope and Light Values
	var hopeValue int = 10
	var lightValue int = 5

	// 4. Print the Hope Value
	fmt.Println("Hope Value:", hopeValue)

	// 5. Print the Light Value
	fmt.Println("Light Value:", lightValue)

	// 6. Create a Function to Increase Hope and Light
	increaseHopeAndLight := func(hope int, light int) (int, int) {
		hope += 5
		light += 5
		return hope, light
	}

	// 7. Call the Function
	hopeValue, lightValue = increaseHopeAndLight(hopeValue, lightValue)

	// 8. Print the New Hope and Light Values
	fmt.Println("New Hope Value:", hopeValue)
	fmt.Println("New Light Value:", lightValue)

	// 9. Create a For Loop to Create Increasing Hope and Light
	for i := 0; i < 10; i++ {
		hopeValue, lightValue = increaseHopeAndLight(hopeValue, lightValue)
		fmt.Printf("Hope and Light Values: %v, %v\n", hopeValue, lightValue)
	}

	// 10. Create a Struct to Store Values
	type Brightness struct {
		hope int
		light int
	}

	// 11. Create an Instance of the Struct
	brighterPlace := Brightness{
		hope:   hopeValue,
		light: lightValue,
	}

	// 12. Print the Struct Values
	fmt.Println("Hope and Light in A Brighter Place:", brighterPlace)

	// 13. Create a Method to Increase Values
	func (b *Brightness) increase() {
		b.hope += 5
		b.light += 5
	}

	// 14. Call the Method
	brighterPlace.increase()

	// 15. Print the New Values
	fmt.Println("Hope and Light in A Brighter Place:", brighterPlace)

	// 16. Create a Slice of Structs
	brightnessList := []Brightness{
		{hope: 10, light: 5},
		{hope: 15, light: 10},
		{hope: 20, light: 15},
		{hope: 25, light: 20},
	}

	// 17. Print the Slice
	fmt.Println("Hope and Light in A Brighter Place List:", brightnessList)

	// 18. Create a Function to Increase List Values
	increaseListValues := func(list []Brightness) []Brightness {
		for i := range list {
			list[i].increase()
		}
		return list
	}

	// 19. Call the Function
	brightnessList = increaseListValues(brightnessList)

	// 20. Print the New List Values
	fmt.Println("Hope and Light in A Brighter Place List:", brightnessList)

	// 21. Create a Map of Structs
	brightnessMap := map[string]Brightness{
		"morning":     {hope: 10, light: 5},
		"afternoon":   {hope: 15, light: 10},
		"evening":     {hope: 20, light: 15},
		"night":       {hope: 25, light: 20},
	}

	// 22. Print the Map
	fmt.Println("Hope and Light in A Brighter Place Map:", brightnessMap)

	// 23. Create a Function to Increase Map Values
	increaseMapValues := func(m map[string]Brightness) map[string]Brightness {
		for k := range m {
			m[k].increase()
		}
		return m
	}

	// 24. Call the Function
	brightnessMap = increaseMapValues(brightnessMap)

	// 25. Print the New Map Values
	fmt.Println("Hope and Light in A Brighter Place Map:", brightnessMap)

	// 26. Create a Pointer to the Struct
	brighterPlacePointer := &Brightness{
		hope:   30,
		light: 25,
	}

	// 27. Print the Pointer
	fmt.Println("Hope and Light in A Brighter Place Pointer:", brighterPlacePointer)

	// 28. Create a Function to Increase Pointer Values
	increasePointerValues := func(b *Brightness) {
		b.hope += 5
		b.light += 5
	}

	// 29. Call the Function
	increasePointerValues(brighterPlacePointer)

	// 30. Print the New Pointer Values
	fmt.Println("Hope and Light in A Brighter Place Pointer:", brighterPlacePointer)

	// 31. Create an Interface
	type BrightnessInterface interface {
		hope() int
		light() int
	}

	// 32. Implement the Interface
	type brighterPlaceInterface struct {
		hope   int
		light int
	}

	// 33. Create a Function for the Interface
	func (b brighterPlaceInterface) hope() int {
		return b.hope
	}

	// 34. Create a Function for the Interface
	func (b brighterPlaceInterface) light() int {
		return b.light
	}

	// 35. Create an Instance of the Interface
	brighterPlaceInt := brighterPlaceInterface{
		hope:   35,
		light: 30,
	}

	// 36. Create a Function to Increase the Interface Values
	increaseInterfaceValues := func(b BrightnessInterface) {
		b.hope() += 5
		b.light() += 5
	}

	// 37. Call the Function
	increaseInterfaceValues(brighterPlaceInt)

	// 38. Print the New Interface Values
	fmt.Println("Hope and Light in A Brighter Place Interface:", brighterPlaceInt)

	// 39. Create an Array
	brightnessArray := [4]Brightness{
		{hope: 10, light: 5},
		{hope: 15, light: 10},
		{hope: 20, light: 15},
		{hope: 25, light: 20},
	}

	// 40. Print the Array
	fmt.Println("Hope and Light in A Brighter Place Array:", brightnessArray)

	// 41. Create a Function to Increase Array Values
	increaseArrayValues := func(list [4]Brightness) [4]Brightness {
		for i := 0; i < len(list); i++ {
			list[i].increase()
		}
		return list
	}

	// 42. Call the Function
	brightnessArray = increaseArrayValues(brightnessArray)

	// 43. Print the New Array Values
	fmt.Println("Hope and Light in A Brighter Place Array:", brightnessArray)

	// 44. Create a Channel
	brightnessChannel := make(chan Brightness)

	// 45. Send Values to the Channel
	go func() {
		brightnessChannel <- Brightness{hope: 10, light: 5}
		brightnessChannel <- Brightness{hope: 15, light: 10}
		brightnessChannel <- Brightness{hope: 20, light: 15}
		brightnessChannel <- Brightness{hope: 25, light: 20}
		close(brightnessChannel)
	}()

	// 46. Receive Values from the Channel
	list := []Brightness{}
	for elem := range brightnessChannel {
		list = append(list, elem)
	}

	// 47. Print the Channel Values
	fmt.Println("Hope and Light in A Brighter Place Channel:", list)

	// 48. Create a Function to Increase Channel Values
	increaseChannelValues := func(list []Brightness) []Brightness {
		for i := range list {
			list[i].increase()
		}
		return list
	}

	// 49. Call the Function
	list = increaseChannelValues(list)

	// 50. Print the New Channel Values
	fmt.Println("Hope and Light in A Brighter Place Channel:", list)

	// 51. Create an Error
	err := fmt.Errorf("There is no hope and light in this place.")

	// 52. Check for the Error
	if err != nil {
		fmt.Println(err)
	}

	// 53. Print an Encouraging Message
	fmt.Println("No matter how dark it seems, there is always a brighter place that you can work towards.")

	// 54. Create a Function to Check if There is Hope
	checkForHope := func(hope int, light int) bool {
		if hope > 0 && light > 0 {
			return true
		}
		return false
	}

	// 55. Call the Function
	hopeExists := checkForHope(hopeValue, lightValue)

	// 56. Print the Result
	if hopeExists {
		fmt.Println("Yes! There is hope in A Brighter Place!")
	} else {
		fmt.Println("No, there is no hope in A Brighter Place.")
	}

	// 57. Create an Array of Structs
	brightnessValueList := [4]struct {
		hope     int
		light int
	}{
		{hope: 10, light: 5},
		{hope: 15, light: 10},
		{hope: 20, light: 15},
		{hope: 25, light: 20},
	}

	// 58. Print the Array of Structs
	fmt.Println("Hope and Light in A Brighter Place Array of Structs:", brightnessValueList)

	// 59. Create a Function to Increase Array of Structs Values
	increaseValueListValues := func(list [4]struct {
		hope     int
		light int
	}) [4]struct {
		hope     int
		light int
	} {
		for i := 0; i < len(list); i++ {
			list[i].hope += 5
			list[i].light += 5
		}
		return list
	}

	// 60. Call the Function
	brightnessValueList = increaseValueListValues(brightnessValueList)

	// 61. Print the New Array of Structs Values
	fmt.Println("Hope and Light in A Brighter Place Array of Structs:", brightnessValueList)

	// 62. Create a Function to Calculate the Total Hope and Light
	calcTotalHopeAndLight := func(list []Brightness) (int, int) {
		hope := 0
		light := 0
		for _, elem := range list {
			hope += elem.hope
			light += elem.light
		}
		return hope, light
	}

	// 63. Call the Function
	totalHope, totalLight := calcTotalHopeAndLight(list)

	// 64. Print the Total Hope and Light
	fmt.Printf("Total hope and light in A Brighter Place: %v, %v\n", totalHope, totalLight)

	// 65. Create a Function to Calculate the Average Hope and Light
	calcAvgHopeAndLight := func(list []Brightness) (float64, float64) {
		hopeSum := 0
		lightSum := 0
		for _, elem := range list {
			hopeSum += elem.hope
			lightSum += elem.light
		}
		hopeAvg := float64(hopeSum) / float64(len(list))
		lightAvg := float64(lightSum) / float64(len(list))
		return hopeAvg, lightAvg
	}

	// 66. Call the Function
	hopeAvg, lightAvg := calcAvgHopeAndLight(list)

	// 67. Print the Average Hope and Light
	fmt.Printf("Average hope and light in A Brighter Place: %.2f, %.2f\n", hopeAvg, lightAvg)

	// 68. Create a Struct to Store Values and Functions
	type BrightnessFunc struct {
		hope  int
		light int
		increase func()
	}

	// 69. Create an Instance of the Struct
	brighterPlaceFunc := BrightnessFunc{
		hope:  40,
		light: 35,
		increase: func() {
			brighterPlaceFunc.hope += 5
			brighterPlaceFunc.light += 5
		},
	}

	// 70. Print the Struct
	fmt.Println("Hope and Light in A Brighter Place Func:", brighterPlaceFunc)

	// 71. Call the Function
	brighterPlaceFunc.increase()

	// 72. Print the New Values
	fmt.Println("Hope and Light in A Brighter Place Func:", brighterPlaceFunc)

	// 73. Create a Function to Join Hope and Light
	joinHopeAndLight := func(hope int, light int) string {
		return fmt.Sprintf("Hope %v and Light %v", hope, light)
	}

	// 74. Call the Function
	hopeLightMsg := joinHopeAndLight(hopeValue, lightValue)

	// 75. Print the Message
	fmt.Println("We join together in ", hopeLightMsg)

	// 76. Create a Function to Compare Hope and Light
	compareHopeAndLight := func(hope int, light int) int {
		if hope > light {
			return 1
		} else if hope < light {
			return -1
		}
		return 0
	}

	// 77. Call the Function
	hopeLightDiff := compareHopeAndLight(hopeValue, lightValue)

	// 78. Print the Result
	if hopeLightDiff == 1 {
		fmt.Println("Hope is greater than light in A Brighter Place.")
	} else if hopeLightDiff == -1 {
		fmt.Println("Light is greater than hope in A Brighter Place.")
	} else {
		fmt.Println("Hope and light are equal