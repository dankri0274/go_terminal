package main
import ( //! Import all the necessary packages
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
	var passwordC string

	var running bool = true
	var root bool = false

	var sym string = " $ "
	var sym_root string = " # "

	currentTime := time.Now()
	
	CLS()
	
	fmt.Print("Enter username: ")
	fmt.Scanln(&user)

	CLS() //! Clear terminal

	fmt.Printf("Password: %s", password "\033[8m")
	fmt.Print("\033[28m")

	CLS()

	fmt.Printf("Confirm assword: %s", passwordC "\033[8m")
	fmt.Print("\033[28m")

	if len(password) < 8 || password != passwordC {
		fmt.Println("Password must be atleast 8 characters")
		time.Sleep(time.Second * 2)
	} else if password == passwordC && len(password) >= 8 && len(passwordC) >= 8 {
		CLS()
		for running {
			if root {
				fmt.Print("root@", getIP(), sym_root)
				fmt.Scanln(&cmd)
			} else {
				fmt.Print(user, "@", getIP(), sym)
				fmt.Scanln(&cmd)
			}
			

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
			} else if cmd == "chg-password" {
				fmt.Print("Enter new password: ")
			} else if cmd == "whoami" {
				fmt.Println("User", user)
			} else if cmd == "help" {
				fmt.Println(
					"Check in the official documentation for full explanations",
					"\nContact Daniel for the documentation",
					"\n1. whoami",
					"\n2. ip",
					"\n3. time",
					"\n4. chg-password",
					"\n5. chg-username",
				)
			} else if cmd == "su" {
				if root != true {
					root = true
				} else {
					root = false
				}
			} else {
				fmt.Println("Invalid command")
			}
		}
	} else {
		fmt.Println("Something went wrong :(")
	}
}
