package bmak_tools

import (
	"encoding/json"
	"fmt"
	"github.com/CrimsonAIO/radix"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (bm *Bmak) SolveChallenge() (string, string, string) {
	bm.MnTs = strconv.Itoa(int(bm.StartTs))
	bm.MnMcLmt = 10
	bm.MnTout = 100
	bm.MnStout = 1000
	bm.MnCt = 1
	bm.MnCd = 10000
	bm.StartTs += int64(randIntRange(100, 200))
	bm.MnLc = make([]string, 1)
	bm.MnLd = make([]int, 1)
	bm.MnAl = make([]string, 1)
	bm.MnIl = make([]int, 1)
	bm.MnR = make(map[string]string)
	bm.MnTcl = make([]int, 1)

	params := bm.get_mn_params_from_abck()

	if len(params) == 0 {
		return "", "", ""
	}

	new_params := bm.mn_get_new_challenge_params(params)
	fmt.Println(new_params)

	if new_params != nil {
		bm.mn_update_challenge_details(new_params)
		if bm.MnSen != 0 {
			bm.MnState = 1
			bm.MnMcIndx = 0
			bm.MnAl = []string{}
			bm.MnIl = []int{}
			bm.MnTcl = []int{}
			bm.MnLg = []string{}
			bm.Mnrts = int(bm.getCfDate())
			bm.MnRt = bm.Mnrts - int(bm.StartTs)
		}
	}

	bm.mn_w()

	C := bm.mn_get_current_challenges()
	B := ""
	x := ""
	M := ""

	if len(C) >= 2 && len(C[1]) != 0 {
		j := C[1]
		if len(bm.MnR[j]) != 0 {
			B = bm.MnR[j]
		}
	}
	if len(C) >= 3 && len(C[2]) != 0 {
		A := C[2]
		if len(bm.MnR[A]) != 0 {
			x = bm.MnR[A]
		}
	}
	if len(C) >= 4 && len(C[3]) != 0 {
		L := C[3]
		if len(bm.MnR[L]) != 0 {
			M = bm.MnR[L]
		}
	}

	return B, x, M
}

func (bm *Bmak) get_mn_params_from_abck() [][]interface{} {
	var t [][]interface{}

	var a = bm.CookieValue

	if true {
		//unescapes "a" as it contains URL encoded values
		aUnescaped, _ := url.QueryUnescape(a)
		//splits a by "~" which is probably to get the challenge indices or something
		var e = strings.Split(aUnescaped, "~")

		if len(e) >= 5 {
			n := e[0]
			//o holds the challenge string at the end of the cookie value
			o := e[4]
			//m now holds o seperated, akamai challenges are determined by the double || at the end
			m := strings.Split(o, "||")
			//check if challenges are existant by checking m array size
			if len(m) > 0 {
				for r := 0; r < len(m); r++ {
					//i holds the first challenge value
					i := m[r]
					c := strings.Split(i, "-")
					//checks for correct length of c
					if len(c) > 5 {
						//sets the indices of the t slice, the 1st,3rd,4th,5th indices are all integers and the second
						b, _ := strconv.Atoi(c[0])
						d := c[1]
						s, _ := strconv.Atoi(c[2])
						k, _ := strconv.Atoi(c[3])
						l, _ := strconv.Atoi(c[4])
						u := 1

						//checks if the length of c is bigger than 6 if yes then reassign u var,
						//in js it is c.length >= 6 && (u = parseInt(c[5])
						if len(c) >= 6 {
							u, _ = strconv.Atoi(c[5])
						}
						//using interface because its the only way to append the payload to an array since the types are mixed
						var v_ = []interface{}{b, n, d, s, k, l, u}

						// in js this is 2 == u ? t.splice(0, 0, _) : t.push(_)
						if 2 == u {
							t = append(t, []interface{}{})
							t = append([][]interface{}{v_}, t...)
						} else {
							t = append(t, v_)
						}

					}

				}
			}
		}
	}
	return t
}

func (bm *Bmak) mn_get_new_challenge_params(t [][]interface{}) []interface{} {
	var a int = -99999
	var e int = -99999
	var n int = -99999

	//checks if array is nil
	if t != nil {
		//makes iterator that mathces the length of the t array
		for o := 0; o < len(t); o++ {
			//assigns m to the first o-th of the t array, which is another array
			m := t[o]
			//checks if m is bigger than 0, so it always works only for the 1st element of the array as its always [[..], []]
			if len(m) > 0 {
				//we do type assertion as the 1st element of the inner array is always an int
				r := m[0].(int)
				//not sure what i is but its the same as the akamai script
				i := bm.MnAbck + fmt.Sprint(bm.StartTs) + fmt.Sprint(m[2])
				//we do type assertion as the 6th element of the inner array is always an int
				c := m[6].(int)

				//we do type assertion as we already know the value type
				var b = 0
				for b = 0; b < bm.MnLcl && (r == 1 && bm.MnLc[b] != i); b++ {
				}
				if bm.MnLcl == b {
					a = o
					if 2 == c {
						e = o
					}
					if 3 == c {
						n = o
					}
				}
			}
		}

		if n != -99999 && bm.Pstate {
			return t[n]
		} else {
			if e == -99999 || bm.Pstate {
				if a == -99999 || bm.Pstate {
					return nil
				} else {
					return t[a]
				}
			} else {
				return t[e]
			}
		}
	}
	return nil
}

func (bm *Bmak) mn_update_challenge_details(t []interface{}) {
	bm.MnSen = t[0].(int)
	bm.MnAbck = t[1].(string)
	bm.MnPsn = t[2].(string)
	bm.MnCd = t[3].(int)
	bm.MnTout = t[4].(int)
	bm.MnStout = t[5].(int)
	bm.MnCt = t[6].(int)
	bm.MnTs = strconv.Itoa(int(bm.StartTs))
	bm.MnCc = bm.MnAbck + strconv.Itoa(int(bm.StartTs)) + bm.MnPsn
}

func (bm *Bmak) mn_w() int {
	a := 0
	e := 0
	n := ""
	o := bm.getCfDate()
	m := bm.MnCd + bm.MnMcIndx
	for t := 0; t == 0; {
		n, _ = radix.ToString(rand.Float64(), 16)
		r := bm.MnCc + fmt.Sprint(m) + n
		i := bm.mn_s(r)
		if 0 == bm.bdm(i, m) {
			t = 1
			e = int(bm.getCfDate() - o)
			bm.MnAl = append(bm.MnAl, n)
			bm.MnTcl = append(bm.MnTcl, e)
			bm.MnIl = append(bm.MnIl, a)
			if bm.MnLg = append(bm.MnLg, bm.MnAbck, bm.MnTs, bm.MnPsn, bm.MnCc, fmt.Sprint(bm.MnCd), fmt.Sprint(m), fmt.Sprint(n), r, marshal(i), strconv.Itoa(bm.MnRt)); bm.MnMcIndx == 0 {
			} else if a += 1; a%1e3 == 0 && (bm.getCfDate()-o) > int64(bm.MnStout) {
				e := bm.getCfDate()
				bm.MnWt += int(e)
				return bm.mn_w()
			}
		}
	}
	bm.MnMcIndx += 1
	if bm.MnMcIndx < bm.MnMcLmt {
		bm.mn_w()
	} else {
		bm.MnMcIndx = 0
		bm.MnLc[bm.MnLcl] = bm.MnCc
		bm.MnLd[bm.MnLcl] = bm.MnCd
		bm.MnLcl += 1
		bm.MnState = 0
		bm.MnLg = append(bm.MnLg, strconv.Itoa(bm.MnWt))
		bm.MnLg = append(bm.MnLg, strconv.Itoa(int(bm.getCfDate())))
		bm.MnR[bm.MnAbck+bm.MnPsn] = bm.mn_pr()
		if bm.JsPost {
			bm.AjType = 8
			if bm.MnCt == 2 {
				bm.Dcs = 1
			}
		}
	}
	return bm.MnWt
}

func (bm *Bmak) mn_get_current_challenges() []string {
	t := bm.get_mn_params_from_abck()
	a := make([]string, 3)

	if t != nil {
		for e := 0; e < len(t); e++ {
			n := t[e]
			if len(n) > 0 {
				o := n[1].(string) + n[2].(string)
				m := n[6].(int)
				a[m] = o
			}
		}
	}
	return a
}

func randIntRange(min, max int) int {
	rand.Seed(time.Now().UnixNano() + time.Now().Unix())
	return min + rand.Intn(max-min+1)
}

func marshal(i interface{}) string {
	v, _ := json.Marshal(i)
	return string(v)
}
