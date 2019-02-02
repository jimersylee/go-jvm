package main

import "fmt"

//入口函数
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}

}

/**
开始虚拟机
*/
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s,class:%s,args:%v", cmd.cpOption, cmd.class, cmd.args)
}
