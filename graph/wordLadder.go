// https://leetcode.com/problems/word-ladder/
package graph

import "fmt"

/*
hit

{
	*it
	h*t
	hi*
	*ot
	d*t
	do* -> [dot, dog]
	*ot -> [dot, hot]
	d*t
}
ait, bit, cit..... zit -> (26-1)words in list ->
hat, hbt, hct......hzt -> (26-1)words in list -> hot
hia, hib, hic......hiz -> (26-1)words in list


hot
1. dot, lot
2. x
3. x

dot
1. lot
2. x
3. dog
*/

type QueueElement struct {
	word        string
	stepCounter int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	patternDict := make(map[string][]string)

	for _, word := range wordList {
		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]

			if _, ok := patternDict[pattern]; !ok {
				patternDict[pattern] = []string{}
			}

			patternDict[pattern] = append(patternDict[pattern], word)
		}
	}

	queue := []QueueElement{{word: beginWord, stepCounter: 1}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		queueElement := queue[0]
		queue = queue[1:]
		word := queueElement.word

		if word == endWord {
			return queueElement.stepCounter
		}

		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]
			for _, word := range patternDict[pattern] {
				if _, ok := visited[word]; !ok {
					queue = append(queue, QueueElement{word: word, stepCounter: queueElement.stepCounter + 1})
					visited[word] = true
				}
			}
		}
	}

	return 0
}

func DoWordLadder() {
	count := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Printf("Going from %s -> %s: took %d iterations\n", "hit", "cog", count)
	count = ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	fmt.Printf("Going from %s -> %s: took %d iterations\n", "hit", "cog", count)
}
