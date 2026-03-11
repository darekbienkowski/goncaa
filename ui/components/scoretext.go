/*
Credit to @dylantientcheu for this implementation originally from nbacli
https://github.com/dylantientcheu/nbacli/blob/master/ui/gameboard/scoretext/scoretext.go
*/
package components

import (
	"fmt"
	"os"
	"strings"

	"github.com/darekbienkowski/goncaa/api/ncaa"
	"github.com/darekbienkowski/goncaa/ui/constants"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	subtle         = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#212121"}
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#5b1b7b", Dark: "#5b1b7b"}).
			Padding(1, 1).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	scoreTextStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginTop(1)

	teamNameStyle = lipgloss.NewStyle().Margin(0, 1)
	infoStyle     = lipgloss.NewStyle().Margin(0, 0)
)

// gotten from https://fsymbols.com/generators/tarty/

var scoreTextFont = map[int]string{
	420: `      
      
█████╗
╚════╝
      
      `,
	0: ` █████╗ 
██╔══██╗
██║  ██║
██║  ██║
╚█████╔╝
 ╚════╝ `,
	1: `  ███╗  
 ████║  
██╔██║  
╚═╝██║  
███████╗
╚══════╝`,
	2: `██████╗ 
╚════██╗
  ███╔═╝
██╔══╝  
███████╗
╚══════╝`,
	3: `██████╗ 
╚════██╗
 █████╔╝
 ╚═══██╗
██████╔╝
╚═════╝ `,
	4: `  ██╗██╗
 ██╔╝██║
██╔╝ ██║
███████║
╚════██║
     ╚═╝`,
	5: `███████╗
██╔════╝
██████╗ 
╚════██╗
██████╔╝
╚═════╝ `,
	6: ` █████╗ 
██╔═══╝ 
██████╗ 
██╔══██╗
╚█████╔╝
 ╚════╝ `,
	7: `███████╗
╚════██║
    ██╔╝
   ██╔╝ 
  ██╔╝  
  ╚═╝   `,
	8: ` █████╗ 
██╔══██╗
╚█████╔╝
██╔══██╗
╚█████╔╝
 ╚════╝ `,
	9: ` █████╗ 
██╔══██╗
╚██████║
 ╚═══██║
 █████╔╝
 ╚════╝ `,
}

func RenderScoreText(boxscore *ncaa.GameinfoNCAA) string {
	away, home := boxscore.Teams[0], boxscore.Teams[1]

	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	hPad, _ := constants.DocStyle.GetFrameSize()

	width = width - hPad

	doc := strings.Builder{}

	{
		// game board
		gameInfo := lipgloss.JoinVertical(lipgloss.Center,
			infoStyle.Render(boxscore.Clock),
			infoStyle.Render(boxscore.CurrentPeriod),
		)

		scoreAwayTeam := lipgloss.JoinVertical(lipgloss.Center, teamNameStyle.Render(away.NameShort), getBigScoreText(away.Score))
		dash := getBigScoreText(420)
		scoreHomeTeam := lipgloss.JoinVertical(lipgloss.Center, teamNameStyle.Render(home.NameShort), getBigScoreText(home.Score))

		ui := lipgloss.JoinVertical(lipgloss.Center, lipgloss.JoinHorizontal(lipgloss.Center, scoreAwayTeam, dash, scoreHomeTeam), gameInfo)

		gameBoard := lipgloss.Place(width, 17,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars(" "),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		board := lipgloss.JoinVertical(lipgloss.Center, gameBoard, infoStyle.Render(fmt.Sprintf("%s - %s", boxscore.Location.City, boxscore.Location.Venue)))

		doc.WriteString(board + "\n")
	}

	return doc.String()
}

func getBigScoreText(number int) string {
	if number == 420 {
		return scoreTextStyle.Render(scoreTextFont[420])
	}

	scoreSlice := splitInt(number)
	result := ""

	for _, v := range scoreSlice {
		result = lipgloss.JoinHorizontal(lipgloss.Top, result, scoreTextStyle.Render(scoreTextFont[v]))
	}

	return result
}

func splitInt(n int) []int {
	if n == 0 {
		return []int{0}
	}

	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	result := []int{}
	for i := range slc {
		result = append(result, slc[len(slc)-1-i])
	}

	return result
}
