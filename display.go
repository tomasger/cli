package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
)

func DisplayServerData (data []byte, sorting string) error {
	serverlist , err := ParseServerData(data)
	if err != nil {
		return ErrParse.Wrap(err, "Displaying server information failed")
	}
	if sorting != "" {
		Sort(serverlist, sorting)
	}
	fmt.Printf("NAME            DISTANCE\n")
	for _, server := range serverlist {
		fmt.Printf("%-20s%4d\n", server.Name, server.Distance)
	}
	fmt.Printf("Total number of servers: %v\n", len(serverlist))
	logrus.Debug("Data was displayed in command line window")
	return nil
}
func Sort(serverList []servers, sorting string) {
	logrus.Debug("Data sorting was called. Argument: ", sorting)
	switch sorting {
	case "best":
		sort.Slice(serverList, func(i, j int) bool {
			return serverList[i].Distance < serverList[j].Distance
		})
	case "alphabetical":
		sort.Slice(serverList, func(i, j int) bool {
			return serverList[i].Name < serverList[j].Name
		})
	default:
		logrus.Warning("Unexpected sorting parameter received. List will not be sorted")
	}
}
