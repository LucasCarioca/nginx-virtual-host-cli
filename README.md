# NGINX Virtual Host Management CLI

![GitHub release (latest by date)](https://img.shields.io/github/v/release/LucasCarioca/nginx-virtual-host-cli)
![GitHub Release Date](https://img.shields.io/github/release-date/LucasCarioca/nginx-virtual-host-cli)
![GitHub all releases](https://img.shields.io/github/downloads/LucasCarioca/nginx-virtual-host-cli/total)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/nginx-virtual-host-cli/Release?label=release)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/LucasCarioca/nginx-virtual-host-cli/CI?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/LucasCarioca/nginx-virtual-host-cli)](https://goreportcard.com/report/github.com/LucasCarioca/nginx-virtual-host-cli)
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




