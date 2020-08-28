# keeper
docker run -itd --env SERVERPORT=10088 --env SERVERNAME=全国通用02 --env IP=47.107.112.182 -p 10089:10089/udp -p 10089:10089 16620808100/shadowsock:v9

docker build . -t 16620808100/shadowsock:v9