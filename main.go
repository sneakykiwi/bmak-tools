package bmak_tools

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Bmak struct {
	Fpcf struct {
		Fpvalstr        string `json:"fpValstr"`
		Fpvalcalculated bool   `json:"fpValCalculated"`
		Rval            string `json:"rVal"`
		Rcfp            string `json:"rCFP"`
		Cache           struct {
		} `json:"cache"`
		Td      int      `json:"td"`
		Plugins []string `json:"PLUGINS"`
	}
	Ver               float64       `json:"ver"`
	KeCntLmt          int           `json:"ke_cnt_lmt"`
	MmeCntLmt         int           `json:"mme_cnt_lmt"`
	MduceCntLmt       int           `json:"mduce_cnt_lmt"`
	PmeCntLmt         int           `json:"pme_cnt_lmt"`
	PduceCntLmt       int           `json:"pduce_cnt_lmt"`
	TmeCntLmt         int           `json:"tme_cnt_lmt"`
	TduceCntLmt       int           `json:"tduce_cnt_lmt"`
	DoeCntLmt         int           `json:"doe_cnt_lmt"`
	DmeCntLmt         int           `json:"dme_cnt_lmt"`
	VcCntLmt          int           `json:"vc_cnt_lmt"`
	DoaThrottle       int           `json:"doa_throttle"`
	DmaThrottle       int           `json:"dma_throttle"`
	SessionID         string        `json:"session_id"`
	JsPost            bool          `json:"js_post"`
	Loc               string        `json:"loc"`
	CfURL             string        `json:"cf_url"`
	ParamsURL         string        `json:"params_url"`
	Auth              string        `json:"auth"`
	APIPublicKey      string        `json:"api_public_key"`
	AjLmtDoact        int           `json:"aj_lmt_doact"`
	AjLmtDmact        int           `json:"aj_lmt_dmact"`
	AjLmtTact         int           `json:"aj_lmt_tact"`
	CeJsPost          int           `json:"ce_js_post"`
	InitTime          int           `json:"init_time"`
	Informinfo        string        `json:"informinfo"`
	Prevfid           int           `json:"prevfid"`
	Fidcnt            int           `json:"fidcnt"`
	SensorData        string        `json:"sensor_data"`
	Ins               string        `json:"ins"`
	Cns               string        `json:"cns"`
	Engetloc          int           `json:"enGetLoc"`
	Enreaddocurl      int           `json:"enReadDocUrl"`
	Disfpcalontimeout int           `json:"disFpCalOnTimeout"`
	Xagg              int           `json:"xagg"`
	Pen               int           `json:"pen"`
	Brow              string        `json:"brow"`
	Browver           string        `json:"browver"`
	Psub              string        `json:"psub"`
	Lang              string        `json:"lang"`
	Prod              string        `json:"prod"`
	Plen              int           `json:"plen"`
	DoadmaEn          int           `json:"doadma_en"`
	Sdfn              []interface{} `json:"sdfn"`
	D2                int           `json:"d2"`
	D3                int           `json:"d3"`
	Thr               int           `json:"thr"`
	Cs                string        `json:"cs"`
	Hn                string        `json:"hn"`
	Z1                int           `json:"z1"`
	O9                int           `json:"o9"`
	Vc                string        `json:"vc"`
	Y1                int           `json:"y1"`
	Ta                int           `json:"ta"`
	Tst               int           `json:"tst"`
	TTst              int64         `json:"t_tst"`
	Ckie              string        `json:"ckie"`
	NCk               string        `json:"n_ck"`
	Ckurl             int           `json:"ckurl"`
	Bm                bool          `json:"bm"`
	Mr                string        `json:"mr"`
	Altfonts          bool          `json:"altFonts"`
	Rst               bool          `json:"rst"`
	Runfonts          bool          `json:"runFonts"`
	Fsp               bool          `json:"fsp"`
	Firstload         bool          `json:"firstLoad"`
	Pstate            bool          `json:"pstate"`
	MnMcLmt           int           `json:"mn_mc_lmt"`
	MnState           int           `json:"mn_state"`
	MnMcIndx          int           `json:"mn_mc_indx"`
	MnSen             int           `json:"mn_sen"`
	MnTout            int           `json:"mn_tout"`
	MnStout           int           `json:"mn_stout"`
	MnCt              int           `json:"mn_ct"`
	MnCc              string        `json:"mn_cc"`
	MnCd              int           `json:"mn_cd"`
	MnLc              []interface{} `json:"mn_lc"`
	MnLd              []interface{} `json:"mn_ld"`
	MnLcl             int           `json:"mn_lcl"`
	MnAl              []interface{} `json:"mn_al"`
	MnIl              []interface{} `json:"mn_il"`
	MnTcl             []interface{} `json:"mn_tcl"`
	MnR               []interface{} `json:"mn_r"`
	MnRt              int           `json:"mn_rt"`
	MnWt              int           `json:"mn_wt"`
	MnAbck            string        `json:"mn_abck"`
	MnPsn             string        `json:"mn_psn"`
	MnTs              string        `json:"mn_ts"`
	MnLg              []interface{} `json:"mn_lg"`
	Loap              int           `json:"loap"`
	Dcs               int           `json:"dcs"`
	Listfunctions     struct {
	} `json:"listFunctions"`
	StartTs     int64  `json:"start_ts"`
	Kact        string `json:"kact"`
	KeCnt       int    `json:"ke_cnt"`
	KeVel       int    `json:"ke_vel"`
	Mact        string `json:"mact"`
	MmeCnt      int    `json:"mme_cnt"`
	MduceCnt    int    `json:"mduce_cnt"`
	MeVel       int    `json:"me_vel"`
	Pact        string `json:"pact"`
	PmeCnt      int    `json:"pme_cnt"`
	PduceCnt    int    `json:"pduce_cnt"`
	PeVel       int    `json:"pe_vel"`
	Tact        string `json:"tact"`
	TmeCnt      int    `json:"tme_cnt"`
	TduceCnt    int    `json:"tduce_cnt"`
	TeVel       int    `json:"te_vel"`
	Doact       string `json:"doact"`
	DoeCnt      int    `json:"doe_cnt"`
	DoeVel      int    `json:"doe_vel"`
	Dmact       string `json:"dmact"`
	DmeCnt      int    `json:"dme_cnt"`
	DmeVel      int    `json:"dme_vel"`
	Vcact       string `json:"vcact"`
	VcCnt       int    `json:"vc_cnt"`
	AjIndx      int    `json:"aj_indx"`
	AjSs        int    `json:"aj_ss"`
	AjType      int    `json:"aj_type"`
	AjIndxDoact int    `json:"aj_indx_doact"`
	AjIndxDmact int    `json:"aj_indx_dmact"`
	AjIndxTact  int    `json:"aj_indx_tact"`
	MeCnt       int    `json:"me_cnt"`
	PeCnt       int    `json:"pe_cnt"`
	TeCnt       int    `json:"te_cnt"`
	NavPerm     string `json:"nav_perm"`
	Brv         int    `json:"brv"`
	Hbcalc      bool   `json:"hbCalc"`
	Fmh         string `json:"fmh"`
	Fmz         int    `json:"fmz"`
	SSH         string `json:"ssh"`
	Wv          string `json:"wv"`
	Wr          string `json:"wr"`
	Weh         string `json:"weh"`
	Wl          int    `json:"wl"`
	Wen         int    `json:"wen"`
	Den         int    `json:"den"`
}

