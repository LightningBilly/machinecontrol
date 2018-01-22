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
			var mac = machinestatus.GetMachineStatus(); 
			fmt.Println(time.Now());
			fmt.Println(mac.String());
		}
	}();

	var s string;
	fmt.Scanln(&s);
}

