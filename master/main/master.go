package main

import (
	"github.com/golazycat/lazycron/common"
	"github.com/golazycat/lazycron/master"
)

func main() {

	master.Start(false)

	common.LoopForever()
}
