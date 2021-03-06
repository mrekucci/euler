// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package euler

// Problem15 is solution for finding how many routes are in a 20×20 grid if
// from the top left corner we are only able to move to the right and down.
func Problem15() int {
	// Solution base on combinatorics.
	// In how many different ways can we choose N out of 2N possible places if
	// the order does not matter? Solution is computed with help of Binomial
	// Coefficient: 40!/((40-20)!*20!); (combination without repetition).
	//
	//	const dim = 20
	//	return int(big.NewInt(0).Binomial(2*dim, dim).Int64())
	//

	// Second solution based on dynamic programming.
	// It's faster but memory less efficient then the previous one.
	const dim = 20
	// +1 'cause inside the grid is number of ways to get from previous to next point.
	var grid [dim + 1][dim + 1]int

	// From previous point to next right point there
	// is one way and so is to next down point.
	for i := 0; i < dim; i++ {
		grid[i][dim] = 1
		grid[dim][i] = 1
	}

	// We count the number of ways backwards, from bottomRight up to topLeft.
	for x := dim - 1; x >= 0; x-- {
		for y := dim - 1; y >= 0; y-- {
			grid[x][y] = grid[x+1][y] + grid[x][y+1]
		}
	}

	return grid[0][0]
}
