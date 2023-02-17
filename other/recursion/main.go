package main

import (
	"fmt"
	"strconv"
)

var mapData = map[int][]string{
	1: {"a", "b", "c"},
	2: {"d", "e", "f"},
	// 3: {"g", "h", "i"},
}

func main() {
	outData := []string{}
	help("122", &outData, "")
	fmt.Println(outData)

	help2(mapData, "122", "")
	fmt.Println(res)
}

func help(put string, outData *[]string, tmp string) {
	if len([]rune(put)) == 0 {
		*outData = append(*outData, tmp)
		return
	}

	stepInt, _ := strconv.Atoi(put[:1])
	for _, val := range mapData[stepInt] {
		help(put[1:], outData, tmp+val)
	}
}

var res []string

func help2(mapList map[int][]string, start string, tmp string) {
	if len(start) == 0 {
		res = append(res, tmp)
		return
	}

	index, _ := strconv.Atoi(start[:1])
	for _, val := range mapList[index] {
		help2(mapList, start[1:], tmp+val)
	}
}
