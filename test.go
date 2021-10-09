package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	readData()
}

func readStyle() {
	f, _ := os.Open("Certificate List - final.csv")
	r := csv.NewReader(f)
	d, _ := r.ReadAll()
	fmt.Println(d[0])
}

func readData() {
	f, err := os.Open("Certificate List - final.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	d, _ := r.ReadAll()
	s, k := -1, -1
	for i := 0; i < len(d); i++ {
		if d[0][i][0] == '*' {
			d[0][i] = d[0][i][1:]
			s = i
		} else if d[0][i][0] == '?' {
			d[0][i] = d[0][i][1:]
			k = i
		}
		println(d[0][i])
	}
	st, i := "", 1
	w, _ := os.Create("test.csv")
	defer w.Close()
	if len(d) < 2 {
		panic("too small like something something")
	}
	for y := 1; y < len(d); y++ {
		if k != -1 && len(d[y][k]) > 0 {
			for x := 0; x < len(d[y]); x++ {
				if x == s {
					if len(d[y][x]) > 0 {
						st = d[y][x]
						i = 1
					}
					w.WriteString(fmt.Sprintf(st+",", i))
					i++
				} else if x+1 == len(d[y]) {
					if len(d[y][x]) == 0 && y > 1 {
						d[y][x] = d[y-1][x]
					}
					w.WriteString(d[y][x] + "\n")
				} else {
					if len(d[y][x]) == 0 && y > 1 {
						d[y][x] = d[y-1][x]
					}
					w.WriteString(d[y][x] + ",")
				}
			}
		} else {
			for x := 0; x < len(d[y]); x++ {
				if x+1 == len(d[y]) {
					w.WriteString(d[y][x] + "\n")
				} else {
					w.WriteString(d[y][x] + ",")
				}
			}
		}
	}
}
