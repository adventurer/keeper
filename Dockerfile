from 16620808100/shadowsock:v1
WORKDIR /root/keeper
RUN ["git","pull"]
ENTRYPOINT ["/bin/bash","/root/keeper/run.sh"]