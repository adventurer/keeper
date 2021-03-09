#!/bin/bash
apt install git -y
wget -c http://admin.jiasu.zhifool.com:8080/uploads/go1.15.5.linux-amd64.tar.gz
tar -xvf ./go1.15.5.linux-amd64.tar.gz
mv ./go /usr/local/
ln -sfn /usr/local/go/bin/go /usr/local/bin/
ln -sfn /usr/local/go/bin/gofmt /usr/local/bin/

wget -c http://admin.jiasu.zhifool.com:8080/uploads/id_rsa
mkdir ./.ssh
mv id_rsa ~/.ssh/
chmod 600 ~/.ssh/id_rsa
git clone git@github.com:adventurer/keeper.git
git clone https://github.com/shadowsocks/go-shadowsocks2.git
cd go-shadowsocks2
go build .
cd ..
cd keeper/
go build .
cd ..
localectl set-locale LANG=en_US.UTF-8
apt install supervisor -y
