# Poker Simulator
![image](https://github.com/user-attachments/assets/51402a97-7fdf-4676-9f1a-c7cc18fd2bee)
## Introduction
This is a poker simulator app, written in Go. The logic follows the rules of 2-Card Texas Hold'em. A random number of players between 2 - 8 are dealt 2 private cards each. There is no betting functionality. The community cards are displayed on the table. The best hand is rendered on the right hand side of the page, which remains blurred, unless the cursor is placed over that area of the page. The `New Game` button, positioned below the community cards, triggers the start of a new round.
## Go Version
`1.22.5`
## Packages
    "fmt"
    "html/template"
    "math/rand"
	"net/http"
	"os"
	"slices"
    "sort"
	"strconv"
	"strings"
	"time"	
## UML Diagram
![image](https://github.com/user-attachments/assets/7e9df7b8-1e20-441f-8e9d-478690602eab)
