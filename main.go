package main

import (
	"main/API/GroopieTracker"
	Start "main/Front"
)

func main() {
	GroopieTracker.InitApi()
	Start.Start()
}
