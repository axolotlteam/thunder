# logger


## Demo


```golang
package main

import (
    "github.com/axolotlteam/thunder/logger"
)


func init () {

    logger.NewLogrus()
    logger.SetLevel(logger.Debug)
    logger.SetServiceInfo("thunder")

}

func main() {

    logger.WithFields(logger.Fields{
        "msg" : "xxxxx",
    }).Info("this a test log")
}


{"fields.msg":"xxxxx","file":"log_test.go:13","func":"github.com/axolotlteam/thunder/logger.Test_Logrus()","hostname":"thunder","level":"info","msg":"this is a test log","service":"thunder","time":"2019-09-22T01:23:54+08:00"}


```




