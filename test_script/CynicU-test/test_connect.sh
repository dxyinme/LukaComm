#!/bin/bash
echo "start cynic server"
nohup ./CynicU_server &
echo "sleep 5 second"
sleep 5s
echo "start cynic client"

echo "cnt clients =" $1

declare -i cntClients=$1
declare -i i=1

while((i <= cntClients))
do
  echo $i "client"
  nohup ./CynicU_client -isSleep &
  let i++
  sleep 0.1s
done

echo "CynicU_server pid =" `ps -aux | grep CynicU_server | grep -v grep | awk '{print $2}'`
