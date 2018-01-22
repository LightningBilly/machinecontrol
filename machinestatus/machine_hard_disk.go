package machinestatus

import(
		"strings"
		"strconv"
		)
/*
   单个分区信息
   */
type FileSystem struct {
	Name string;
	Total int;	//总容量，默认KB
	Used int;	//已经使用
	Unused int;	//剩余
	UsedPercent int;	//使用百分比 %
}

func (fs *FileSystem) String() string {
	var res = "";
	res += fs.Name;
	res += "\t" + strconv.Itoa(fs.Total) + " KB";
	res += "\t" + strconv.Itoa(fs.Used) + " KB";
	res += "\t" + strconv.Itoa(fs.Unused) + " KB";
	res += "\t" + strconv.Itoa(fs.UsedPercent) + "%";
	res += "\n";

	return res;
}

//----------单个分区信息 end-----------------//

/*
   系统硬盘信息
   */
type DiskPartionInfo struct {
	DiskPartion []FileSystem;	//多个分区信息
	Summary FileSystem;			//综合统计
}

func (dp *DiskPartionInfo) String() string {
	var res = "硬盘: \n";

	res += "名称\t总容量\t已使用\t未使用\t使用比\n";

	for _, fs := range dp.DiskPartion {
		res += fs.String();
	}
	res += dp.Summary.String();

	return res;
}
//-------------系统硬盘信息 end------------//

//-----------静态方法区--------------------//
/*
   获取系统硬盘信息
   */
/*
Filesystem     1K-blocks     Used Available Use% Mounted on
/dev/vda1       41152832 21254256  17785092  55% /
devtmpfs          931536        0    931536   0% /dev
tmpfs             941868        0    941868   0% /dev/shm
tmpfs             941868      556    941312   1% /run
tmpfs             941868        0    941868   0% /sys/fs/cgroup
tmpfs             188376        0    188376   0% /run/user/0
*/
func getDiskPartionInfo() DiskPartionInfo {
	var str = exec_shell("df -k");
	var arr = strings.FieldsFunc(str, split('\n'));
	var res DiskPartionInfo;
	//res.DiskPartion = append(res.DiskPartion, FileSystem{});
	res.Summary = FileSystem{"总计", 0, 0, 0, 0};
	
	var len = len(arr);
	for i:=1; i<len; i++ {
		var fieldsarr = strings.Fields(arr[i]);
		var name = fieldsarr[0];
		var total, _ = strconv.Atoi(fieldsarr[1]);
		var used, _ = strconv.Atoi(fieldsarr[2]);
		var unused, _ = strconv.Atoi(fieldsarr[3]);
		var percent = used * 100 / total;
		res.Summary.Total += total;
		res.Summary.Used += used;
		res.Summary.Unused += unused;

		res.DiskPartion = append(res.DiskPartion, 
				FileSystem{Name:name, Total:total, Used:used, Unused:unused, UsedPercent:percent});
	}

	res.Summary.UsedPercent = res.Summary.Used * 100 / res.Summary.Total;
	return res;
}


//-----------静态方法区 end----------------//
