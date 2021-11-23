package main
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {						// FOR LINUX
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls", "/c", "cls")		// FOR WINDOWS
		cmd.Stdout = os.Stdout
		cmd.Run()
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is not supported!")
	}
}

func main() {
	var cmd string
	var user string
	
	for cmd != "exit()" {
		fmt.Print("Enter username: ")
		fmt.Scanln(&user)

		CallClear() // Clear terminal
		
		fmt.Print("Enter command: ")
		fmt.Scanln(&cmd)
		if cmd == "whoami" {
			fmt.Println("User", user)
		} else {
			fmt.Println("Only command available at the moment is \"whoami\"")
		}
	}
}
