package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

//加息
func Parse(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndClassExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//如果用户没有提供-classpath/-cp选项,则使用当前目录作为用
//户类路径。ReadClass()方法依次从启动类路径、扩展类路径和用户
//类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	data, entry, err := self.bootClassPath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	data, entry, err = self.extClassPath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	return self.userClassPath.readClass(className)

}

func (self *Classpath) String() string {
	return self.userClassPath.String()
}

func (self *Classpath) parseBootAndClassExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)
	// jre/ext/*
	jreExtPath := filepath.Join(jreDir, "ext", "*")
	self.extClassPath = newWildcardEntry(jreExtPath)
}

// 1.先在jreOption里找jre目录,找到返回
// 2.如果没有则在当前目录找jre目录,找到则返回
// 3.如果没有则取系统环境变量,找到则返回
// 4.如果都没有,报错
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("can't find jre folder")
}

//判断目录是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return false
		}
	}

	return true
}

//parseUserClasspath()方法的代码相对简单一些,如下:
func (self *Classpath) parseUserClasspath(classpathOption string) {
	if classpathOption == "" {
		classpathOption = "."
	}
	self.userClassPath = newEntry(classpathOption)

}
