rm -rf /data
docker rm -v -f $(docker ps -aq)