package main
import (
	"os"
	"fmt"
	"net"
	"log"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {							//! <-- For Linux
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	
}

func callClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is not supported!")
	}
}

func getIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func main() {
	var cmd string
	var user string
	
	callClear()
	
	fmt.Print("Enter username: ")
	fmt.Scanln(&user)
	
	fmt.Print("\033[H\033[2J")
	
	callClear() //! Clear terminal
	
	for cmd != "exit()" {
		fmt.Print(">>> ")
		fmt.Scanln(&cmd)
		if cmd == "whoami" {
			fmt.Println("User", user)
		} else if cmd == "exit()" {
			os.Exit(0)
		} else if cmd == "ip" {
			fmt.Println(getIP())
		} else {
			fmt.Println("Only command available at the moment is \"whoami\"")
		}
	}
}
