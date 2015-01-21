package main

//"github.com/martini-contrib/binding"
import (
	"fmt"
	"github.com/chai2010/gettext-go/gettext"
	"github.com/codegangsta/cli"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yetist/genpass"
	"github.com/yetist/middleware/i18n"
	"html/template"
	"os"
	"strconv"
)

const (
	PkgName    = "genpass"
	PkgVersion = "0.1"
)

var webFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "port, p",
		Value: 8080,
		Usage: __("Http server port."),
	},
}

var cmdFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "user, u",
		Value: "username",
		Usage: __("User name about website."),
	},

	cli.StringFlag{
		Name:  "domain, d",
		Value: "baidu.com",
		Usage: __("The domain about your password used for."),
	},
	cli.StringFlag{
		Name:  "flag, f",
		Value: "alpha",
		Usage: __("Which chars should include in password, valid option is:\n\tupper|lower|alpha|digit|punct|xdigit|alpha|alnum|graph"),
	},
	cli.StringFlag{
		Name:  "extra, e",
		Value: "",
		Usage: __("Which extra chars can used for part of password."),
	},
	cli.StringFlag{
		Name:  "method, m",
		Value: "sha256",
		Usage: __("Which method should use, valid options is: md5|sha1|sha256|sha512"),
	},
	cli.IntFlag{
		Name:  "reversion, r",
		Value: 0,
		Usage: __("Password version, for update password."),
	},
	cli.IntFlag{
		Name:  "length, l",
		Value: 8,
		Usage: __("Password length, default is 8."),
	},
}

var CmdGen = cli.Command{
	Name:        "gen",
	Usage:       __("Generate password now"),
	Action:      runGen,
	Flags:       cmdFlags,
	Description: __("Generate your password."),
}

var CmdServer = cli.Command{
	Name:   "server",
	Usage:  __("Run http server to user web generate password."),
	Action: runServer,
	Flags:  webFlags,
}

func runServer(c *cli.Context) {
	port := c.Int("port")
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Funcs:     []template.FuncMap{{"__": __}},
	}))
	m.Use(i18n.I18n(i18n.Options{
		Domain:    PkgName,
		Directory: "locale",
		Parameter: "lang",
		Inited:    true,
	}))
	m.Use(martini.Static("static", martini.StaticOptions{
		Prefix: "static",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})
	m.Post("/", func(r render.Render) string {
		return "ok"
	})
	martini.Env = martini.Prod
	m.RunOnAddr(":" + strconv.Itoa(port))
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
	fmt.Printf(__("Password: %s\n"), p)
}

func __(msgid string) string {
	return gettext.PGettext("", msgid)
}

func main() {
	gettext.BindTextdomain(PkgName, "locale", nil)
	gettext.Textdomain(PkgName)
	app := cli.NewApp()
	app.Name = PkgName
	app.Usage = __("Generate Password Service")
	app.Version = PkgVersion
	app.Commands = []cli.Command{
		CmdServer,
		CmdGen,
	}
	app.Run(os.Args)
}
