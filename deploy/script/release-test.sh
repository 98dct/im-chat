need_start_server_shell=(
  im-ws-test.sh
  im-api-test.sh
  im-rpc-test.sh
  user-rpc-test.sh
  user-api-test.sh
  social-rpc-test.sh
  social-api-test.sh
  task-mq-test.sh
)

for i in ${need_start_server_shell[*]}; do
  chmod +x $i
  ./$i
done

docker ps
docker exec -it etcd-chat etcdctl get --prefix ""
