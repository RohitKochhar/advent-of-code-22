package main

import "testing"

func TestPart1(t *testing.T) {
	// Use table-driven testing from the examples given
	testcases := []struct {
		Name     string
		Input    string
		Expected int
	}{
		{"TestCase1", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"TestCase2", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"TestCase3", "nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"TestCase4", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"TestCase5", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			if res := Part1(tc.Input); res != tc.Expected {
				t.Fatalf("Got %d, expected %d", res, tc.Expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	// Use table-driven testing from the examples given
	testcases := []struct {
		Name     string
		Input    string
		Expected int
	}{
		{"TestCase1", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"TestCase2", "bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"TestCase3", "nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"TestCase4", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"TestCase5", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			if res := Part2(tc.Input); res != tc.Expected {
				t.Fatalf("Got %d, expected %d", res, tc.Expected)
			}
		})
	}
}
