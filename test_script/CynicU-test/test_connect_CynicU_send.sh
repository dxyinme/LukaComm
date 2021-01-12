#!/bin/bash
echo "start cynic send server"
nohup ./CynicU_send_server &
echo "sleep 5 second"
sleep 5s
echo "start cynicU send client"

echo "cnt clients =" $1

declare -i cntClients=$1
declare -i i=1

while((i <= cntClients))
do
  echo $i "client"
  nohup ./CynicU_send_client -isSleep &
  let i++
  sleep 0.1s
done

echo "CynicU_send_server pid =" `ps -aux | grep CynicU_send_server | grep -v grep | awk '{print $2}'`
