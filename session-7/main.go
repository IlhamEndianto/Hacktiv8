package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	const SSH_ADDRESS = "localhost:22"
	const SSH_USERNAME = "demo"
	const SSH_PASSWORD = "password"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Dial Failed." + err.Error())
	}

	session, err := client.NewSession()
	if session != nil {
		defer session.Close()
	}

	if err != nil {
		log.Fatal("Failed session creation." + err.Error())
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// err = session.Run("dir")
	// if err != nil {
	// 	log.Fatal("Command execution error. " + err.Error())
	// }

	// err = session.Start("/bin/bash")
	// if err != nil {
	// 	log.Fatal("Error starting bash." + err.Error())
	// }

	// err = session.Wait()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var stdout, stderr bytes.Buffer
	// session.Stdout = &stdout
	// session.Stderr = &stderr

	// outputErr := stderr.String()
	// fmt.Println("=============== ERROR")
	// fmt.Println(strings.TrimSpace(outputErr))

	// outputString := stdout.String()
	// fmt.Println("=============== OUTPUT")
	// fmt.Println(strings.TrimSpace(outputString))

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client." + err.Error())
	}
	// err = session.Run("touch test-file.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fDestination, err := sftpClient.Create("fileDestination.txt")
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("test-file.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy source file into destination file." + err.Error())
	}

	log.Println("File copied.")
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}
