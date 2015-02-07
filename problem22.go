// Copyright 2015 Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package euler

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Problem22 is solution for finding the total of all the name scores in the
// file names.txt. Score for every name is computed on alphabetical sorted list,
// where working out the alphabetical value for each name, multiply this value
// by its alphabetical position in the list. Alphabetical value of the name is
// defined as a sum of the numbers representing the order of the individual
// characters of the alphabet from which consists a particular name.
func Problem22() (int, error) {
	total := 0

	f, err := ioutil.ReadFile("data/names.txt")
	if err != nil {
		return 0, fmt.Errorf("read file error: %s", err)
	}

	// We may use csv reader instead of the code below, but below code is faster.
	//	names, err := csv.NewReader(bytes.NewReader(f)).Read()
	//	if err != nil {
	//		return 0, fmt.Errorf("parse error: %s", err)
	//	}
	//
	names := strings.Split(string(f), ",")
	for i, n := range names {
		names[i] = n[1 : len(n)-1] // Removes '"' character from beginning and the end.
	}

	sort.Strings(names)

	// We assume that all names are in uppercase, so we don't need conversion.
	for i, n := range names {
		score := len(n) // We have to add +1 for every letter in name, because 'A'-'A' = 0 not 1.
		for _, c := range n {
			score += int(c - 'A')
		}
		score *= (i + 1) // i+1 'cause 'names' index starts from 0 not 1.
		total += score
	}

	return total, nil
}
