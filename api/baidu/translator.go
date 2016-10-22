package baidu

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bitly/go-simplejson"
)

const (
	appid   string = "20151113000005375"
	key     string = "s7TBETv56ZaZjyYoS472"
	baseurl string = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

func genarateSign(query string) (sign, salt string) {
	salt = strconv.FormatInt(rand.Int63n(10000), 5)
	sign = appid + query + salt + key

	m := md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))
	return
}

// Translator translate words from one language to another.
func Translator(query, from, to string) (result string, err error) {
	si, sa := genarateSign(query)
	q := url.QueryEscape(query)
	transURL := baseurl + "?appid=" + appid + "&q=" + q + "&from=" + from +
		"&to=" + to + "&salt=" + sa + "&sign=" + si

	res, err := http.Get(transURL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		data, _ := simplejson.NewFromReader(res.Body)
		result, _ = data.Get("trans_result").GetIndex(0).Get("dst").String()
	}

	return
}
