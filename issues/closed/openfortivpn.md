# openfortivpn

[git](https://github.com/adrienverge/openfortivpn)

## setup

```shell
host={ip}
port={port}
username={user_name@mail.com}
trusted-cert=b297187804c765aae5a583ad79d06aac22eee7b419450a53c31fd4bb9fc71402
set-routes=0
pppd-ipparam={ip.1} {ip.2} {ip.3} {ip.4}
```

## backgroud running

```shell
# background
function bg-start() {
	sudo screen -mS $1 ${@:2}
}
function bg-stop() {
	bg-pid $1 | sudo xargs kill -9
}
function bg-pid() {
	ps aux | grep $1 | awk '{print $2}'
}

# vpn
export HX_VPN_FILE=$HOME/.vpn/hx
alias vpn-start="bg-start vpn-session sudo openfortivpn -c $HX_VPN_FILE"
alias vpn-stop="bg-stop openfortivpn"
function vpn-clean() {
	local addr=`awk -F "=" '/host/{gsub(/ /, "", $2); print $2}' $HX_VPN_FILE`
	sudo route delete $addr
}
function vpn() {
	if pgrep openfortivpn > /dev/null; then
		red "vpn stop"
		vpn-stop
	else
		green "vpn start"
		vpn-start
	fi
}

# Color Print alias
function green {
    printf "\e[32m%s\e[0m\n" "$1"
}

function red {
    printf "\e[33m%s\e[0m\n" "$1"
}

function gray {
    printf "\e[90m%s\e[0m\n" "$1"
}
```
