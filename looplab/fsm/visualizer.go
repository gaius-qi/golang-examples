package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

const (
	PeerStatePending        = "Pending"
	PeerStateReceivedSmall  = "ReceivedSmall"
	PeerStateReceivedNormal = "ReceivedNormal"
	PeerStateRunning        = "Running"
	PeerStateBackToSource   = "BackToSource"
	PeerStateSucceeded      = "Succeeded"
	PeerStateFailed         = "Failed"
	PeerStateLeave          = "Leave"
)

const (
	PeerEventDownload                 = "Download"
	PeerEventRegisterSmall            = "RegisterSmall"
	PeerEventRegisterNormal           = "RegisterNormal"
	PeerEventDownloadFromBackToSource = "DownloadFromBackToSource"
	PeerEventDownloadSucceeded        = "DownloadSucceeded"
	PeerEventDownloadFailed           = "DownloadFailed"
	PeerEventLeave                    = "Leave"
)

const (
	TaskStatePending   = "Pending"
	TaskStateRunning   = "Running"
	TaskStateSucceeded = "Succeeded"
	TaskStateFailed    = "Failed"
)

const (
	TaskEventDownload          = "Download"
	TaskEventDownloadSucceeded = "DownloadSucceeded"
	TaskEventDownloadFailed    = "DownloadFailed"
)

func main() {
	peer := fsm.NewFSM(
		PeerStatePending,
		fsm.Events{
			{Name: PeerEventRegisterSmall, Src: []string{PeerStatePending}, Dst: PeerStateReceivedSmall},
			{Name: PeerEventRegisterNormal, Src: []string{PeerStatePending}, Dst: PeerStateReceivedNormal},
			{Name: PeerEventDownload, Src: []string{PeerStateReceivedSmall, PeerStateReceivedNormal}, Dst: PeerStateRunning},
			{Name: PeerEventDownloadFromBackToSource, Src: []string{PeerStateRunning}, Dst: PeerStateBackToSource},
			{Name: PeerEventDownloadSucceeded, Src: []string{PeerStateRunning, PeerStateBackToSource}, Dst: PeerStateSucceeded},
			{Name: PeerEventDownloadFailed, Src: []string{
				PeerStatePending, PeerStateReceivedSmall, PeerStateReceivedNormal,
				PeerStateRunning, PeerStateBackToSource, PeerStateSucceeded,
			}, Dst: PeerStateFailed},
			{Name: PeerEventLeave, Src: []string{PeerStateFailed, PeerStateSucceeded}, Dst: PeerEventLeave},
		},
		fsm.Callbacks{},
	)

	outputPeer, err := fsm.VisualizeWithType(peer, fsm.MermaidStateDiagram)
	if err != nil {
		panic(err)
	}

	fmt.Println("=========== peer ===========")
	fmt.Println(outputPeer)
	fmt.Println("============================")

	task := fsm.NewFSM(
		TaskStatePending,
		fsm.Events{
			{Name: TaskEventDownload, Src: []string{TaskStatePending, TaskStateFailed}, Dst: TaskStateRunning},
			{Name: TaskEventDownloadSucceeded, Src: []string{TaskStateRunning, TaskStateFailed}, Dst: TaskStateSucceeded},
			{Name: TaskEventDownloadFailed, Src: []string{TaskStateRunning}, Dst: TaskStateFailed},
		},
		fsm.Callbacks{},
	)

	outputTask, err := fsm.VisualizeWithType(task, fsm.MermaidStateDiagram)
	if err != nil {
		panic(err)
	}

	fmt.Println("=========== task ===========")
	fmt.Println(outputTask)
	fmt.Println("============================")
}
