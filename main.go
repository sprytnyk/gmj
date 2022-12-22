package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
)

const usage = `Usage of gmj:

	$ gmj -l
	Lists all available emojis, its codepoint, shortcode, and usage.

	$ gmj -c bug
	Gets an emoji codepoint by a provided shortcode.

	$ gmj -u bug
	Shows a use case by a provided shortcode.

	$ gmj -s dev
	Goes through a list of emojis and returns more feasiable emojis to be
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

	keys := make([]string, 0, len(Emojis))
	for key := range Emojis {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if *emojisList {
		fmt.Println(color.HiBlueString("Codepoint - Shortcode - Usage"))
		for _, key := range keys {
			fmt.Printf(
				"%s - :%s: - %s\n",
				Emojis[key]["codepoint"],
				color.RedString(key),
				color.CyanString(Emojis[key]["usage"]),
			)
		}
		fmt.Println(
			color.HiBlueString("Total number of emojis: "),
			len(keys),
		)
		os.Exit(0)
	}

	if len(*lookupKeyword) > 0 {
		sum := 0
		fmt.Println(color.HiBlueString("Codepoint - Shortcode - Usage"))
		lookupKeyword := strings.ToLower(*lookupKeyword)
		for _, key := range keys {
			usage := strings.ToLower(Emojis[key]["usage"])
			key := strings.ToLower(key)
			if strings.Contains(usage, lookupKeyword) ||
				strings.Contains(key, lookupKeyword) {
				fmt.Printf(
					"%s - :%s: - %s\n",
					Emojis[key]["codepoint"],
					color.RedString(key),
					color.CyanString(Emojis[key]["usage"]),
				)
				sum += 1
			}
		}
		if sum == 0 {
			fmt.Printf(
				"None of emojis found by a shortcode %s\n",
				color.HiYellowString(lookupKeyword),
			)
		}
		fmt.Println(color.HiBlueString("Total number of emojis: "), sum)
		os.Exit(0)
	}

	if len(*emojiCode) > 0 {
		emojiCode := strings.ToLower(*emojiCode)
		value, key := Emojis[emojiCode]
		if key {
			fmt.Println(value["codepoint"])
		} else {
			fmt.Printf(
				"Cannot get an emoji by a shortcode: %s\n",
				color.HiRedString(emojiCode),
			)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if len(*emojiUsage) > 0 {
		emojiUsage := strings.ToLower(*emojiUsage)
		value, key := Emojis[emojiUsage]
		if key {
			fmt.Printf("%s\n", color.HiBlueString(value["usage"]))
			os.Exit(0)
		} else {
			fmt.Printf(
				"Cannot get an emoji usage by a shortcode: %s\n",
				color.HiRedString(emojiUsage),
			)
			os.Exit(1)
		}
	}

	flag.Usage()
	os.Exit(1)
}
