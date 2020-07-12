package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	token = "Bearer "
	// 获取首页列表数据接口
	homeURL = "https://mastodon.social/api/v1/timelines/home"
	// 点赞接口，注意不包含 id 以及尾部的 /favourite
	starURL = "https://mastodon.social/api/v1/statuses/"
	// 需要点赞的用户
	users []string
	// 请求一次首页数据（20条）的间隔，默认60s
	interval int16 = 60
)

func main() {
	color.Green("读取配置...")
	readConfig()
	for {
		color.Green("开始请求首页数据（20条）...")
		query()

		color.Green("一次请求结束...")
		time.Sleep(time.Second * time.Duration(interval))
	}
}

type Account struct {
	Username string
}

type Reblog struct {
	Favourited bool
	Reblogged  bool
	Reblogount int8 `json:"reblogs_count"`
}

type Result struct {
	ID         string
	Favourited bool
	Account    Account
	Reblog     Reblog
}

// 发起home请求，然后判断是否需要点赞
func query() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", homeURL, nil)

	request.Header.Add("authorization", token)

	if err != nil {
		color.Red(err.Error())
		return
	}

	resp, requestErr := client.Do(request) // resp 可能为 nil，不能读取 Body。所以不能先defer
	if requestErr != nil {
		color.Red(requestErr.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red(readErr.Error())
		return
	}

	var result []Result

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		color.Red("解析字符串出错！")
		return
	}

	for _, item := range result {
		reblog := item.Reblog
		username := item.Account.Username
		id := item.ID

		// 如果为空，说明对用户没有限制，跳过
		// 如果不为空，说明只需要给这里面的用户点赞
		if len(users) != 0 {
			// 如果不需要点赞，跳过
			if !has(users, username) {
				continue
			}
		}

		// 如果有转嘟且已经点赞 或 嘟文已经点赞，跳过
		if reblog.Favourited || item.Favourited {
			continue
		}

		// 没点赞，不管是转嘟还是自己嘟文。
		star(id, username)

		// 点赞的间隔
		time.Sleep(time.Second * 2)
	}
}

// 发起点赞
func star(id string, username string) {
	color.Cyan("准备给【%s】点赞, 嘟文id: %s \n", username, id)

	client := &http.Client{}

	request, err := http.NewRequest("POST", starURL+id+"/favourite", nil)

	request.Header.Add("authorization", token)

	if err != nil {
		color.Red(err.Error())
		return
	}

	resp, doError := client.Do(request) // resp 可能为 nil，不能读取 Body。所以不能先defer
	if doError != nil {
		color.Red(doError.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	color.Blue("%s，点赞完成✅\n", id)
}

type Config struct {
	Token    string
	HomeURL  string
	StarURL  string
	Users    []string
	Interval int16
}

// 读取本地配置文件
func readConfig() {
	file, fileErr := os.Open("./config.json")
	defer file.Close()

	if fileErr != nil {
		color.Red(fileErr.Error())
		os.Exit(1)
	}

	var conf Config
	err := json.NewDecoder(file).Decode(&conf)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	if conf.Token == "" {
		color.Red("没有检测到Token，程序退出！请在 config.json 中添加 token 字段。")
		os.Exit(1)
	} else {
		token += conf.Token
	}

	if conf.HomeURL == "" {
		color.Yellow("HomeURL 为空，将使用默认URL：%s\n", homeURL)
	} else {
		homeURL = conf.HomeURL
	}

	if conf.StarURL == "" {
		color.Yellow("StarURL 为空，将使用默认URL：%s\n", starURL)
	} else {
		starURL = conf.StarURL
	}

	if len(conf.Users) == 0 {
		color.Yellow("需要点赞的用户数(users)为空，将对主页的所有嘟文进行点赞！")
	} else {
		users = conf.Users
	}

	if conf.Interval == 0 {
		color.Yellow("没有设置请求间隔时间，将使用默认时间：60s")
	} else {
		interval = conf.Interval
	}

}

// 判断数组是否存在某个值
func has(arr []string, value string) bool {
	for _, e := range arr {
		if e == value {
			return true
		}
	}

	return false
}
