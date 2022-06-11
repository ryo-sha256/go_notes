package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

    arguments := os.Args

    if len(arguments) == 1 {
        fmt.Println("Please provide port number")
        return
    }


    // create the string variable that can be used 
    // as port variable
    PORT := ":" + arguments[1]
    l, err := net.Listen("tcp", PORT)
    if err != nil {
            fmt.Println(err)
            log.Fatalln("PORT is not provided")
    }
    defer l.Close()
   
    c, err := l.Accept()
    if err != nil {
        fmt.Println(err)
        log.Fatalln(err)
    }

    for {
        netData, err := bufio.NewReader(c).ReadString('\n')
        if err != nil {
                fmt.Println(err)
                return 
        }

        if strings.TrimSpace(string(netData)) == "STOP" {
                fmt.Println("Exiting TCP server")
                return
        }
    
        fmt.Print("Received >>>>", string(netData))
        t := time.Now() 
        myTime := t.Format(time.RFC3339) + "\n"
        // type casting
        c.Write([]byte(myTime))
    }

}
