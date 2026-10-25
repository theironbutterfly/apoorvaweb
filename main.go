package main

import "fmt"

func reboot(serverID int) {
    fmt.Println("Rebooting server", serverID)
}

func main() {
    for i := 1; i <= 3; i++ {
        go reboot(i)
    }

    fmt.Println("All servers rebooted")
}