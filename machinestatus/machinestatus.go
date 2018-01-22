/*
   主文件，另外的一些子属性放在其他文件里
   MachineStatus是一个 总状态，里面有很多的子状态
   */
package machinestatus

import (
		"fmt"
		"log"
		"os/exec"
		"bytes"
		)
/*
   机器状态结构体，包括cpu, 内存等
   */
type MachineStatus struct {
	Cpu CPUStatus;		//cpu状态
	Mem MemeryStatus;	//内存状态
	Disk DiskPartionInfo;	//硬盘状态
}

func (mac *MachineStatus) String() string {
	var res = "";

	res += mac.Cpu.String();
	res += mac.Mem.String();
	res += mac.Disk.String();

	return res;
}

//----------------MachineStatus end----------------//

//---------------公用方法----------------//

/*
   运行一条命令并返回字符串结果
   */
func exec_shell(s string) string {
	cmd:=exec.Command("/bin/bash", "-c", s);
	var out bytes.Buffer;

	cmd.Stdout=&out;
	err:=cmd.Run();
	if err!=nil {
		log.Fatal(err);
	}
	//fmt.Printf("%v\n", strings.FieldsFunc(out.String(), Split('\n'))[2]);
	return out.String();
}

//闭包方法返回一个字符串分割规则
func split(s rune) func(rune) bool {
	return func(c rune) bool {
		 return c == s; 
	}
}

/*
   生成一个machinestatus
   */
func GetMachineStatus() MachineStatus {
	fmt.Println("func:GetMachineStatus");
	var cpu, mem = getCPUAndMemery();
	var disk = getDiskPartionInfo();
	var mac = MachineStatus {Cpu:cpu, Mem:mem, Disk:disk};
	return mac;
}
//---------------公用方法 end----------------//


