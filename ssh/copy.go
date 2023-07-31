package ssh

import (
	"fmt"

	"github.com/povsister/scp"
)

func CopyFileFromRemote(sessionInfo SessionInfo, sourcePath string, destinationPath string) error {
	client := sessionInfo.Client

	conn, err := scp.NewClientFromExistingSSH(client, &scp.ClientOption{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	defer conn.Close()

	err = conn.CopyFileFromRemote(sourcePath, destinationPath, &scp.FileTransferOption{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return err
}

func CopyFileToRemote(sessionInfo SessionInfo, sourcePath string, destinationPath string) error {
	client := sessionInfo.Client
	conn, err := scp.NewClientFromExistingSSH(client, &scp.ClientOption{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	defer conn.Close()

	err = conn.CopyFileToRemote(sourcePath, destinationPath, &scp.FileTransferOption{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return err
}

func CopyDirFromRemote(sessionInfo SessionInfo, sourcePath string, destinationPath string) error {
	client := sessionInfo.Client

	conn, err := scp.NewClientFromExistingSSH(client, &scp.ClientOption{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	defer conn.Close()

	err = conn.CopyDirFromRemote(sourcePath, destinationPath, &scp.DirTransferOption{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return err
}

func CopyDirToRemote(sessionInfo SessionInfo, sourcePath string, destinationPath string) error {
	client := sessionInfo.Client

	conn, err := scp.NewClientFromExistingSSH(client, &scp.ClientOption{})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	defer conn.Close()

	err = conn.CopyDirToRemote(sourcePath, destinationPath, &scp.DirTransferOption{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return err
}
