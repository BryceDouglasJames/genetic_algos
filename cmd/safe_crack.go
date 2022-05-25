package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const COMBINATION = "3875307572"

var trace = false
var attempt = "0000000000"

//If we want to generate comparator functions, we have to create an ordered interface...
type ordered interface {
	~string | ~int
}
type tuple[T ordered, U ordered] struct {
	First  T
	Second U
}

//we'll keep this just in case generics become too much
/*type tuple struct {
	First  string
	Second string
}*/

func Fitness_Function(code []string, attempt []string) int {
	var temp_score = 0
	zipped := Zip(code, attempt)
	if trace {
		fmt.Println("FITNESS FUNCTION")
		fmt.Println(zipped)
	}
	for i := range zipped {
		temp := zipped[i]
		if trace {
			fmt.Printf("Tuple 1: %+v\tTuple 2: %+v\n", temp.First, temp.Second)
		}
		if strings.Compare(temp.First, temp.Second) == 0 {
			if trace {
				fmt.Printf("\t\tFOUND A PAIR: %s %s\n", temp.First, temp.Second)
			}
			temp_score += 1
		}
	}
	return temp_score
}

func Naive_Hill_Climb() {
	var count = 0

	//generate a character slice out of combination
	combo_list := strings.Split(COMBINATION, "")

	//generate a character slice out of attempt
	best_attempt := strings.Split(attempt, "")

	//generate inital fitness grade
	best_fit := Fitness_Function(combo_list, best_attempt)

	for {

		//Break if combo found
		if strings.Join(best_attempt, "") == COMBINATION {
			break
		}

		//the slice type has an unsafe pointer index.... this is annoying because
		//we cant implicitly reassign next_attmempt. Pass-by-value is really funky sometimes.
		var next_attempt = []string{}
		next_attempt = append(next_attempt, best_attempt...)

		//We will be using the cyrpto package instead...
		//rand.Seed(time.Now().UnixNano())
		//next_index := rand.Intn(len(COMBINATION))

		//assign random index HARD CODED DOMAIN
		answer_len := big.NewInt(10)
		next_index, _ := rand.Int(rand.Reader, answer_len)

		//assign random value HARD CODED DOMAIN
		val_len := big.NewInt(11)
		rand_val, _ := rand.Int(rand.Reader, val_len)

		//mutate
		next_attempt[next_index.Int64()] = rand_val.String()

		//apply fitness function see if mutation is better
		next_grade := Fitness_Function(combo_list, next_attempt)

		//if the mutation is better fit, save it.
		if next_grade > best_fit {
			fmt.Printf("\nFOUND A BETTER FIT: %d is the last fit, %d is the new best. Found on try #%d!\n", best_fit, next_grade, count)
			best_attempt = next_attempt
			best_fit = next_grade
			fmt.Printf("best is %s\n", strings.Join(best_attempt, ""))

		}

		count += 1

		//It should not reach this... but just in case :)
		if count >= 100000 {
			fmt.Println("Okay something may be wrong here...")
			fmt.Printf("\nThe most tuples matched was %d and the best offspring was %s\n", best_fit, strings.Join(best_attempt, ""))
			break
		}
	}

	fmt.Printf("\nThis is the best attempt: %s\nCracked in %d attempts\nBest fit was %d\n", best_attempt, count, best_fit)
}

//go generics are so cool!!!...kinda
//zip helper function. create tuple tuple object.
func Zip[T, U ordered](t_vals []T, u_vals []U) []tuple[T, U] {

	//make sure lists are symetric
	if len(t_vals) != len(u_vals) {
		panic("slices have different length")
	}

	//generate tuple slice
	tuples := make([]tuple[T, U], len(t_vals))
	for i := 0; i < len(t_vals); i++ {
		tuples[i] = tuple[T, U]{First: t_vals[i], Second: u_vals[i]}
	}

	//return tuple slice
	return tuples
}