//jrs, returns 2 float64 variables, it does some weird shit, i think it has to do with time parsing not sure
func (bm *Bmak) jrs(t int64) (float64, float64) {
	var a float64
	var e string
	var n int
	var m bool
	var o []int64
	a = math.Floor(1e5*rand.Float64() + 1e4)
	e = fmt.Sprintf("%f", float64(t)+a)
	n = 0
	m = len(e) >= 18
	for len(o) < 6 {
		intN, _ := strconv.ParseInt(e[n:n+2], 10, 32)
		o = append(o, intN)
		if m {
			n += 3
		} else {
			n += 2
		}
	}
	return a, bm.cal_dis(o)
}

//cal_dis does some weird math things that i dont rly understand but i was lucky enough to just basically copy and paste it
func (bm *Bmak) cal_dis(t []int64) float64 {
	var a = t[0] - t[1]
	var e = t[2] - t[3]
	var n = t[4] - t[5]
	var o = math.Sqrt(float64(a*a + e*e + n*n))
	return math.Floor(o)
}

//pi just takes a float and returns the int part as an int type var
func (bm *Bmak) pi(t float64) int {
	return int(t)
}

//ab does some weird ascii code shit, dont rly understand it but it returns an int, i think it just sums up all the ascii
//codes that are less than 128
func (bm *Bmak) ab(t *string) int {
	if t == nil {
		return -1
	}
	var a = 0
	for e := 0; e < len(*t); e++ {
		n := []rune(*t)[e]
		if n < 128 {
			a += int(n)
		}
	}
	return a
}

//od, does some weird ascii character shit that returns a result, its huge twice inside the generateSensor (bpd) function
func (bm *Bmak) od(a, t string) string {
	var e []string
	n := len(t)
	if n > 0 {
		for o, v := range a {
			r := string(v)
			i := t[o%n]

			vInt := int(v)
			m := bm.rir(vInt, 47, 57, int(i))
			if m != vInt {
				r = string(rune(m))
			}
			e = append(e, r)
		}
		if len(e) > 0 {
			return strings.Join(e, "")
		}
	}
	return a
}
//rir takes in 4 diff params, if the a param is between 47 and 57 it does some calculations, at the end of the day it returns
//always an int
func (bm *Bmak) rir(a, b, c, d int) int {
	if a > 47 && a <= 57 {
		a += d % (c - b)
		if a > c {
			a = a - c + b
		}
	}
	return a
}

//cc returns a function that returns an int64, weird right? It gets even weirder, that func's return value is determined by
//the t parameter of the parent funcion
func (bm *Bmak) cc(t int) func(t, a int) int64 {
	var a = t % 4
	if a == 2 {
		a = 3
	}
	var e = 42 + a
	n := func(t, a int) int64 {
		return 0
	}
	if e == 42 {
		n = func(t, a int) int64 {
			return int64(t * a)
		}
	} else if e == 43 {
		n = func(t, a int) int64 {
			return int64(t + a)
		}
	} else {
		n = func(t, a int) int64 {
			return int64(t - a)
		}
	}

	return n
}