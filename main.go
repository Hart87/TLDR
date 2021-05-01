package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/hart87/tldr/aggregator"
)

func main () {
	fmt.Println("T L D R ...")

	//Subcommands
	twitterCmd := flag.NewFlagSet("twitter", flag.ExitOnError)
    twitterSearchKey := twitterCmd.String("key", "", "Whaty you're searching for")
    twitterCmdEnable := twitterCmd.Bool("enable", false, "A boolean value to enable something guy.")
    

	redditCmd := flag.NewFlagSet("reddit", flag.ExitOnError)
    RedditSearchKey := redditCmd.String("key", "", "Whaty you're searching for")
    redditLevel := redditCmd.Int("level", 0, "how high or low a level is")

	if len(os.Args) < 2 {
        fmt.Println("expected 'twitter' or 'reddit' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
	case "twitter":
        twitterCmd.Parse(os.Args[2:])
        s := aggregator.GetElementSliceString("urls.twitter")
        aggregator.Iterate(s, *twitterSearchKey)
        fmt.Println("  enable:", *twitterCmdEnable)
        fmt.Println("  tail:", twitterCmd.Args())

    case "reddit":
        redditCmd.Parse(os.Args[2:])
        s := aggregator.GetElementSliceString("urls.reddit")
        aggregator.Iterate(s, *RedditSearchKey)
        fmt.Println("  level:", *redditLevel)
        fmt.Println("  tail:", redditCmd.Args())
        
		

    default:
        fmt.Println("expected 'twitter' or 'reddit' subcommands")
        os.Exit(1)
	}
}