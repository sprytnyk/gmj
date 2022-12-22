package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
)

const usage = `Usage of gmj:

	$ gmj -l
	Lists all available emojis, its codepoint, shortcode, and usage.

	$ gmj -c bug
	Gets an emoji codepoint by a provided short code.

	$ gmj -u bug
	Shows a use case by a provided shortcode.

	$ gmj -s dev
	Goes through a list of emojis and returns more feasiable codepoints to be
	used by a provided keyword.

`

func main() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
	}

	emojisList := flag.Bool("l", false, "")
	emojiCode := flag.String("c", "", "")
	emojiUsage := flag.String("u", "", "")
	lookupKeyword := flag.String("s", "", "")

	flag.Parse()

	if *emojisList {
		keys := make([]string, 0, len(Emojis))
		for key := range Emojis {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			fmt.Printf(
				"%s - :%s: - %s\n",
				Emojis[key]["codepoint"],
				color.RedString(key),
				color.CyanString(Emojis[key]["usage"]),
			)
		}
		fmt.Println(
			color.HiBlueString("Total number of Emojis: "),
			len(keys),
		)
		os.Exit(0)
	}

	if len(*lookupKeyword) > 0 {
		fmt.Println("WIP")
		os.Exit(0)
	}

	if len(*emojiCode) > 0 {
		value, key := Emojis[*emojiCode]
		if key {
			fmt.Println(value["codepoint"])
		} else {
			fmt.Printf(
				"Cannot get an emoji by a shortcode: %s.\n",
				color.HiRedString(*emojiCode),
			)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if len(*emojiUsage) > 0 {
		value, key := Emojis[*emojiUsage]
		if key {
			fmt.Printf("%s\n", color.YellowString(value["usage"]))
			os.Exit(0)
		} else {
			fmt.Printf(
				"Cannot get an emoji usage by a shortcode: %s.\n",
				color.HiRedString(*emojiCode),
			)
			os.Exit(1)
		}
	}

	flag.Usage()
	os.Exit(1)
}
