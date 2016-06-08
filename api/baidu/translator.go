package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

// APPID, KEY, BAIDUTRANS is request from baidu fanyi api
const (
	APPID      string = "20151113000005375"
	KEY        string = "s7TBETv56ZaZjyYoS472"
	BAIDUTRANS string = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

// TransResultJSON save the translate result
type TransResultJSON struct {
	From   string
	To     string
	Result []transResultObject `json:"trans_result"`
}

type transResultObject struct {
	Src string
	Dst string
}

func genarateSign(query string) (sign, salt string) {
	salt = strconv.FormatInt(rand.Int63n(10000), 5)
	sign = APPID + query + salt + KEY

	m := md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))
	return
}

// Translator translate words from one language to another.
func Translator(query, from, to string) (result TransResultJSON, err error) {
	si, sa := genarateSign(query)
	q := url.QueryEscape(query)
	transURL := BAIDUTRANS + "?appid=" + APPID + "&q=" + q + "&from=" + from +
		"&to=" + to + "&salt=" + sa + "&sign=" + si

	res, err := http.Get(transURL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		data, _ := ioutil.ReadAll(res.Body)

		err = json.Unmarshal(data, &result)
		if err != nil {
			return
		}
	}

	return
}
