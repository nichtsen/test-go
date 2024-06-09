# dump interface lo, with incoming direction, tcp protocol and port 7077
sudo tcpdump -i lo -A -Q in 'tcp port 7077
