# Setup ip-up script in macOS

- issue: [https://github.com/adrienverge/openfortivpn/wiki#method-one---short-version---pass-routes-by-pppd-ipparam](https://github.com/adrienverge/openfortivpn/wiki#method-one---short-version---pass-routes-by-pppd-ipparam)
- path: /etc/ppp/ip-up
- don't forget: `chmod a+x /etc/ppp/ip-up`

```sh
# !/bin/sh -e
now=`date +%Y-%m-%d_%Hh%Mm%Ss`
logfile=/var/log/ip-up.log

echo "" >> $logfile
echo "$0 called at $now with following params: ${6}" >> $logfile
echo "The VPN interface (e.g. ppp0): $1" >> $logfile
echo "Unknown, was 0, in my case: $2" >> $logfile
echo "IP of the VPN server: $3" >> $logfile
echo "VPN gateway address: $4" >> $logfile
echo "Regular (non-vpn) gateway for your lan connections: $5" >> $logfile

IPS=${6}
for IP in ${IPS}
do
 /sbin/route add $IP -interface $1 >> $logfile 2>&1
done

echo "./ip-up DONE" >> $logfile
```
