package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(actual int) int {
	/*
		the actual minutes the lasagna has been in the oven as a parameter and
		returns how many minutes the lasagna still has to remain in the oven,
		based on the expected oven time in minutes from the previous task.
	*/

	return OvenTime - actual

}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers int) int {
	/*
		takes the number of layers you added to the lasagna as a parameter
		and returns how many minutes you spent preparing the lasagna,
		assuming each layer takes you 2 minutes to prepare.
	*/

	return 2 * layers

}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, mins_in_oven int) int {
	/*
		takes two parameters: the first parameter is the number of
		layers you added to the lasagna, and the second parameter
		is the number of minutes the lasagna has been in the oven.
		The function should return how many minutes in total you've
		worked on cooking the lasagna, which is the sum of the
		preparation time in minutes, and the time in minutes the
		lasagna has spent in the oven at the moment.
	*/

	return mins_in_oven + layers*2

}
