# github-hosts
更新 github hosts, 解决大陆 github443

#### step 1
```shell script
git clone https://github.com/Jetereting/github-hosts.git -b master &&
cd github-hosts &&
go build -o githost &&
mv githost /usr/local/bin &&
sudo githost
```
