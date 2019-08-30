---
title: "Coding Questions in Golang"
excerpt: "Programming trivia in Golang"
header:
  teaser: assets/images/coding_questions_in_golang_01.jpg
feature_row:
  - image_path: /assets/images/gophercises_punching.jpg
  - image_path: /assets/images/gophercises_lifting.jpg
---

{% include feature_row %}

{% include toc %}

In this post, we will answer several programming challenge questions in Golang. These questions are often asked as part of interviews. Although there may be built-in methods to solve these problems, we shall only use basic data structures to solve each of them.

We prioritise recursive, concurrent, and clean codes. The solution codes are pretty self explanatory.

## Questions

1. [Find pair of numbers, with a given sum, from two arrays](#findPair)
2. [Delete single node, given a single pointer, in a singly linked list](#deleteSinglyLinkedNode)
3. [K-Means Clustering algorithm](#kmeans)
4. [Given two sorted lists, merge them into a single sorted list](#mergeSortedList)
5. [Print fibonanci numbers](#fibonanci)
6. [Reverse a string](#reverseString)
7. [Match brackets in a string](#matchBrackets)
8. [Range sum query](#rangeSum)
9. [Longest valid parentheses substring](#longestSubstring)
10. [Test driven development of Set](#tddSet)
11. [Sorted word count](#sortedWordCount)
12. [Fuzzy String Match](#fuzzyStringMatch)
13. [Breadth First Search & Depth First Search](#bfs&dfs)

{: .notice--success}
More programming challenge questions in Golang will be added as time permits. Let me know if there is any particular problem you would like to have solved here.

## Repository
The [repository](https://github.com/Adaickalavan/Coding-Questions-in-Golang) contains the Golang solution codes.

## Solutions

1. <a name="findPair"></a> Find every pair of numbers (inclusive of duplicates) from two integer arrays (i.e. one number from each array), whose sum equals a given number. A hash map is used to store previously seen numbers, for quick O(1) retrieval later.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/findPair).

    ```go
    package main

    import (
      "fmt"
    )

    //Retrieve pairs of numbers from two arrays which satisfies a given sum
    func findPair(arr1 []int, arr2 []int, n int) [][]int {
      hm := make(map[int]int)
      output := [][]int{}

      for _, val := range arr1 {
        hm[val] = hm[val] + 1
      }

      for _, val := range arr2 {
        temp := n - val
        if count, ok := hm[temp]; ok {
          for jj := count; jj > 0; jj-- {
            output = append(output, []int{temp, val})
          }
        }
      }

      return output
    }

    func main() {

      arr1 := []int{-1, -2, 4, 4, -2, -6, 5, -7}
      arr2 := []int{0, 6, 3, 4, 0}
      output := findPair(arr1, arr2, 4)
      fmt.Print(output)

    }

    ```

    Expected output:

    ```go
    [[4 0] [4 0] [-2 6] [-2 6] [4 0] [4 0]]
    ```

1. <a name="deleteSinglyLinkedNode"></a> Delete a single node, given only the pointer to that node, in a singly-linked list. The desired node is deleted by copying the contents (i.e., value and pointer) of the next node into the current node.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/deleteSinglyLinkedNode).

    {: .notice--warning}
    This solution does not work for deleting the last node in a singly linked list.

    ```go
    package main

    import (
      "fmt"
    )

    //Node contains value and pointer to next node
    type node struct {
      val      int
      nextNode *node
    }

    //Deletes node given by pointer
    func delete(ptr *node) {
      ptrNext := ptr.nextNode
      ptr.val = ptrNext.val
      ptr.nextNode = ptrNext.nextNode
    }

    func main() {

      sixthNode := &node{val: 6, nextNode: nil}
      fifthNode := &node{val: 5, nextNode: sixthNode}
      fourthNode := &node{val: 4, nextNode: fifthNode}
      thirdNode := &node{val: 3, nextNode: fourthNode}
      secondNode := &node{val: 2, nextNode: thirdNode}
      firstNode := &node{val: 1, nextNode: secondNode}

      delete(fourthNode)

      for ii := firstNode; ii != nil; ii = ii.nextNode {
        fmt.Println(ii)
      }
    }
    ```

    Expected output:

    ```go
    &{1 0xc04204a200}
    &{2 0xc04204a1f0}
    &{3 0xc04204a1e0}
    &{5 0xc04204a1c0}
    &{6 <nil>}
    ```

1. <a name="kmeans"></a> Given a list of data points and a list of centroids, perform `K-Means Clustering` to return a new set of updated centroids. Goroutines are used to parallelize the code execution.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/kmeans).

1. <a name="mergeSortedList"></a> Given two sorted lists, merge them into a single sorted list.

   Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/mergeSortedList).

    ```go
    package main

    import (
      "fmt"
    )

    //Recursively merge two ordered list
    func arrange(arr1 []int, arr2 []int, output *[]int) {
      switch {
      case len(arr1) == 0 && len(arr2) == 0:
        return
      case len(arr1) == 0 && len(arr2) > 0:
        *output = append(*output, arr2...)
        return
      case len(arr2) == 0 && len(arr1) > 0:
        *output = append(*output, arr1...)
        return
      default:
        if arr1[0] <= arr2[0] {
          *output = append(*output, arr1[0])
          arrange(arr1[1:], arr2, output)
        } else {
          *output = append(*output, arr2[0])
          arrange(arr1, arr2[1:], output)
        }
      }
    }

    func main() {
      arr1 := []int{1, 4, 7}
      arr2 := []int{3, 4, 5}
      output := []int{}
      arrange(arr1, arr2, &output)
      fmt.Println(output)
    }
    ```
    Expected output:
    ```go
    [1 3 4 4 5 7]
    ```

1. <a name="fibonanci"></a> Print Fibonanci numbers using closures in Golang.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/fibonanci).

    ```go
    package main

    import (
      "fmt"
    )

    //Closure returns a fibonanci function
    func evenFib() func() float64 {
      x := float64(0)
      y := float64(1)
      return func() float64 {
        x, y = y, x+y
        return y
      }
    }

    func main() {
      f := evenFib()
      for ii := 0; ii < 5; ii++ {
        result := f()
        fmt.Println(result)
      }
    }
    ```
    Expected output:
    ```go
    1
    2
    3
    5
    8
    ```
1. <a name="reverseString"></a> Reverse a given string. The string is reversed in parallel using goroutines and channels.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/reverseString).

    ```go
    package main

    import "fmt"

    //Reverse a string in parallel
    func reverser(arr1 []rune, length int, c0 chan int) {
      c1 := make(chan int)
      switch {
      case length == 0:
        c0 <- 1
        return
      case length == 1:
        c0 <- 1
        return
      case length >= 2:
        go reverser(arr1[1:length-1], length-2, c1)
        arr1[0], arr1[length-1] = arr1[length-1], arr1[0]
        <-c1
        c0 <- 1
      }
    }

    func main() {

      str := "abcdefgh"
      str1 := []rune(str)

      ch := make(chan int)

      go reverser(str1, len(str1), ch)
      <-ch

      fmt.Print(string(str1))
    }
    ```
    Expected output:
    ```text
    hgfedcba
    ```

1. <a name="matchBrackets"></a> Given a string containing just the characters '(', ')', '[', ']', '{', '}', determine if the input string is valid. An input string is valid if :
    + Open brackets must be closed by the same type of brackets.
    + Open brackets must be closed in the correct order.
  
    Note that an empty string is also considered valid. String below are considered valid strings.
    ```text
    ""
    "()"
    "()[]{}"
    "{[]}"
    ```
    Strings below are considered as invalid strings.
    ```text
    "["
    "}{"
    "{()"
    "(]"
    "([)]"
    ```

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/matchBrackets).

    ```go
    package main

    import (
      "fmt"
    )

    func main() {
      strList := []string{"", "()", "()[]{}", "{[]}", "[", "}{", "{()", "(]", "([)]"}
      for _, str := range strList {
        input := []rune(str)
        valid := matchBrackets(&[]rune{}, input)
        fmt.Println("The string", str, "is", valid)
      }
    }

    //Map closing brackets to corresponding opening brackets
    var bracketMap = map[rune]rune{
      rune(')'): rune('('),
      rune('}'): rune('{'),
      rune(']'): rune('['),
    }

    //Recursively compare brackets in a string
    func matchBrackets(stack *[]rune, input []rune) bool {
      switch {
      case len(*stack) == 0 && len(input) == 0:
        return true
      case len(*stack) != 0 && len(input) == 0:
        return false
      case len(*stack) == 0 && len(input) != 0:
        switch _, ok := bracketMap[input[0]]; ok {
        case true: //closing bracket
          return false
        default: //opening bracket
          *stack = append(*stack, input[0])
          return matchBrackets(stack, input[1:])
        }
      default: //len(*stack) != 0 && len(input) != 0
        switch elem, ok := bracketMap[input[0]]; ok {
        case true: //closing bracket
          if elem == (*stack)[len(*stack)-1] {
            *stack = (*stack)[:len(*stack)-1]
            return matchBrackets(stack, input[1:])
          }
          return false
        default: //opening bracket
          *stack = append(*stack, input[0])
          return matchBrackets(stack, input[1:])
        }
      }
    }
    ```
    Expected output:
    ```text
    The string  is true
    The string () is true
    The string ()[]{} is true
    The string {[]} is true
    The string [ is false
    The string }{ is false
    The string {() is false
    The string (] is false
    The string ([)] is false
    ```

1. <a name="rangeSum"></a> For an array $A$, its range sum $S$ from $i$ to $j$ is the sum of elements between indices $i$ and $j$, inclusive, where $i≤j$.

    $$
    S(i,j)=\sum_{t=i}^jA\left[t\right]
    $$

    Implement a struct with a receiver method $S(i,j)$. The struct will only be instantiated once, whereas the receiver method may be called several times during runtime. The optimization goal would be to minimize
    + usage of memory,
    + execution time of range sum receiver method $S$, and
    + execution time of struct instantiation.

    Several test cases are shown below.
    ```text
    A = [-2, 0, 3, -5, 2, -1]
    S(0, 2) //answer = 1
    S(2, 5) //answer = -1
    S(0, 5) //answer = -3
    ```

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/rangeSum).

1. <a name="longestSubstring"></a> Given a string containing just characters '(' and ')', return the start and end indexes of the longest valid parentheses substring. Several test cases are shown below.
    ```text
    Input: "" // answer: startIndex = 0, endIndex = 0
    Input: "(()" // answer: startIndex = 1, endIndex = 2
    Input: ")()())" // answer: startIndex = 1, endIndex = 4
    Input: "())((())" // answer: startIndex = 4, endIndex = 7
    Input: "())(()" // answer: startIndex = 0, endIndex = 1
    ```

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/longestSubstring).

    ```go
    package main

    import (
      "fmt"
    )

    //Returns the longest valid parentheses substring
    func longestSubstr(arr []rune) (int, int) {
      start := 0
      end := 0
      longestStart := 0
      longestEnd := 0
      stack := []int{-1}

      if len(arr) <= 1 {
        return longestStart, longestEnd
      }

      for ii, val := range arr {
        switch val {
        case rune('('):
          stack = append(stack, ii)
        default: //case ')':
          stack = stack[:len(stack)-1] //pop opening bracket
          if len(stack) != 0 {
            start = stack[len(stack)-1]
            end = ii
            if end-start > longestEnd-longestStart {
              longestStart = start
              longestEnd = end
            }
          } else { //empty stack
            stack = append(stack, ii)
          }
        }
      }
      return longestStart + 1, longestEnd
    }

    func main() {
      strList := []string{"", "(()", ")()())", "())((())", "())(()"}
      for _, str := range strList {
        input := []rune(str)
        start, end := longestSubstr(input)
        fmt.Println("Longest valid substring of", str, "is", string(input[start:end+1]), ", from ", start, "to ", end)
      }
    }
    ```
    Expected output:
    ```text
    Longest valid substring of  is   , from  0 to  0
    Longest valid substring of (() is () , from  1 to  2
    Longest valid substring of )()()) is ()() , from  1 to  4
    Longest valid substring of ())((()) is (()) , from  4 to  7
    Longest valid substring of ())(() is () , from  0 to  1
    ```

1. <a name="tddSet"></a> Implement a Set in Golang using test driven development method. The Set should implement `Add`,`Remove`,`Contains`,`Union`, and `Intersection`, functions.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/tddSet).

    ```go
    package set

    import (
      "fmt"
    )

    //Set implements the set
    type Set map[interface{}]struct{}

    //NewSet returns a new set
    func NewSet() *Set {
      s := Set{}
      return &s
    }

    //Print prints the keys within the set
    func (s *Set) Print() {
      fmt.Print("Contents of Set: ")
      for key := range *s {
        fmt.Print(key, ", ")
      }
    }

    //Add adds key to the set
    func (s *Set) Add(key interface{}) {
      (*s)[key] = struct{}{}
    }

    //Remove removes key from the set
    func (s *Set) Remove(key interface{}) {
      delete((*s), key)
    }

    //Contains checks whether a key exists in set
    func (s *Set) Contains(key interface{}) bool {
      _, ok := (*s)[key]
      return ok
    }

    //Union joins two sets
    func (s *Set) Union(other *Set) *Set {
      newSet := NewSet()
      for key := range *s {
        (*newSet)[key] = struct{}{}
      }
      for key := range *other {
        (*newSet)[key] = struct{}{}
      }
      return newSet
    }

    //Intersection returns intersection between two sets
    func (s *Set) Intersection(other *Set) *Set {
      newSet := NewSet()
      for key := range *s {
        if _, ok := (*other)[key]; ok {
          (*newSet)[key] = struct{}{}
        }
      }
      return newSet
    }
    ```

1. <a name="sortedWordCount"></a> Given an input text, produce a table of words and the number of times those words have occurred in the input text. Sort the words first by descending order of cardinality, then by alphabetical order. Other constraints:
   + Split words on whitespaces (newlines, spaces, tabs)
   + Remove any non-letter characters. In regex terms, remove anything not of the class [a-zA-Z]
   + Convert words to lowercase
   + Only list the first 10 most frequently occurring words
   + The text can include punctuation, non-letter characters, and mixed case
  
    Output should be formatted as: `<word><space><count><newline>`

    Example input text:
    ```text
    The quick brown fox jumped over the lazy dog's bowl.
    The dog was angry with the fox for considering him lazy.
    ```
    Expected output:
    ```text
    the 4
    fox 2
    lazy 2
    angry 1
    bowl 1
    brown 1
    considering 1
    dog 1
    dogs 1
    for 1
    ```

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/sortedWordCount).

    ```go
    package main

    import (
      "fmt"
      "log"
      "regexp"
      "sort"
      "strings"
    )

    type byCount []wordCount

    func (words byCount) Len() int {
      return len(words)
    }

    func (words byCount) Swap(i int, j int) {
      words[i], words[j] = words[j], words[i]
    }

    // Less sorts words by count, in descending order followed by alphabetical order
    func (words byCount) Less(i int, j int) bool {
      switch {
      case words[i].count > words[j].count:
        return true
      case words[i].count == words[j].count && words[i].word < words[j].word:
        return true
      default:
        return false
      }
    }

    type wordCount struct {
      word  string
      count int
    }

    func sorter(words []string) {
      m := make(map[string]int)
      var list []wordCount

      for _, word := range words {
        m[word] = m[word] + 1
      }

      for key, val := range m {
        list = append(list, wordCount{word: key, count: val})
      }

      sort.Sort(byCount(list))

      // Print the 10 most frequent words
      for _, elem := range list[:10] {
        fmt.Println(elem.word, elem.count)
      }

    }

    func clean(str string) []string {
      // Convert letters to lowercase and split string into words
      words := strings.Fields(strings.ToLower(str))

      // Remove all non letter characters
      reg, err := regexp.Compile("[^a-zA-Z0-9]+")
      if err != nil {
        log.Fatal(err)
      }
      for ii, elem := range words {
        words[ii] = reg.ReplaceAllString(elem, "")
      }

      return words
    }

    func main() {
      // Example string
      str := `The quick brown fox jumped over the lazy dog's bowl.
      The dog was angry with the fox for considering him lazy.`

      // Clean up the string
      words := clean(str)

      // Sort the words
      sorter(words)
    }
    ```

1. <a name="fuzzyStringMatch"></a> Given a DNA sequence (i.e., input string) and a DNA subsequence (i.e., search string), compute the percentage match of the DNA subsequence at every index location of the DNA sequence.
    Example input text:
    ```text
    Index:            012345678
    DNA sequence:    "ATGCATTGC"
    DNA subsequence: "CTG"
    ```
    Expected output:
    ```text
    0: 0.66
    1: 0
    2: 0
    3: 0.33
    4: 0.33
    5: 0.66
    6: 0
    7: 0
    8: 0.33
    ```
    Explanation of the output: At index location 0, 2/3 of the letters match between DNA sequence (i.e., ATG) and subsequence (i.e., CTG), hence percentage match is 66.6%. At index location 1, 0/3 of the letters match between DNA sequence (i.e., TGC) and subsequence (i.e., CTG), hence percentage match is 0%. At index location 7, 0/3 of the letters match between DNA sequence (i.e., GC-) and subsequence (i.e., CTG), hence percentage match is 0%. At index location 8, 1/3 of the letters match between DNA sequence (i.e., C--) and subsequence (i.e., CTG), hence percentage match is 33.3%.

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/fuzzyStringMatch).

    ```go
    package main

    import (
      "fmt"
    )

    func main() {
      // Input data
      sequence := "ATGCATTGC"
      subSequence := "CTG"
      sequenceLen := len(sequence)
      subSequenceLen := len(subSequence)

      var truncatedSequence string
      m := make(map[int]float64)
      var c = make(chan res)

      // Obtain substring and perform comparison
      for ii := 0; ii < sequenceLen; ii++ {
        if ii+subSequenceLen > sequenceLen {
          truncatedSequence = sequence[ii:sequenceLen]
        } else {
          truncatedSequence = sequence[ii : ii+subSequenceLen]
        }
        // Start goroutine for comparison
        go comp(truncatedSequence, subSequence, ii, c)
      }

      // Retrieve results from each comparison
      for ii := 0; ii < sequenceLen; ii++ {
        r := <-c
        m[r.key] = r.val
      }

      // Print results
      fmt.Println(m)
    }

    type res struct {
      key int
      val float64
    }

    // comp compares sequence and subSequence, letter-by-letter, returning the percentage match
    func comp(sequence string, subSequence string, key int, c chan res) {
      sequenceLen := len(sequence)
      subSequenceLen := len(subSequence)
      var count int
      for jj := 0; jj < sequenceLen; jj++ {
        if sequence[jj] == subSequence[jj] {
          count++
        }
      }
      // Return result
      c <- res{key: key, val: float64(count) / float64(subSequenceLen)}
    }
    ```

1. <a name="bfs&dfs"></a> Traverse a tree, in a breadth first and depth first manner, from the starting node till the ending node. The first node, namely `node1`, is the starting node. The node with a field value `end:true` marks the ending node. Return true if an end node is reached, else return false. Remember to avoid possible cycles in the tree.
  
    Example input tree:
    ```text
    Tree structure
                    node1
                      |
            ---------------------
            |                   |
          node2               node 3
            |                   |
        -----------             |
        |         |             |
      node4     node5         node 6
                  |
             -----------
             |         |
           node7     node8
    ```
    Expected output:
    ```text
    Breath first search = 1 2 3 4 5 6 7 8 true
    Depth first search = 1 2 4 5 7 8 true
    ```

    Link to solution [code](https://github.com/Adaickalavan/Coding-Questions-in-Golang/tree/master/breadthFirstSearch-DepthFirstSearch).  

    ```go
    package main

    import "fmt"

    var m = make(map[string]bool)

    type node struct {
      next []*node
      end  bool
      name string
    }

    type queue []*node

    // depth first search
    func dfs(curNode *node) bool {
      if _, ok := m[curNode.name]; ok {
        return false
      }

      fmt.Print(curNode.name, " ")

      if curNode.end {
        return true
      }

      for _, newNode := range curNode.next {
        if dfs(newNode) {
          return true
        }
      }

      return false
    }

    // breadth first search
    func bfs(q *queue) bool {
      if len(*q) == 0 {
        return false
      }

      curNode := (*q)[0]

      if _, ok := m[curNode.name]; ok {
        return false
      }

      fmt.Print(curNode.name, " ")

      if curNode.end {
        return true
      }

      *q = append((*q)[1:], curNode.next...)

      if bfs(q) {
        return true
      }

      return false
    }

    func main() {

      node1 := &node{end: false, name: "1"}
      node2 := &node{end: false, name: "2"}
      node3 := &node{end: false, name: "3"}
      node4 := &node{end: false, name: "4"}
      node5 := &node{end: false, name: "5"}
      node6 := &node{end: false, name: "6"}
      node7 := &node{end: false, name: "7"}
      node8 := &node{end: true, name: "8"}

      nexta := []*node{node2, node3}
      nextb := []*node{node4, node5}
      nextc := []*node{node6}
      nextd := []*node{node7, node8}

      node1.next = nexta
      node2.next = nextb
      node3.next = nextc
      node5.next = nextd

      // Breadth first search
      fmt.Print("Breath first search = ")
      n := queue([]*node{node1})
      tree := bfs(&n)
      fmt.Println(tree)

      // Reset map
      m = make(map[string]bool)

      // Depth first search
      fmt.Print("Depth first search = ")
      tree = dfs(node1)
      fmt.Println(tree)
    }
    ```