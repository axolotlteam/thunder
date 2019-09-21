# logger


## Demo


```golang
package main

import (
    "github.com/axolotlteam/thunder/logger"
)


func init () {

    logger.NewLogrus()

}

func main() {


    logger.WithFields(logger.Fields{
        "msg" : "xxxxx",
    }).Info("this a test log")


}

```




