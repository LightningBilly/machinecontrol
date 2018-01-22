/*
   主要放一些和CPU，内在相关的方法和属性
   */
package machinestatus

import(
		"strconv"
		"strings"
		//"fmt"
		)
/*
   CPU状态结构体
   */
type CPUStatus struct {
	Used int	//已经使用百分比
	Unused int	//未使用百分比
}

func (cpu *CPUStatus) String() string {
	var res = "";
	res = "CPU : \n";
	res += strconv.Itoa(cpu.Used) + "% 已使用, ";
	res += strconv.Itoa(cpu.Unused) + "% 空闲. ";
	res += "\n";
	return res;
}

//----------------CPUStatus end------------------------//



/*
   内在状态结构
   */
type MemeryStatus struct {
	Free int	//空闲未使用 默认单位KB
	Buff int	//缓冲区 默认单位KB
	Cache int	//缓存 默认单位KB
}

func (mem *MemeryStatus) String() string {
	var res = "";

	res = "内存: \n";
	res += "free : " + strconv.Itoa(mem.Free) + " KB\n";
	res += "buff : " + strconv.Itoa(mem.Buff) + " KB\n";
	res += "cache : " + strconv.Itoa(mem.Cache) + " KB\n";
	res += "总剩余 : " + strconv.Itoa(mem.Free + mem.Buff + mem.Cache) + " KB\n";

	return res;
}

//--------------MemeryStatus end-----------------------//



//--------------------静态方法区--------------------//

/*
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 0  1      2      3      4      5    6    7     8     9   10   11 12 13 14 15 16
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 4  0      0  91548   3104 305736    0    0   121    15    1    1  0  0 99  0  0
*/
func getCPUAndMemery() (CPUStatus, MemeryStatus) {

	res := exec_shell("vmstat");
	var arr = strings.Fields(strings.FieldsFunc(res, split('\n'))[2]);
	//fmt.Println(arr);
	var cpuUnused,_ = strconv.Atoi(arr[14]);
	var cpu = CPUStatus{Used: 100-cpuUnused, Unused:cpuUnused};

	var free, _ =  strconv.Atoi(arr[3]);
	var buff, _ = strconv.Atoi(arr[4]);
	var cache, _ = strconv.Atoi(arr[5]);

	var mem = MemeryStatus{Free:free, Buff:buff, Cache:cache};

	return cpu, mem;
}

//--------------------静态方法区 end--------------------//

