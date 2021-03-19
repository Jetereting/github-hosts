# github-hosts
更新 github hosts, 解决大陆 github 443

#### 首次使用(适用于mac,linux; windows请自行修改路径)
```shell script
# git clone https://github.com/Jetereting/github-hosts.git -b master &&
git clone https://gitee.com/jetereting/github-hosts.git -b master &&
cd github-hosts &&
go build -o githost &&
mv githost /usr/local/bin &&
sudo githost
```

#### 非首次使用
```shell script
sudo githost
```
