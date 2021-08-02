// CGFオープンを目指す囲碁プログラム

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
MainLoop:
	for scanner.Scan() {
		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			fmt.Print("= \n\n")
		case "clear_board":
			fmt.Print("= \n\n")
		case "quit":
			break MainLoop
		case "protocol_version":
			fmt.Print("= 2\n\n")
		case "name":
			fmt.Print("= kifuwarabe\n\n")
		case "version":
			fmt.Print("= 0.0.1\n\n")
		case "list_commands":
			fmt.Print("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")
		case "komi":
			fmt.Print("= 6.5\n\n") // TODO コミ
		case "undo":
			fmt.Printf("= \n\n")
		// genmove b
		case "genmove":
			fmt.Printf("= %s\n\n", "pass")
		// play b a3
		// play w d4
		// play b pass
		// play w pass
		case "play":
			fmt.Print("= \n\n")
		default:
			fmt.Print("? unknown_command\n\n")
		}
	}

}
