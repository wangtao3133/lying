package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	goarch  string
	goos    string
	gocc    string
	gocxx   string
	cgo     string
	race    bool
	version string
)

// Usage: go run build.go [,options...] [args...]
// [options:]
//      -cc 		  CC参数
// 	    -cgo-enabled CGO_ENABLED
//      -cxx         CXX参数
// 	    -goarch      GOARCH (default "amd64")
//      -goos		 Use goos windows linux darwin (default current os)
// 	    -race		 Use race detector
//		-v           Use version x.xx.xx (default 1.0.0)
// [args:]
// 		build <name>     build a single main program
// 		builds <name...> build multiple main programs
// 		build-run <name> build and run a single main program
// 		test <,name>     run test
// 		clean			 Clearing the bin directory and reset
//
// Sample:
//   go run build.go build-run server   	            (编译并运行server主程序,文件格式以当前操作系统为主)
//   go run build.go -goos linux -v 2.2.3 build server (编译server主程序,并指定操作系统为linux,程序版本为2.2.3)
//   go run build.go clean 					            (清空bin目录并重置)
func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	flag.StringVar(&goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&gocc, "cc", "", "CC")
	flag.StringVar(&gocxx, "cxx", "", "CXX")
	flag.StringVar(&cgo, "cgo-enabled", "", "CGO_ENABLED")
	flag.BoolVar(&race, "race", race, "Use race detector")
	flag.StringVar(&version, "v", "1.0.0", "User version x.xx.xx(default 1.0.0)")
	flag.Parse()

	// 设置gopath
	ensureGoPath()

	if flag.NArg() == 0 {
		log.Println("Usage: go run build.go build")
		return
	}
	cmds := flag.Args()

	// 对应的命令
	args := make([]string, 0)
	for k, cmd := range cmds {
		if k == 0 {
			continue
		}
		args = append(args, cmd)
	}
	switch cmds[0] {
	case "build":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin name]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
	case "builds":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin names]")
			return
		}
		for _, binary := range args {
			build(binary, "./cmd/"+binary, []string{})
		}
	case "build-run":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin names]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
		runPrint("../bin/" + args[0] + "_" + goos)
	case "test":
		if len(args) == 0 {
			test("./..")
		} else {
			test(args[0])
		}
	case "clean":
		clean()
	default:
		log.Fatalf("Unknown command %q", cmds[0])
	}
}

// 配置gopath
func ensureGoPath() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	g := os.Getenv("GOPATH")
	var gopath string
	if g == "" {
		gopath = filepath.Clean(filepath.Join(pwd, "/../"))
	} else {
		gopath = filepath.Clean(filepath.Join(pwd, "/../")) + string(os.PathListSeparator) + g
	}
	log.Println("GOPATH is", gopath)
	os.Setenv("GOPATH", gopath)
}

// 运行命令
func runPrint(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// 配置文件
func clean() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rmr(filepath.Join(cwd, "../bin"))
	createDir("../bin")
	createDir("../bin/etc") // 创建etc目录
	createDir("../bin/log") // 创建log目录
	copyFiles(getFileNames("etc", "../bin/etc"))
}

// 创建目录
func createDir(dirName string) {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// 获取文件名
func getFileNames(src string, dst string) (srcFileNames []string, dstFileNames []string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			if strings.Count(dst, "/") >= 20 {
				log.Println("文件夹层级过深")
				return
			}
			createDir(dst + "/" + file.Name())
			childSrcFileNames, childDstFileNames := getFileNames(src+"/"+file.Name(), dst+"/"+file.Name())
			srcFileNames = append(srcFileNames, childSrcFileNames...)
			dstFileNames = append(dstFileNames, childDstFileNames...)
		} else {
			srcFileNames = append(srcFileNames, src+"/"+file.Name())
			dstFileNames = append(dstFileNames, dst+"/"+file.Name())
		}
	}
	return
}

// 批量复制文件
func copyFiles(srcNames, dstNames []string) {
	time1 := time.Now().UnixNano()
	for i := 0; i < len(srcNames); i++ {
		copyFile(srcNames[i], dstNames[i])
	}
	log.Println("耗时", time.Now().UnixNano()-time1)
}

// 复制文件
func copyFile(srcName, dstName string, wg ...*sync.WaitGroup) {
	if len(wg) > 0 {
		defer wg[0].Done()
	}
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err, 153)
		return
	}
	defer src.Close()
	_, err = os.Open(dstName)
	if err == nil {
		// 文件存在不修改
		return
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err, 358)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err, 364)
	}
}

// 循环删除目录
func rmr(paths ...string) {
	for _, path := range paths {
		log.Println("rm -r", path)
		os.RemoveAll(path)
	}
}

// 编译
func build(binaryName, pkg string, tags []string, paths ...string) {
	var binary string
	switch len(paths) {
	case 1:
		binary = paths[0] + binaryName
	case 0:
		binary = "../bin/" + binaryName
	default:
		log.Fatal("paths count err,there can only be one")
	}
	switch goos {
	case "windows":
		binary += "_windows.exe"
	case "linux":
		binary += "_linux"
	case "darwin":
		binary += "_darwin"
	}

	rmr(binary, binary+".md5")
	args := []string{"build", "-ldflags", ldflags()}
	if len(tags) > 0 {
		args = append(args, "-tags", strings.Join(tags, ","))
	}
	if race {
		args = append(args, "-race")
	}
	args = append(args, "-o", binary)
	args = append(args, pkg)
	setBuildEnv()

	runPrint("go", "version")
	runPrint("go", args...)

	// Create an md5 checksum of the binary, to be included in the archive for
	// automatic upgrades.
	err := md5File(binary)
	if err != nil {
		log.Fatal(err)
	}

	err = shaFile(binary)
	if err != nil {
		log.Fatal(err)
	}
}

// 编译参数
func ldflags() string {
	var b bytes.Buffer
	b.WriteString("-w")
	b.WriteString(fmt.Sprintf(" -X global.VERSION=%s", version))
	b.WriteString(fmt.Sprintf(" -X global.BUILDTIME=%d", time.Now().Unix()))
	return b.String()
}

// 设置编译环境参数
func setBuildEnv() {
	os.Setenv("GOOS", goos)
	if strings.HasPrefix(goarch, "armv") {
		os.Setenv("GOARCH", "arm")
		os.Setenv("GOARM", goarch[4:])
	} else {
		os.Setenv("GOARCH", goarch)
	}
	if goarch == "386" {
		os.Setenv("GO386", "387")
	}
	if cgo != "" {
		os.Setenv("CGO_ENABLED", cgo)
	}
	if gocc != "" {
		os.Setenv("CC", gocc)
	}
	if gocxx != "" {
		os.Setenv("CXX", gocxx)
	}
}

// md5加密文件
func md5File(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := md5.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".md5")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}

func shaFile(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := sha256.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".sha256")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}

// 运行测试
func test(pkg string) {
	setBuildEnv()
	runPrint("go", "test", "-short", "-timeout", "60s", pkg)
}
