#!/bin/bash
kill -9 `ps -aux | grep CynicU_send_client | grep -v grep | awk '{print $2}'`