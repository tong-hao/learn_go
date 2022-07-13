package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func print(reader io.Reader) {
	log.Println("Enter print")
	for {
		p := make([]byte, 4)
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
	log.Println("Leave print")
}

func sshAgent() (ssh.AuthMethod, func(), error) {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, func() {}, err
	}
	method := ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	return method, func() { _ = sshAgent.Close() }, nil
}

func exec(session *ssh.Session, cmd string) {
	log.Println("cmd:", cmd)
	result, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to run command, Err:%s", err.Error())
	}
	fmt.Println("log:", string(result))

}

func execAsync(session *ssh.Session, cmd string) {
	var wg sync.WaitGroup
	consumeStream := func(r io.Reader) {
		log.Println("begin: consume ", r)
		info := make(chan string, 3)
		errc := make(chan error, 3)
		scan := bufio.NewScanner(r)
		scan.Split(bufio.ScanLines)
		for scan.Scan() {
			info <- scan.Text()
		}
		if err := scan.Err(); err != nil {
			errc <- err
		}

		defer wg.Done()
		log.Println(info)
		log.Println(errc)
		log.Println("end: consume ", r)
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		fmt.Printf("opening stderr: %v", err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Printf("opening stdout: %v", err)
	}

	wg.Add(1)
	go consumeStream(stderr)
	go consumeStream(stdout)

	log.Println("begin: ls -a ")
	if err := session.Start("ls -a"); err != nil {
		fmt.Println(err)
	}
	log.Println("end: ls -a ")
	wg.Add(1)

	go func() {
		if err := session.Wait(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		wg.Wait()
	}()

}

func main() {
	auth, done, err := sshAgent()
	if err != nil {
		log.Fatalf("no ssh agent: %v", err)
	}
	defer done()

	//var hostKey ssh.PublicKey
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", "192.168.8.88:22", &ssh.ClientConfig{
		User:            "caton.hao",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{auth},
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	//exec(session, "ls -a")
	//go func() {time.Sleep(1000); log.Println("kill session") ;session.Close()}()
	//execAsync(session, "ls -a")
	cmd := "pstree -p 5664"
	exec(session, cmd)
	//go func() {time.Sleep(1000); log.Println("kill session") ;session.Close()}()
	defer log.Println("over.")

	client.Close()
	client.Close()
}
