package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fgimenez/cihealth/pkg/ghclient"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	opt := &options{
		Path:  "/tmp/test",
		Token: "",
	}

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, arguments []string) {
			if err := opt.Run(); err != nil {
				log.Fatalf("error: %v", err)
			}
		},
	}
	flag := cmd.Flags()

	flag.StringVar(&opt.Path, "path", opt.Path, "The directory to save results to.")
	flag.StringVar(&opt.Token, "gh-token", opt.Token, "OAuth2 token to interact with GitHub API.")

	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

type options struct {
	Path  string
	Token string
}

func (o *options) Run() error {
	if o.Token == "" {
		return fmt.Errorf("You need to specify the GitHub token path with --gh-token")
	}

	client := ghclient.New()

	results, err := client.Run()
	if err != nil {
		return err
	}

	d, err := yaml.Marshal(&results)
	if err != nil {
		return err
	}

	log.Printf("Writing output file %s", o.Path)
	err = ioutil.WriteFile(o.Path, d, 0644)
	if err != nil {
		return err
	}

	return nil
}
