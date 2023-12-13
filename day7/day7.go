package day7

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Bid           int
	Cards         string
	WildCardsCoef int
	HandMaximums  [2]int
}

// Card power mapping with Jacks (normal)
var pow_map = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

// Card power mapping with Jokers (wild) & no Jacks
var pow_map_j = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func Solution(input string) (solution1, solution2 int) {
	if input == "" {
		input = strings.Trim(TEST_INPUT, "\r\n")
	}

	// Solution 1
	/*
		Representation combinations as min-max card counts:
			five of a kind:  [5,0]
			four of a kind:  [4,1]
			full house:      [3,2]
			three of a kind: [3,1]
			two of a kind:   [2,1]
			one of a kind:   [1,1]
	*/

	hands := strings.Split(input, "\n")
	var interm_strengths [7][]Hand
	for _, hand_s := range hands {
		bid, _ := strconv.Atoi(strings.Split(hand_s, " ")[1])
		cards := strings.Split(hand_s, " ")[0]
		wild := strings.Count(cards, "J")
		hand := Hand{Cards: cards, Bid: bid, WildCardsCoef: wild}

		// consider max and min cardz to determine combination cases (full house)
		for ki := 0; ki < len(cards); ki++ { // max 4 iterations
			cc := strings.Count(cards, string(cards[ki]))
			if hand.HandMaximums[0] < cc {
				hand.HandMaximums[1] = hand.HandMaximums[0]
				hand.HandMaximums[0] = cc

				// Cut out registered and jump back
				cards = strings.ReplaceAll(cards, string(cards[ki]), "")
				ki--
			} else if hand.HandMaximums[1] < cc {
				hand.HandMaximums[1] = cc

				// Cut out registered  and jump back
				cards = strings.ReplaceAll(cards, string(cards[ki]), "")
				ki--
			}
		}

		switch hand.HandMaximums[0] { // Maximum desides the type of combination
		case 5: // five of a kind
			interm_strengths[0] = append(interm_strengths[0], hand)
		case 4: // four of a kind
			interm_strengths[1] = append(interm_strengths[1], hand)
		case 3: // full house or three of a kind
			if hand.HandMaximums[1] == 2 { // fullhouse (max 3, min 2)
				interm_strengths[2] = append(interm_strengths[2], hand)
			} else { // three of a kind
				interm_strengths[3] = append(interm_strengths[3], hand)
			}
		case 2: // two pair or pair
			if hand.HandMaximums[1] == 2 { // two pair
				interm_strengths[4] = append(interm_strengths[4], hand)
			} else { // pair
				interm_strengths[5] = append(interm_strengths[5], hand)
			}
		case 1: // high card
			interm_strengths[6] = append(interm_strengths[6], hand)
		}
	}

	var rhs []Hand
	for i := len(interm_strengths) - 1; i >= 0; i-- { // reversed sequence for ease of summing up later results of rank * bid
		strength := interm_strengths[i]
		sort.Slice(strength,
			func(x, y int) bool {
				for i := 0; i < len(strength[x].Cards); i++ { // compare every char until diff
					if pow_map[rune(strength[x].Cards[i])] < pow_map[rune(strength[y].Cards[i])] {
						return true
					} else if pow_map[rune(strength[x].Cards[i])] > pow_map[rune(strength[y].Cards[i])] {
						return false
					}
				}
				return false
			},
		)

		rhs = append(rhs, strength...) // merge sequantially
	}

	for i, hand := range rhs {
		solution1 += hand.Bid * (i + 1)
	}

	// Solution 2
	hands = strings.Split(input, "\n")
	interm_strengths = [7][]Hand{}
	for _, hand_s := range hands {
		bid, _ := strconv.Atoi(strings.Split(hand_s, " ")[1])
		cards := strings.Split(hand_s, " ")[0]
		ctrunc := strings.ReplaceAll(cards, "J", "") // cut out wild cards
		wild := strings.Count(cards, "J")            // wild card impact value
		hand := Hand{Cards: cards, Bid: bid, WildCardsCoef: wild}

		// consider max and min cardz to determine combination cases (full house)
		for ki := 0; ki < len(ctrunc); ki++ { // max 4 iterations
			cc := strings.Count(cards, string(ctrunc[ki]))
			if hand.HandMaximums[0] < cc {
				hand.HandMaximums[1] = hand.HandMaximums[0]
				hand.HandMaximums[0] = cc

				// Cut out registered  and jump back
				ctrunc = strings.ReplaceAll(ctrunc, string(ctrunc[ki]), "")
				ki--
			} else if hand.HandMaximums[1] < cc {
				hand.HandMaximums[1] = cc

				// Cut out registered  and jump back
				ctrunc = strings.ReplaceAll(ctrunc, string(ctrunc[ki]), "")
				ki--
			}
		}

		// Wild cards add to maximum for best hand combination
		switch hand.HandMaximums[0] + hand.WildCardsCoef { // Maximum + wild card impact desides the type of combination
		case 5: // five of a kind
			interm_strengths[0] = append(interm_strengths[0], hand)
		case 4: // four of a kind
			interm_strengths[1] = append(interm_strengths[1], hand)
		case 3: // full house or three of a kind
			if hand.HandMaximums[1] == 2 { // fullhouse (max 3, min 2)
				interm_strengths[2] = append(interm_strengths[2], hand)
			} else { // three of a kind
				interm_strengths[3] = append(interm_strengths[3], hand)
			}
		case 2: // two pair or pair
			if hand.HandMaximums[1] == 2 { // two pair
				interm_strengths[4] = append(interm_strengths[4], hand)
			} else { // pair
				interm_strengths[5] = append(interm_strengths[5], hand)
			}
		case 1: // high card
			interm_strengths[6] = append(interm_strengths[6], hand)
		}
	}

	rhs = []Hand{}
	for i := len(interm_strengths) - 1; i >= 0; i-- { // reversed sequence for ease of summing up later results of rank * bid
		strength := interm_strengths[i]
		sort.Slice(strength,
			func(x, y int) bool {
				for i := 0; i < len(strength[x].Cards); i++ { // compare every char until diff
					if pow_map_j[rune(strength[x].Cards[i])] < pow_map_j[rune(strength[y].Cards[i])] {
						return true
					} else if pow_map_j[rune(strength[x].Cards[i])] > pow_map_j[rune(strength[y].Cards[i])] {
						return false
					}
				}
				return false
			},
		)

		rhs = append(rhs, strength...) // merge sequantially
	}

	for i, hand := range rhs {
		solution2 += hand.Bid * (i + 1)
	}

	return
}

//go:embed test_data.txt
var TEST_INPUT string
