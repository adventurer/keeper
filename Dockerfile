from ubuntu
WORKDIR /root/keeper
RUN ["git","pull"]
ENTRYPOINT ["/bin/bash","/root/keeper/run.sh"]