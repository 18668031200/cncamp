# cncamp
golang homework

docker run shell
#######

docker run -d -p 9100:9100 -c 2 -m 4g --name=httpserver ygdxd/httpserver:v1.0

nsenter
#######
PID=$(docker inspect --format {{.State.Pid}} httpserver)
nsenter -m -u -i -n -p -t $PID ifconfig
