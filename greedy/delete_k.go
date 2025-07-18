package greedy

import (
	"bufio"
	"fmt"
	"os"
)

type DeleteK struct {
}

func NewDeleteK() *DeleteK {
	return &DeleteK{}
}

func (d *DeleteK) Delete() {
	reader := bufio.NewReader(os.Stdin)
	var (
		nums string
		k    int
	)
	fmt.Fscanln(reader, &nums)
	fmt.Fscanln(reader, &k)

	res := "0"
	for _, num := range nums {
		for k > 0 && num < int32(res[len(res)-1]) {
			k--
			res = res[0 : len(res)-1]
		}
		res = res + string(num)
	}
	for k > 0 {
		k--
		res = res[0 : len(res)-1]
	}
	k = 0
	for k < len(res) && res[k] == '0' {
		k++
	}
	if k == len(res) {
		fmt.Println(0)
	} else {
		res = res[k:len(res)]
		fmt.Println(res)
	}

}
