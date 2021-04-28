package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/hart87/tldr/aggregator"
)

func main () {
	fmt.Println("Greetings from TLDR...")

	//Subcommands
	twitterCmd := flag.NewFlagSet("twitter", flag.ExitOnError)
    twitterCmdEnable := twitterCmd.Bool("enable", false, "A boolean value to enable something guy.")
    twitterSearchKey := twitterCmd.String("name", "", "the name of something")

	redditCmd := flag.NewFlagSet("reddit", flag.ExitOnError)
    redditLevel := redditCmd.Int("level", 0, "how high or low a level is")

	if len(os.Args) < 2 {
        fmt.Println("expected 'twitter' or 'reddit' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
	case "twitter":
        twitterCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'twitter'")
        fmt.Println("  enable:", *twitterCmdEnable)
        fmt.Println("  name:", *twitterSearchKey)
        fmt.Println("  tail:", twitterCmd.Args())
        s := aggregator.GetElementSliceString("urls.twitter")
        t := aggregator.GetElementString("app")
        fmt.Println(s)
        fmt.Println(t)
    case "reddit":
        redditCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'reddit'")
        fmt.Println("  level:", *redditLevel)
        fmt.Println("  tail:", redditCmd.Args())
        s := aggregator.GetElementSliceString("urls.reddit")
		fmt.Println(s)

    default:
        fmt.Println("expected 'twitter' or 'reddit' subcommands")
        os.Exit(1)
	}
}