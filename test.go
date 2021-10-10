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
	sCount, k := -1, -1
	w, _ := os.Create("test.csv")
	defer w.Close()
	n := len(d[0])
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if d[0][i][0] == '*' {
			d[0][i] = d[0][i][1:]
			sCount = i
		} else if d[0][i][0] == '?' {
			d[0][i] = d[0][i][1:]
			k = i
		}
		printCSV(&i, &n, w, &d[0][i])
	}
	st, sm := "", make(map[string]int)
	if len(d) < 2 {
		panic("too small like something something")
	}
	for y := 1; y < len(d); y++ {
		if k != -1 && len(d[y][k]) > 0 {
			for x := 0; x < n; x++ {
				var at *string
				at = &d[y][x]
				c := len(*at)
				if x == sCount {
					if c > 0 {
						st = *at
					}
					sm[st]++
					str := fmt.Sprintf(st, sm[st])
					printCSV(&x, &n, w, &str)
					continue
				} else if c == 0 && y > 1 {
					at = &d[arr[x]][x]
				} else {
					arr[x] = y
				}
				printCSV(&x, &n, w, at)
			}
		}
	}
}

func printCSV(i, size *int, w *os.File, str *string) {
	if *i+1 == *size {
		w.WriteString(*str + "\n")
	} else {
		w.WriteString(*str + ",")
	}
}
