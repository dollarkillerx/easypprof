/**
 * @Author: DollarKillerX
 * @Description: pprof.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:55 2019/12/31
 */
package easypprof

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func BasicPprofTime(timeout time.Duration) {
	cpuf, e := os.Create("cpu_profile")
	if e != nil {
		log.Fatalln(e)
	}
	pprof.StartCPUProfile(cpuf)

	defer pprof.StopCPUProfile()

	time.Sleep(timeout)

	log.Println("关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||  CPU")

	memf, err := os.Create("mem_profile")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	memf.Close()

	log.Println("关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析||| MEMF")
}
