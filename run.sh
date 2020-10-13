#!/bin/bash

# start 1
echo "begin keeper..."
/root/keeper -name=$SERVERNAME -ip=$IP -port=$SERVERPORT -bandwidth=$BANDWIDTH > /var/log/keeper.log 2>&1 &
echo "wait for 5 second..."
sleep 2
echo "begin ss-server..."
# start 2
/bin/ss-server -c /root/config.json -u > /var/log/ss-server.log 2>&1 &

# echo "begin goiftop..."
# /root/keeper/goiftop -i eth0 2>&1 &

echo "done..."
# just keep this script running
while true; do
       sleep 1
done
