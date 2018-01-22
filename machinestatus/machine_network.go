/*
   获取网络相关信息模块
   */
package machinestatus

import (
		"strings"
		"strconv"
		)

/*
   网络发送与接收数据量
   */
type NetworkStatus struct {
	Receive float64;	//接收速率KB/s
	Send float64;		//发送速率KB/s
}

func (net *NetworkStatus) String() string {
	var res = "";
	res = "网络\n";
	res += "receive : ";
	res += strconv. FormatFloat(net.Receive, 'f', 2, 64);
	res += " KB/s, ";
	res += "send : ";
	res += strconv. FormatFloat(net.Send, 'f', 2, 64);
	res += " KB/s\n";

	return res;
}

//---------------网络结构 end----------------//

//-----------静态方法区--------------------//
/*
   获取网络信息
   */
/*
1 Linux 3.10.0-514.16.1.el7.x86_64 (iZbp16yfh8mr7v2c43rbtpZ) 	01/22/2018 	_x86_64_	(1 CPU)
2 
3 07:57:07 PM     IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
4 07:57:08 PM      eth0      0.00      0.00      0.00      0.00      0.00      0.00      0.00
5 07:57:08 PM      eth1      1.00      1.00      0.05      0.10      0.00      0.00      0.00
6 07:57:08 PM        lo      2.00      2.00      0.10      0.10      0.00      0.00      0.00
7
8 Average:        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s
9 Average:         eth0      0.00      0.00      0.00      0.00      0.00      0.00      0.00
10 Average:         eth1      1.00      1.00      0.05      0.10      0.00      0.00      0.00
11 Average:           lo      2.00      2.00      0.10      0.10      0.00      0.00      0.00

*/
func getNetworkStatus() NetworkStatus {
	var str = exec_shell("sar -n DEV 1 1");
	var arr = strings.FieldsFunc(str, split('\n'));
	var res NetworkStatus = NetworkStatus{0.0, 0.0};
	
	var len = len(arr);
	for i:=4; i<len; i++ {
		if arr[i]=="" {
			break;
		}

		var fieldsarr = strings.Fields(arr[i]);
		var receive, _ = strconv.ParseFloat(fieldsarr[5], 64);
		var send, _ = strconv.ParseFloat(fieldsarr[6], 64);
		res.Receive += receive;
		res.Send += send;
	}

	return res;
}

//-----------静态方法区 end----------------//
