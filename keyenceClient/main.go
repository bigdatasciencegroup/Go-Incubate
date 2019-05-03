package main

import "net"
import "fmt"
import "bufio"

func main() {

  // connect to this socket
  conn, err := net.Dial("tcp", "192.168.0.10:8500")
  if err !=nil {
	  fmt.Println("Error",err)
	  return
  }
  for { 
    // read in input from stdin
    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Text to send: ")
    // text, _ := reader.ReadString('\n')
    // // send to socket
    // fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}