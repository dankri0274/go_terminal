package main
import (
	"os"
	"fmt"
	"net"
	"log"
	"time"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {						//! <-- For Linux
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")		//! <-- For Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	
}

//! Clear the terminal
func CLS() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is not supported!")
	}
}

//! Get IPv4 of local machine
func getIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

//! Main function
func main() {
	var cmd string
	var user string
	var password string
	var running bool = true

	currentTime := time.Now()
	
	CLS()
	
	fmt.Print("Enter username: ")
	fmt.Scanln(&user)

	CLS() //! Clear terminal

	fmt.Print("Password: ", "\033[8m")
	fmt.Scanln(&password)
	fmt.Print("\033[28m")

	if len(password) < 8 {
		fmt.Println("Password must be atleast 8 characters")
		time.Sleep(time.Second * 2)
	} else {
		CLS()
		for running {
			fmt.Print(user, "@", getIP(), "$ ")
			fmt.Scanln(&cmd)

			//* Commands for SYSTEM

			if cmd == "exit()" {
				running = false
				os.Exit(0)
			} else if cmd == "ip" {
				fmt.Println(getIP())
			} else if cmd == "clear" || cmd == "cls" {
				CLS()
			} else if cmd == "time" {
				fmt.Println("Today is", currentTime.Format("Monday 02.01.2006"),
				"\nCurrent time:", currentTime.Format("15:04:05"))
			} else if cmd == "chg-username" { //! Commands for USER/ACCOUNT
				fmt.Print("Enter new username: ")
				fmt.Scanln(&user)
			} else if cmd == "whoami" {
				fmt.Println("User", user)
			} else {
				fmt.Println("Invalid command")
			}
		}
	}
}