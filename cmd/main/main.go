package main

import (
	"fmt"

	"github.com/TanguturiKishore/broadpeakgo"
)

const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImJyb2FkcGVha2FwaWtleSIsInRlbmFudElkIjoxMzg0LCJpYXQiOjE2NzUwODUyMzJ9.CkDFu8DYKN_vtBuQx6iEQes0LA1Au8gFnLSp595m35E"

func main() {
	broadpeak := broadpeakgo.GetBroadPeak(apiKey)

	fmt.Println(broadpeak)

	broadpeak.GetAllSources()
}
