from 16620808100/shadowsock:v1
WORKDIR /root/keeper
RUN ["apt","update","-y"]
RUN ["apt","install","libpcap-dev","-y"]
RUN ["git","fetch","--all"]
RUN ["git","reset","--hard","origin/master"]
COPY ./keeper /root/keeper/keeper
COPY ./goiftop /root/keeper/goiftop
COPY ./run.sh /root/keeper/run.sh
ENTRYPOINT ["/bin/bash","/root/keeper/run.sh"]
