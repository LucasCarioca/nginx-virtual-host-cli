# NGINX Virtual Host Management CLI

## Install
Using prepared install-scripts. These can also be used to upgrade to newer versions. Please see the releases section of this repo for information on each release.
```shell
curl -L https://raw.githubusercontent.com/LucasCarioca/proxy-reverse-proxy-cli/main/scripts/install-<PLATFORM>.sh | sh -s <VERSION>
```
Examples 
```shell
# on a linux machine
curl -L https://raw.githubusercontent.com/LucasCarioca/proxy-reverse-proxy-cli/main/scripts/install-linux.sh | sh -s 0.3.0
# on a macOS machine (x86)
curl -L https://raw.githubusercontent.com/LucasCarioca/proxy-reverse-proxy-cli/main/scripts/install-linux.sh | sh -s 0.3.0
```

### Manual installation
```shell
curl -L https://github.com/LucasCarioca/proxy-reverse-proxy-cli/releases/download/<RELEASE>/<BINARY_NAME>.tar.gz -o proxy.tar.gz
tar -xzf proxy.tar.gz
sudo mv proxy /usr/local/bin/proxy
```

## Uninstall 
```shell
curl -L https://raw.githubusercontent.com/LucasCarioca/proxy-reverse-proxy-cli/main/scripts/uninstall.sh | sh
```




