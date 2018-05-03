package main

import (
	"math/rand"
	"strconv"
	"time"

	. "github.com/mobvisita/g_met"
)

//The following example is to demonstrate how to use GMet.
func main() {
	//Create GMet Instance with default settings (SeelogWriter and LtrFormatter)
	gmet := CreateGMetInstanceByDefault("../configs/g_met_seelog/g_met_log.xml")

	//To set the writer and formatter, you can refer to the following code
	//	writer, err := CreateMetWriterBySeeLog("../configs/g_met_seelog/g_met_log.xml")
	//	if err != nil {
	//		panic(err)
	//	}
	//	//create GMet instance by given the writer and the formatter
	//	gmet := CreateGMetInstance(writer, &LtrFormatter{})

	//Create a metric item of host IP, the name of the metric item is "HostAddr"
	addr, _ := IpAddress()
	for i := 0; i < 100; i++ {
		gmet.Send(addr, Metric("input_bytes", strconv.Itoa(rand.Intn(100))),
			Metric("output_bytes", strconv.Itoa(rand.Intn(100))))

		gmet.Flush() //in your real case, DON'T flush for each sending.
		//For seelog writer, the auto-flushing can be set in the log configuration
		time.Sleep(time.Second * 1)

	}
}
