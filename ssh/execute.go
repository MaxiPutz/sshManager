package ssh

import (
	"bytes"
	"fmt"
)

func Execute(sessionInfo SessionInfo, command string) (string, error) {
	session, err := sessionInfo.Client.NewSession()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	defer session.Close()
	var outBuf bytes.Buffer
	session.Stdout = &outBuf

	err = session.Run(command)

	return outBuf.String(), err

}
