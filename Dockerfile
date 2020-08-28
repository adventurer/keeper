from 16620808100/shadowsock:v1
WORKDIR /root/keeper
RUN ["git","fetch","--all"]
RUN ["git","reset","--hard","origin/master"]
COPY ./keeper /root/keeper/keeper
ENTRYPOINT ["/bin/bash","/root/keeper/run.sh"]