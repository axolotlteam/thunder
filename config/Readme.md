# Config



## ConsulKV
```golang
package main

import (
    "github.com/axolotlteam/thunder/config"
)

func main() {
    v , err := config.ConsulKV("consulhost" , "key" , "ftype{json/yaml}")
    if err != nil {
        log.Fatal(err)
    }

    v.GetString("xxx.xxx")


    x := config.Database{}

    // maping to struct
    err := viper.Unmarshal(&x)

}

```
