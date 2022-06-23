package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	name      string
	alamat    string
	pekerjaan string
}

func main() {
	var coworker = []struct {
		biodata  Biodata
		motivasi string
	}{
		{biodata: Biodata{name: "Kristian", alamat: "Jaksel", pekerjaan: "Database Developer"}, motivasi: "Mau jago"},
		{biodata: Biodata{name: "Bakrie", alamat: "Tangerang", pekerjaan: ".NET Developer"}, motivasi: "Mau jago"},
		{biodata: Biodata{name: "Kevin", alamat: "Jakbar", pekerjaan: ".NET Developer"}, motivasi: "Mau jago"},
		{biodata: Biodata{name: "Chandra", alamat: "Jakbar", pekerjaan: ".NET Developer"}, motivasi: "Mau jago"},
	}
	t := os.Args[1]
	idx, err := strconv.ParseInt(t, 0, 16)
	if err == nil {
		fmt.Println("==========================================")
	}
	fmt.Printf("Nama %8s %s\n", ":", coworker[idx-1].biodata.name)
	fmt.Printf("Alamat %6s %s\n", ":", coworker[idx-1].biodata.alamat)
	fmt.Printf("Pekerjaan %3s %s\n", ":", coworker[idx-1].biodata.pekerjaan)
	fmt.Printf("Alasan %6s %s\n", ":", coworker[idx-1].motivasi)

}
