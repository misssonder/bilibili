package video

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var tr = map[string]int64{}
var s = []int64{11, 10, 3, 8, 4, 6}
var xor int64 = 177451812
var add int64 = 8728348608

func init() {
	tableByte := []byte(table)
	for i := 0; i < 58; i++ {
		tr[string(tableByte[i])] = int64(i)
	}
}

func BvIDToAID(bv string) int64 {
	var r int64
	arr := []rune(bv)

	for i := 0; i < 6; i++ {
		r += tr[string(arr[s[i]])] * int64(math.Pow(float64(58), float64(i)))
	}
	return (r - add) ^ xor
}

func AIDtoBvID(av int64) string {
	x := (av ^ xor) + add
	r := []string{"B", "V", "1", " ", " ", "4", " ", "1", " ", "7", " ", " "}
	for i := 0; i < 6; i++ {
		r[s[i]] = string(table[int(math.Floor(float64(x/int64(math.Pow(float64(58), float64(i))))))%58])
	}
	var result string
	for i := 0; i < 12; i++ {
		result += r[i]
	}
	return result
}

var bvIDRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`),
	regexp.MustCompile(`(video\/BV([a-z]|[0-9]|[A-Z])+\/)`),
	regexp.MustCompile(`BV([a-z]|[0-9]|[A-Z])+`),
}

var ssIDRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`),
	regexp.MustCompile(`(play\/ss([0-9])+)`),
	regexp.MustCompile(`([0-9])+`),
}

var epIDRegexpList = []*regexp.Regexp{
	regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`),
	regexp.MustCompile(`(play\/ep([0-9])+)`),
	regexp.MustCompile(`([0-9])+`),
}

func ExtractBvID(id string) (string, error) {
	if aid, err := strconv.ParseInt(id, 10, 64); err == nil {
		return AIDtoBvID(aid), nil
	}
	for _, re := range bvIDRegexpList {
		if isMatch := re.MatchString(id); isMatch {
			subs := re.FindStringSubmatch(id)
			id = subs[0]
		}
	}
	if len(id) < 12 {
		return "", fmt.Errorf("invalid characters in id")
	}
	return id, nil
}

func ExtractSSID(id string) (string, error) {
	for _, re := range ssIDRegexpList {
		if isMatch := re.MatchString(id); isMatch {
			subs := re.FindStringSubmatch(id)
			id = subs[0]
		}
	}
	return id, nil
}
func ExtractEpID(id string) (string, error) {
	for _, re := range epIDRegexpList {
		if isMatch := re.MatchString(id); isMatch {
			subs := re.FindStringSubmatch(id)
			id = subs[0]
		}
	}
	return id, nil
}

func IsSSID(id string) bool {
	return regexp.MustCompile(`(play\/ss([0-9])+)`).MatchString(id)
}

func IsEpID(id string) bool {
	return regexp.MustCompile(`(play\/ep([0-9])+)`).MatchString(id)
}
