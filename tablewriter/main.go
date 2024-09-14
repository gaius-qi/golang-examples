package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Blob URL", "Hostname", "IP", "Cluster ID", "State", "Description"})

	data := [][]string{
		{"https://blob1", "host1", "1", "cluster1", "active", "blob1"},
		{"https://blob2", "host2", "2", "cluster2", "active", "blob2"},
		{"https://blob1", "host1", "1", "cluster3", "active", "blob3"},
		{"https://blob2", "host2", "2", "cluster4", "active", "blob4"},
		{"https://blob1", "host1", "1", "cluster4", "active", "blob4"},
		{"https://blob2", "host2", "2", "cluster4", "active", "blob4"},
	}
	for _, v := range data {
		table.Append(v)
	}

	table.SetAutoMergeCells(true)
	table.Render()
}
