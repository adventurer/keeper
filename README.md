# keeper
docker run --cap-add NET_ADMIN --cap-add NET_RAW -itd --env SERVERPORT=10088 --env SERVERNAME=全国通用02 --env IP=47.57.180.168 -p 10089:10089/udp -p 10089:10089 16620808100/shadowsock:v22

docker build . -t 16620808100/shadowsock:v9