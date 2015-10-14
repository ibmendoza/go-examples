//https://github.com/blacklabeldata/sshh/blob/master/examples/shell/main.go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/blacklabeldata/sshh"
	log "github.com/mgutz/logxi/v1"
	"golang.org/x/crypto/ssh"
)

var privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDjzAhRGLLcnQhs7Xe/2TrbjpHOkeBwVfmI0z+mZot87AXyIVcr
+OepPl/8UekPb352bz3zAwn2x5zCT/hW+1CBwp6fqhAvlxlYFEYr40L2dYKMmZyT
3kq18P3fTmAIKyXv7XOtVXiNLHc0Ai+3aN4J+yHKwbf42nNU3Qb1NRp9KQIDAQAB
AoGANgZyxoD8EpRvph3fs7FaYy356KryNtI9HzUyuE1DsbnsYxODMBuVHa98ZkQq
6Q1BSedyIstKtqt6wx7iQAbUfa9VxYht2DnxJDG7AhbQS1jd8ifSPCyhsp7HqCL5
pPbJBoW2M2qVL95+TMaZKYDDQcpFIHsEzJ/6lnWatGdBxfECQQDwv+cFSe5i8hqU
5BmLH3131ez5jO4yCziQxNwZaEavDXPDsqeKl/8Oj9EOcVyysyOLR9z7NzOCV2wX
8u0hpO69AkEA8joVv2rZdb+83Zc1UF/qnihMt4ZqYafPMXEtl2YTZtDmQOZG0kMw
a/iPjkUt/t8+CNR/Z5RLUYA5NVJSlsI03QJBANUZaEo8KLCYkILebOXCl/Ks/zfd
UTIm0IkEV7Z9oKNuitvclYSOCgw/rNLV8TGUc4/jqm0LbaKf82Q3eULglRkCQBsi
4rjVEZOdbV0tyW09sZ0SSrXsuxJBqHaThVYGu3mzQXhX0+tOV6hg6kQ3/9Uj0WFP
3Q4PkPiKct5EYLg+/YkCQCpHiRgfbESG2J/eYtTdyDvm+r0m0pc4vitqKsRGjd2u
LZxh0eGWnXXd+Os/wOVMSzkAWuzc4VTxMUnk/yf13IA=
-----END RSA PRIVATE KEY-----
`

func main() {

	// Create logger
	writer := log.NewConcurrentWriter(os.Stdout)
	logger := log.NewLogger(writer, "sshh")

	// Get private key
	privateKey, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		logger.Warn("Private key could not be parsed", "error", err.Error())
	}

	// Setup server config
	config := sshh.Config{
		Deadline: time.Second,
		Logger:   logger,
		Bind:     ":9022",
		Handlers: map[string]sshh.SSHHandler{
			"session": NewShellHandler(logger),
		},
		PrivateKey: privateKey,
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (perm *ssh.Permissions, err error) {
			if conn.User() == "admin" && string(password) == "password" {

				// Add username to permissions
				perm = &ssh.Permissions{
					Extensions: map[string]string{
						"username": conn.User(),
					},
				}
			} else {
				err = fmt.Errorf("Invalid username or password")
			}
			return
		},
		AuthLogCallback: func(conn ssh.ConnMetadata, method string, err error) {
			if err == nil {
				logger.Info("Successful login", "user", conn.User(), "method", method)
			}
		},
		// PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (perm *ssh.Permissions, err error) {
		// 	return nil, fmt.Errorf("Unauthorized")
		// },
	}

	// Create SSH server
	sshServer, err := sshh.NewSSHServer(&config)
	if err != nil {
		logger.Error("SSH Server could not be configured", "error", err.Error())
		return
	}

	// Start servers
	sshServer.Start()

	// Handle signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	// Wait for signal
	logger.Info("Ready to serve requests")

	// Block until signal is received
	<-sig

	// Stop listening for signals and close channel
	signal.Stop(sig)
	close(sig)

	// Shut down SSH server
	logger.Info("Shutting down servers.")
	sshServer.Stop()
}
