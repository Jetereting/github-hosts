package main

import (
	"fmt"
	util "github.com/Jetereting/go_util"
	"github.com/astaxie/beego/httplib"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	hostsPath := "/etc/hosts"
	if runtime.GOOS == "windows" {
		hostsPath = "C:/Windows/System32/drivers/etc/hosts"
	}
	oldHosts := getFileContent(hostsPath)
	newHosts := ""
	isFirstByAu := !strings.Contains(oldHosts, "github end by Au.")
	if strings.Contains(oldHosts, "github.com") && isFirstByAu {
		fmt.Println("请先将 " + hostsPath + " 中关于 github.com 的删除掉")
		return
	}

	githubHost := "\n# github start by Au\t" + time.Now().Format("2006-01-02 15:04:05")
	//github.com
	g, err := httplib.Get("https://github.com.ipaddress.com/").String()
	if err != nil {
		return
	}
	g = util.Str(g).GetBetween(`class="comma-separated"><li>`, `</li>`)
	githubHost += "\n" + g + "\t" + "github.com"

	//github.global.ssl.fastly.net
	gSSL, err := httplib.Get("https://fastly.net.ipaddress.com/github.global.ssl.fastly.net#ipinfo").String()
	if err != nil {
		return
	}
	gSSL = util.Str(gSSL).GetBetween(`class="comma-separated"><li>`, `</li>`)
	githubHost += "\n" + gSSL + "\t" + "github.global.ssl.fastly.net"

	//assets-cdn.github.com
	gAssets, err := httplib.Get("https://github.com.ipaddress.com/assets-cdn.github.com").String()
	if err != nil {
		return
	}
	gAssets = util.Str(gAssets).GetBetween(`<div>assets-cdn.github.com resolves to the following 4 IPv4 addresses:`, `</div>`)
	for _, v := range strings.Split(gAssets, "</a>") {
		gAssets = util.Str(v).GetBetween(`/ipv4/`, `">`)
		if !strings.Contains(gAssets, ".") {
			continue
		}
		githubHost += "\n" + gAssets + "\t" + "assets-cdn.github.com"
	}
	githubHost += "\n# github end by Au.\n"

	if isFirstByAu {
		newHosts = oldHosts + githubHost
	} else {
		newHosts = strings.ReplaceAll(oldHosts, util.Str(oldHosts).GetBetween(`github start by Au`, `github end by Au.`), util.Str(githubHost).GetBetween(`github start by Au`, `github end by Au.`))
	}
	fmt.Println(newHosts)
	createFile(hostsPath, newHosts)
}

func getFileContent(fileName string) (content string) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	content = string(buf)
	return
}
func createFile(fileName, content string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		_, err = f.Write([]byte(content))
		if err != nil {
			fmt.Println(err)
		}
	}
}
