package main

import (
	"fmt"
	"time"
	"github.com/machinecontrol/machinestatus"
//	"log"
//	"os/exec"
//	"bytes"
//	"strings"
//	"strconv"
)

func main() {

	go func() {
		ticker := time.NewTicker(time.Second * 1);
		for _ = range ticker.C {
			fmt.Println(time.Now());
			var mac = machinestatus.GetMachineStatus(); 
			fmt.Println(mac.String());
		}
	}();

	var s string;
	fmt.Scanln(&s);
}

