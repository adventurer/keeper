from ubuntu
WORKDIR /root
RUN ["apt","update","-y"]
RUN ["apt","install","git","-y"]
RUN ["apt","install","shadowsocks-libev","-y"]
RUN ["apt","install","iptables","-y"]
RUN ["apt","install","libpcap-dev","-y"]
COPY ./keeper /root/keeper
COPY ./goiftop /root/goiftop
COPY ./run.sh /root/run.sh
ENTRYPOINT ["/bin/bash","/root/run.sh"]
