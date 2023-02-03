package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	bilibili "github.com/misssonder/bilibili/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cookieDir  = os.Getenv("HOME")
	cookieFile = ".bilibili_cookie.txt"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login bilibili through qrcode (default is $HOME/.bilibili_cookie.txt).",
	Args:  cobra.ExactArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		exitOnError(createLoginDir())
	},
	Run: func(cmd *cobra.Command, args []string) {
		exitOnError(login())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func isLogin() bool {
	cookie, err := readCookieFromFile()
	if err != nil {
		return false
	}
	client.SetCookie(cookie)
	info, err := client.NavInfo()
	if err != nil {
		return false
	}
	return info.Data.IsLogin
}

func login() error {
	if !isLogin() {
		logrus.Info("Please login")
		responses, err := client.LoginWithQrCode(os.Stdout)
		if err != nil {
			return err
		}
		for resp := range responses {
			switch resp.LoginStatus {
			case bilibili.LoginSuccess:
				client.SetCookie(resp.Cookie)
				if err = saveCookieFile(resp.Cookie); err != nil {
					return err
				}
				return nil
			case bilibili.LoginExpired:
				return fmt.Errorf("login qrcode expired")
			default:
				continue
			}
		}
	}
	return nil
}

func saveCookieFile(cookie []string) error {
	file, err := os.OpenFile(path.Join(cookieDir, cookieFile), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(strings.Join(cookie, "\n")))
	return err
}

func readCookieFromFile() ([]string, error) {
	file, err := os.Open(path.Join(cookieDir, cookieFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func createLoginDir() error {
	_, err := os.Stat(cookieDir)
	switch err {
	case nil:
	default:
		if os.IsNotExist(err) {
			if err := os.MkdirAll(cookieDir, 0o755); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
