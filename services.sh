#!/bin/bash

api_service="./cmd/api/api"
comment_service="./cmd/comment/comment"
favorite_service="./cmd/favorite/favorite"
feed_service="./cmd/feed/feed"
message_service="./cmd/message/message"
consumers_service="./cmd/mykafka/consumers/consumers"
publish_service="./cmd/publish/publish"
relation_service="./cmd/relation/relation"
user_service="./cmd/user/user"

function killProsess() {
	NAME=$1
	echo $NAME

    # 获取进程 PID
	# PID=$(ps -e | grep $NAME | awk '{print $1}')
	PID=$(ps -ef|grep $NAME| grep -v grep | awk '{print $2}')

	echo "PID: $PID"
    #  杀死进程
	kill -9 $PID
}

function start() {
	(nohup $api_service &) &&
	(nohup $comment_service &) &&
	(nohup $favorite_service &) &&
	(nohup $feed_service &) &&
	(nohup $message_service &) &&
	(nohup $consumers_service &) &&
	(nohup $publish_service &) &&
	(nohup $relation_service &) &&
	(nohup $user_service &) &&
	ps -ef | grep service
}

function stop() {
	killProsess $api_service
  killProsess $comment_service
  killProsess $favorite_service
  killProsess $feed_service
  killProsess $message_service
  killProsess $consumers_service
  killProsess $publish_service
  killProsess $relation_service
  killProsess $user_service
}

function restart() {
	stop
	start
}

case "$1" in
	start )
		start
		echo "monitor start sucess"
		;;
	stop )
		stop
		echo "monitor stop sucess"
		;;
	restart )
		restart
		echo "monitor restart sucess"
		;;
	* )
		echo "you can use ./services start|stop|restart"
		;;
esac