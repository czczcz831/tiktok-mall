#!/bin/bash

. scripts/list_app.sh

get_app_list

readonly root_path=$(pwd)
session="microservices"

# 创建新的 tmux session
tmux new-session -d -s $session

# 为每个微服务创建一个窗口并运行命令
for app_path in ${app_list[*]}; do
    service_name=$(basename $app_path)
    tmux new-window -t $session -n $service_name
    tmux send-keys -t $session:$service_name "cd ${root_path}/${app_path} && air" C-m
done

# 选择第一个窗口
tmux select-window -t $session:0

# 附加到 tmux session
tmux attach-session -t $session