#!/bin/bash

# start 1
echo "begin keeper..."
/root/keeper/keeper -name=test-server -ip=47.90.62.113 -port=13389 > /var/log/keeper.log 2>&1 &
echo "wait for 5 second..."
sleep 5
echo "begin ss-server..."
# start 2
/bin/ss-server -c config.json -u > /var/log/ss-server.log 2>&1 &

echo "done..."
# just keep this script running
while true; do
       sleep 1
done