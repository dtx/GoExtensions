#This file will change the GOPATH to the current directory and add the bin directory to the path. 
#Only use this for testing purposes if you dont know what a GOPATH is
export GOPATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export PATH=$PATH:$GOPATH/bin
echo $GOPATH
