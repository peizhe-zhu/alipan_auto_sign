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
	BilibiliCookie string `yaml:"bilibili_cookie"`
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
