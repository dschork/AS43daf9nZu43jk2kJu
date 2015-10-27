// main
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func main() {

	proz_name := "wish"
	//	proz_name2 := "xmaple"

	startProcess(proz_name)

}

func StartProcessTest(*testing.T) {
	name := ""
	process_cmd := exec.Command(name)
	process := process_cmd.Start()
	fmt.Println(process)
}

func startProcess(name string) int {

	// fmt.Println("Prozess: ", name)

	// process_cmd := exec.Command(name)
	process_cmd := exec.Command(name)
	process_cmd.Stdin = strings.NewReader("xxx")

	var out bytes.Buffer
	process_cmd.Stdout = &out

	process := process_cmd.Run()
	if process != nil {
		fmt.Println("An Error occured!")
		return 1
	} else {
		fmt.Println("Process runs!")
		fmt.Println("Output: ", out)
		return 0
	}

}

//func get_process_id(name string) int {

//	ident := ps aux | grep name | awk '{print $11 " " $2}'
//	args := []string{ "| grep name",
//					 "| awk '{print $11 " " $2}'"
//					}

//	// ident := "ps aux"

//	id_cmd := exec.Command(ident, args)
//	id := id_cmd.Start()
//	idReader, err := id_cmd.StdoutPipe()

//	if err != nil {
//		fmt.Println("An Error occured while Stdout Pipe!")
//	}

//	if id != nil {
//		// fmt.Println("An Error occured!")
//		return 1
//	} else {
//		fmt.Println("Process Id:", id)
//		fmt.Println(idReader)
//		return 0
//	}
//}
