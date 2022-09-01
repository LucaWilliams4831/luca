package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/nikola43/panda141035/crypto"
)

type UploadCommand struct {
	args *Args
}

const (
	uploadCommand     = "upload"
	uploadDescription = "upload file to chain"
)

func newUploadCommand() Command {
	return Command{
		Name:        uploadCommand,
		Description: uploadDescription,
		Exec:        &UploadCommand{},
	}
}

func (c *UploadCommand) ExecCommand(ctx context.Context, args []string) error {
	c.args = &Args{args}
	if len(args) == 0 {
		return ErrorFromString(fmt.Sprintf("file not found"))
	}

	fi, err := os.Stat(args[0])
	if err != nil {
		return ErrorFromString(fmt.Sprintf("file not found"))
	}
	// get the size
	size := fi.Size()
	if size == 0 {
		return ErrorFromString(fmt.Sprintf("file corupted"))
	}

	if len(args) < 2 {
		return ErrorFromString(fmt.Sprintf("server ip"))
	}

	isValid := checkIPAddress(args[1])
	if !isValid {
		return ErrorFromString(fmt.Sprintf("ip invalid format"))
	}

	if len(args) < 3 {
		return ErrorFromString(fmt.Sprintf("server port"))
	}

	dat, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		return ErrorFromString(err.Error())
	}

	sCipher := &crypto.GCM{
		Passphrase: "6368616e676520746869732070617373776f726420746f206120736563726574",
	}
	sCipher.Init()

	b, err := sCipher.Encrypt([]byte(dat))
	if err != nil {
		fmt.Println("Encrypt")
		panic(err)
	}

	conn, err := net.Dial("udp", args[1]+":"+args[2])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if conn == nil {
		return ErrorFromString("Invalid connection")
	}
	bytes, err := fmt.Fprintf(conn, string(b))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(bytes)

	//return ErrorFromString(fmt.Sprintf("%s: invalid subcommand passed", uploadCommand))
	return nil
}

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		fmt.Printf("IP Address: %s - Invalid\n", ip)
		return false
	} else {
		fmt.Printf("IP Address: %s - Valid\n", ip)
		return true
	}
}
