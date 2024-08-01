//Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
//
//
//Example 1:
//
//Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//Example 2:
//
//Input: strs = [""]
//Output: [[""]]
//Example 3:
//
//Input: strs = ["a"]
//Output: [["a"]]

package main

func group_anagram(strs []string) [][]string {
	store := make(map[string][]string)
	for _, s := range strs {
		hashKey := [26]byte{}
		for i := 0; i < len(s); i++ {
			hashKey[s[i]-'a']++
		}
		// convert the hash,key to string
		hashKeyStr := string(hashKey[:])
		store[hashKeyStr] = append(store[hashKeyStr], s)
	}
	results := make([][]string, 0)
	for _, v := range store {
		results = append(results, v)
	}
	return results
}
