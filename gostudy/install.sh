# GOROOT go install path; GOPATH, Multiple paths can be set , the first is default package path, for example : $HOME/go
# GOPATH/
#   src ，自己的代码以及第三方依赖的代码
# 
# 

# go env # show all vars , file: $HOME/.config/go/env 
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on
export GOPROXY=https://goproxy.cn