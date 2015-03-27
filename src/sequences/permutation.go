package main

import "strings"
import "fmt"

//put an element at every index in a string in a list of strings
func concat(n string, seq []string) []string{
    var finalSeq []string

    //for every index of seq, create a new string by inserting n
    for i:= 0; i<len(seq); i++ {
        items := strings.Split(seq[i], "")
        temp := append(items, "dummy")
        for j := 0; j < len(temp); j++{
            copy(temp[j+1:], temp[j:])
            temp[j] = n
            finalSeq = append(finalSeq, strings.Join(temp, ""))
            //temp is dirty, remake it 
            temp = append(items, "dummy")
        }
    }
    return finalSeq
}

func permutation(sequence []string) []string{
    //perm(n) = n.concat(perm(n-1))

    //base case
    if(len(sequence) == 1){
        return sequence
    }

    var finalSeq []string
    var n string
    n = sequence[0]
    subPermutation := permutation(sequence[1:])
    finalSeq = concat(n, subPermutation)
    return finalSeq
}

func main(){
    items := []string{"ab", "ba"}
    sequence := []string{"a", "b", "c", "d"}
    fmt.Printf("%q\n", concat("d", items))
    fmt.Printf("%q\n", permutation(sequence))
}
