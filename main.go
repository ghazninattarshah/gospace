package main

import (
    channel "github.com/ghazninattarshah/gospace/channels"
    conv "github.com/ghazninattarshah/gospace/conv"
    cpy "github.com/ghazninattarshah/gospace/copy"
    mp "github.com/ghazninattarshah/gospace/map"
    slice "github.com/ghazninattarshah/gospace/slice"
    strs "github.com/ghazninattarshah/gospace/structs"
    "github.com/ghazninattarshah/gospace/singleton"
    "github.com/ghazninattarshah/gospace/helloworld"
)

func main() {

   helloworld.Run()
   channel.Run()
   conv.Run()
   cpy.Run()
   mp.Run()
   slice.Run()
   strs.Run()
   singleton.Run()
} 