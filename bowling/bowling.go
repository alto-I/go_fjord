package bowling

import (
	"strconv"
)

func CalcScore(bins []string) int {
	// Xを10にして、全てを数値型に変換
	var intbins []int
	for _, v := range bins {
		if v == "X" {
			intbins = append(intbins, 10)
			intbins = append(intbins, 0)
		} else {
			i, _ := strconv.Atoi(v)
			intbins = append(intbins, i)
		}
	}

	// fmt.Println(intbins)

	// 成績をフレーム単位に整形
	var frames [][]int
	sliceSize := len(intbins)

	for i := 0; i < sliceSize; i += 2 {
		end := i + 2
		if sliceSize < end {
			end = sliceSize
		}
		frames = append(frames, intbins[i:end])
	}

	// fmt.Println(frames)

	// スコア計算
	sum := 0
	for i, v := range frames {
		if i >= 9 {
			for _, v := range v {
				sum += v
			}
		} else {
			if v[0] == 10 {
				if frames[i+1][0] == 10 {
					sum += v[0] + frames[i+1][0] + frames[i+2][0]
				} else {
					sum += v[0] + frames[i+1][0] + frames[i+1][1]
				}
			} else if v[0]+v[1] == 10 {
				sum += v[0] + v[1] + frames[i+1][0]
			} else {
				sum += v[0] + v[1]
			}
		}
	}

	return sum
}
