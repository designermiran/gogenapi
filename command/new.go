package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/designermiran/gogenapi/gogenapi"
	"github.com/tcnksm/go-gitconfig"
)

const (
	defaultDatabase = "sqlite"
	defaultVCS      = "github.com"
)

type NewCommand struct {
	Meta

	vcs       string
	username  string
	project   string
	namespace string
	database  string
}

func (c *NewCommand) Run(args []string) int {
	if err := c.parseArgs(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Fprintln(os.Stderr, "Error: $GOPATH is not found")
		return 1
	}

	return gogenapi.Skeleton(gopath, c.vcs, c.username, c.project, c.namespace, c.database)
}

func (c *NewCommand) parseArgs(args []string) error {
	flag := flag.NewFlagSet("gogenapi", flag.ContinueOnError)

	flag.StringVar(&c.vcs, "vcs", defaultVCS, "VCS")
	flag.StringVar(&c.username, "u", "", "Username")
	flag.StringVar(&c.username, "user", "", "Username")
	flag.StringVar(&c.namespace, "n", "", "Namespace of API")
	flag.StringVar(&c.namespace, "namespace", "", "Namespace of API")
	flag.StringVar(&c.database, "d", defaultDatabase, "Database engine [sqlite,postgres,mysql]")
	flag.StringVar(&c.database, "database", defaultDatabase, "Database engine [sqlite,postgres,mysql]")

	if err := flag.Parse(args); err != nil {
		return err
	}
	if 0 < flag.NArg() {
		c.project = flag.Arg(0)
	}

	if c.project == "" {
		return errors.New("Please specify project name.")
	}

	if c.username == "" {
		var err error
		c.username, err = gitconfig.GithubUser()
		if err != nil {
			c.username, err = gitconfig.Username()
			if err != nil || strings.Contains(c.username, " ") {
				return errors.New("Cannot find github username in `~/.gitconfig` file.\n" +
					"Please use -u option")
			}
		}
	}
	return nil
}

func (c *NewCommand) Synopsis() string {
	return "Generate boilerplate"
}

func (c *NewCommand) Help() string {
	helpText := `
Usage: gogenapi new [options] PROJECTNAME

  Generate go project and its boilerplate

Options:
  -database=database, -d     Database engine [sqlite,postgres,mysql] (default: sqlite)
  -namespace=namepace, -n    Namespace of API (default: "" (blank string))
  -user=name, -u             Username of VCS (default: username of github in .gitconfig)
  -vcs=name                  Version controll system to use (default: github.com)
`
	return strings.TrimSpace(helpText)
}
