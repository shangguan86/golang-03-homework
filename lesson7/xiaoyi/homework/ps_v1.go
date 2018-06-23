package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func PidNum() []string {
	var b = true
	var PidData []string

	Dir, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`^[\d]{1,}`)
	for _, dir := range Dir {
		b = reg.MatchString(dir.Name())
		if dir.IsDir() && true == b {
			PidData = append(PidData, dir.Name())
		}
	}

	return PidData
}

func PidStat(Pid []string) {
	var totalVirtualMem int64
	var totalRSSMem int64
	var osPageSize = os.Getpagesize()
	var TotalScollectorMemoryMB uint64

	FileName, err := os.OpenFile("./tmp.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(FileName)
	defer FileName.Close()

	for _, pid := range Pid {
		file_status, err := os.Stat("/proc/" + pid)
		if err != nil {
			log.Fatal(err)
			continue
		}
		start_ts := file_status.ModTime().Unix()

		stats_file, err := ioutil.ReadFile("/proc/" + pid + "/stat")
		if err != nil {
			log.Fatal(err)
			continue
		}
		stats := strings.Fields(string(stats_file))
		if len(stats) < 24 {
			err = fmt.Errorf("stats too short")
			continue
		}

		us, err := strconv.ParseInt(stats[13], 10, 64)
		if err != nil {
			log.Fatal(err)
			continue
		}
		sys, err := strconv.ParseInt(stats[14], 10, 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		virtual, err := strconv.ParseInt(stats[22], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		totalVirtualMem += virtual

		rss, err := strconv.ParseInt(stats[23], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if pid == string(os.Getegid()) {
			TotalScollectorMemoryMB = uint64(rss) * uint64(osPageSize) / 1024 / 1024
		}

		totalRSSMem += rss

		fmt.Fprintf(w, "Pid:% -6v StartTime:% -6v us:% -4v sys:% -4v totalVirtualMem:% -6v TotalScollectorMemoryMB:% -4v \n", pid, start_ts, us, sys, totalVirtualMem, TotalScollectorMemoryMB)
	}
}

func ReadPsFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./tmp.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileline, err := ioutil.ReadAll(file)
	fmt.Fprintf(w, string(fileline))
}

func main() {
	PidStat(PidNum())

	http.HandleFunc("/", ReadPsFile)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}