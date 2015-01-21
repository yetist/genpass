package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/yetist/genpass"
	"os"
)

var webFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "port, p",
		Value: 8080,
		Usage: "language for the greeting",
	},
}

var cmdFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "user, u",
		Value: "username",
		Usage: "User name about website.",
	},

	cli.StringFlag{
		Name:  "domain, d",
		Value: "baidu.com",
		Usage: "The domain about your password used for.",
	},
	cli.StringFlag{
		Name:  "flag, f",
		Value: "alpha",
		Usage: `Which chars should include in password, valid option is:
	upper|lower|alpha|digit|punct|xdigit|alpha|alnum|graph`,
	},
	cli.StringFlag{
		Name:  "extra, e",
		Value: "",
		Usage: "Which extra chars can used for part of password.",
	},
	cli.StringFlag{
		Name:  "method, m",
		Value: "sha256",
		Usage: "Which method should use, valid options is: md5|sha1|sha256|sha512",
	},
	cli.IntFlag{
		Name:  "reversion, r",
		Value: 0,
		Usage: "Password version, for update password.",
	},
	cli.IntFlag{
		Name:  "length, l",
		Value: 8,
		Usage: "Password length, default is 8.",
	},
}

var CmdGen = cli.Command{
	Name:        "gen",
	Usage:       "Generate password now",
	Action:      runGen,
	Flags:       cmdFlags,
	Description: "Generate your password.",
}

var CmdServer = cli.Command{
	Name:   "server",
	Usage:  "Run http server to user web generate password.",
	Action: runServer,
	Flags:  webFlags,
}

func runServer(c *cli.Context) {
	port := c.Int("port")
	fmt.Printf("run http server on %v ...\n", port)
}

func genFlag(flag string) int {
	flagtable := map[string]int{
		"alnum":  genpass.CharAlnum,
		"alpha":  genpass.CharAlpha,
		"digit":  genpass.CharDigit,
		"graph":  genpass.CharGraph,
		"lower":  genpass.CharLower,
		"punct":  genpass.CharPunct,
		"upper":  genpass.CharUpper,
		"xdigit": genpass.CharXdigit,
	}
	return flagtable[flag]
}

func runGen(c *cli.Context) {
	opt := genpass.Options{
		Primary:     c.String("user"),
		Description: c.String("domain"),
		ExtraChars:  c.String("extra"),
		Method:      c.String("method"),
		Flag:        genFlag(c.String("flag")),
		Reversion:   c.Int("reversion"),
		Length:      c.Int("length"),
	}
	p := genpass.Gen(opt)
	fmt.Printf("Password: %s\n", p)
}

func main() {
	app := cli.NewApp()
	app.Name = "genpass"
	app.Usage = "Generate Password Service"
	app.Version = "1.0"
	app.Commands = []cli.Command{
		CmdServer,
		CmdGen,
	}
	app.Run(os.Args)
}
