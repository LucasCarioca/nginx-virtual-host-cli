#!/etc/bin/sh

curl -L https://github.com/LucasCarioca/proxy-reverse-proxy-cli/releases/download/v0.1.2/proxy_0.1.2_Darwin_x86_64.tar.gz -o proxy.tar.gz
tar -xzf proxy.tar.gz
sudo mv proxy /usr/local/bin/proxy
