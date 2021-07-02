# Hashicorp Sentinel Exceptions Parser

### Summary
Hashicorp Sentiel Excaptions management tool. This code goes through the included `exceptions.json` and modifies the `enforcement_level` within the `sentinel.hcl` files.


### Installing Golang
1. Update Packages
```
sudo apt update
sudo apt upgrade -y
```
2. Download latest version: https://golang.org/dl/
3. Extract downloaded package and add to path.
```
sudo tar -C /usr/local/ -xzf <filename>
cd /usr/local/
sudo nano $HOME/.profile
export PATH=$PATH:/usr/local/go/bin
source .profile
```
4. Verify Installation
```
go version
```

### Required Packages
```
go get https://github.com/hashicorp/hcl
go get https://github.com/zclconf/go-cty
go get https://github.com/hashicorp/hcl/tree/main/hclsyntax
go get https://github.com/hashicorp/hcl/tree/main/hclwrite
```

