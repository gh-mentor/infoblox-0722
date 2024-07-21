/*
This application uses channels to create a pipeline of goroutines.
The pipeline consists of two goroutines: a generator and a transformer.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Create a goroutine 'generator' 
Arguments:
- 'numbers': target channel of the generator.
- 'min': minimum value of the generated numbers.
- 'max': maximum value of the generated numbers.
Details:
- Generates integer values between `min` and `max` and sends them to the `numbers` channel.
- It closes the channel when it is done.
- Produces numbers at an irregular pace by adding a random delay between 50 and 200 milliseconds.
*/
func generator(numbers chan int, min, max int) {
	for i := 0; i < 10; i++ {
		numbers <- rand.Intn(max-min+1) + min
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(150)+50))
	}
	close(numbers)
}

/*
Create a goroutine 'transformer'
Arguments:
- 'numbers': source channel of the transformer.
- 'cubes': target channel of the transformer.
Details:
- Receives numbers from the `numbers` channel.
- Calculates the cube of each number
- Sends the result to the `cubes` channel.
- It continues this process until the `numbers` channel is closed, and then it closes the `cubes` channel.
*/
func transformer(numbers, cubes chan int) {
	for num := range numbers {
		cubes <- num * num * num
	}
	close(cubes)
}


// Create a 'main' function as the entry point of the program.
// Details:
// It creates two channels, `numbers` and `cubes`,
// and starts two goroutines: `generator` and `transformer`.
// It then receives values from the `cubes` channel and prints them.
func main() {
	numbers := make(chan int)
	cubes := make(chan int)

	go generator(numbers, 1, 10)
	go transformer(numbers, cubes)

	for cube := range cubes {
		fmt.Println(cube)
	}
}