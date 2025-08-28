package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"buvid3=6B96BB52-859F-4DBE-C237-D5D184F96AE040371infoc; b_nut=1750684040; _uuid=AEFF68BC-10ADC-A31D-7A10D-226B67B10B102840879infoc; buvid_fp=80d10e63269f2dc75007e0e1e5be4a6d; enable_web_push=DISABLE; enable_feed_channel=ENABLE; home_feed_column=5; browser_resolution=1528-738; rpdid=0zbfVLKTv2|9GLEHrzt|2vO|3w1UtGTB; DedeUserID=493239443; DedeUserID__ckMd5=02a9d31303da4905; CURRENT_QUALITY=80; header_theme_version=OPEN; theme-tip-show=SHOWED; theme-avatar-tip-show=SHOWED; buvid4=C9DD301B-AA4A-2C9F-C98A-7C00C303983E30895-022012415-SK3hbof5R8mR/P9KIc92lcN/B/tuRecWKl7mrGcS1SJpc+0da5MFdw%3D%3D; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY0NjM4NjQsImlhdCI6MTc1NjIwNDYwNCwicGx0IjotMX0.-abBVNt7E2lbxThP2ZQDQZc2ACTzv8N8iVuD0VNIJWU; bili_ticket_expires=1756463804; SESSDATA=7f316349%2C1771756666%2C3b119%2A81CjAoCaz17WAQLnXmcre1XOQBESyuiuhltLNJHUrtEVziXKuM8v1M6vN9klzTTP9N9nQSVlo1X3RRVE9mRXl6Tk1NY2xNMjRpclIwUHlXY1B6NmE0a1RSSGhCbUtFWkR3LXlUUktWckZfSWFJZFowcThqUmJoTURjbWd5eEowNnN0SWllVEh6T2F3IIEC; bili_jct=2487b7e40d2cc17c8d01d60e06f3433f; sid=5i66okdl; CURRENT_FNVAL=4048; b_lsid=37453432_198EEC9EFE3; bsource=search_bing; bp_t_offset_493239443=1106016261256511488"`
	KKCookie       string `yaml:"__itrace_wid=35d3e356-be28-4ef5-add2-cd901199ffdb; b-user-id=b0d0c7d2-b499-33d5-4a34-7b54c71045e2; _UP_A4A_11_=wb9ca15f764d40909b92fa6c94f1304f; _qk_bx_ck_v1=eyJkZXZpY2VJZCI6ImVDeSNBQVBzdFBuVEhqeExlRVF0NzQ3N1NSTEtobEc5YXBVeHJveUR1Vm5WWUZhMXlxWEdZUDhTVUxnczIvdnR2RldUdXZFPSIsImRldmljZUZpbmdlcnByaW50IjoiODBkMTBlNjMyNjlmMmRjNzUwMDdlMGUxZTViZTRhNmQifQ==; __itrace_wid=35d3e356-be28-4ef5-add2-cd901199ffdb; __sdid=AAROX61SnjQydqnIkH3rwe8EPy+I+ors8ugzHub4lxkf2bnOxymPOyivmr++EpV4qgy45gI2jblFS+cfb62vi9w/OmrLq2SzV3OazyizIeMI+g==; _UP_D_=pc; __pus=d8bb01fc273e0fe4665a4973925f56d2AAQmNa/cznmXIb6cXREC2J62JAP3HmAkvmet0TaVzLro7IS51ByJ00L/GiKvo+rnyKewJY/b7HakGT5S9WpQCMN3; __kp=88986bc0-7f5a-11f0-b18e-f3a891fe767b; __kps=AAQR+iqgeyC2Bq1WeagOaOZx; __ktd=qKbb1K6neNtgZlivhZPvYg==; __uid=AAQR+iqgeyC2Bq1WeagOaOZx; xlly_s=1; isg=BPf3l3_IU2MYmNhWwh6qGlR6hutBvMseWMWB8EmkqkYt-BU6UY-GbleR31imL6OW; tfstk=gCYmDhfBwnSXCD4bmfbjcJfb-yiJlZ_1FdUOBNBZ4TW5BjO9_NzwBCQAbxO9EOWP_nU47r7Ps19hgszZnLcMHCaO3tLOSGJduGtABtBGSC9nwv3KJIOf5ZkKp2FdjQIORSya0RIP46bGg_jUJ0Of5Nk8wPoK-I9jEg210Nllz1fT0NSVbgllGsbauGz448W5_N7NbGRPU65Q7OWa0bAPF17NQd7ZZaW5_NWwQNkMNmWX7U8rI7R3UWbs2X1lm9RVEyRJrskNK2677P8lieXeg7za7UfcmF9zvZQ5DhRCAEdtSz_vt37PsFGUjOjHbEQyuYuFVG-wvGx83jpeIQJ1kUla09-vHa82x5zwZZXwKUSUnYXelQ8Czg3EtQ8WHITk65uNwpByGE7i8X_cut7lNFH7uOxybEIf5-kGCUAyuhjyGurUhKz1afLzflsVN_Xp7MXitxsimzhoZkl10_1fpbcuflsVN_XKZbqE5i55G9C..; __puus=8842cbfe51d46b57dc9a275c339b7acdAAQYElJ0mbdWXsBPBOLgaEr1wPwqJtUwCkVmw8wY9FWmETWrQjbatj6I9NfK7B5IlsLUVgJiJhorhCqZrWxy0w14obfyg+hhtEvGlYjvE23kWkaW1JJ7eFv/0OIeF/4Xwt2E2eCcsG4Vqd9ZQ0N0rZ+OrDi1NJDeGzmRnh+AbY+BGRJ18JqkagDeejBLJB8g7HfZvYwn0Dufafg5vjkGjzR5"`
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
