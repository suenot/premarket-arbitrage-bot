package display

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/suenot/w_premarket_arbitrage/scanner"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7B2FBE")).
			Padding(0, 1)

	greenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF87"))
	yellowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFDA6B"))
	redStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6B6B"))
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	titleStyle  = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF79C6")).
			Padding(1, 0)
	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8BE9FD")).
			Italic(true)
	balanceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#50FA7B")).
			Bold(true)
)

func colorAPR(apr float64) string {
	s := fmt.Sprintf("%.1f%%", apr)
	switch {
	case apr >= 20:
		return greenStyle.Render(s)
	case apr >= 10:
		return yellowStyle.Render(s)
	default:
		return redStyle.Render(s)
	}
}

func colorProfit(pct float64) string {
	s := fmt.Sprintf("%.1f%%", pct)
	switch {
	case pct >= 5:
		return greenStyle.Render(s)
	case pct >= 3:
		return yellowStyle.Render(s)
	default:
		return redStyle.Render(s)
	}
}

func colorDepth(usd float64) string {
	s := fmt.Sprintf("$%.0f", usd)
	switch {
	case usd >= 100:
		return greenStyle.Render(s)
	case usd >= 20:
		return yellowStyle.Render(s)
	default:
		return dimStyle.Render(s)
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-1] + "…"
}

// PortfolioInfo holds optional portfolio data to display.
type PortfolioInfo struct {
	TotalBalance    float64
	PolygonBalance  float64
	ActivePositions int
	TotalValue      float64
}

// Render prints the arbitrage opportunities as a styled CLI table.
func Render(pairs []scanner.ArbitragePair, total int, dryRun bool, portfolio *PortfolioInfo) {
	mode := greenStyle.Render("🟢 DRY RUN")
	if !dryRun {
		mode = redStyle.Render("🔴 LIVE TRADING")
	}

	fmt.Println(titleStyle.Render("⚡ Premarket Arbitrage Scanner"))

	// Status line
	statusParts := []string{mode, fmt.Sprintf("Found: %d matched / %d total", len(pairs), total)}

	// Portfolio info
	if portfolio != nil {
		balanceStr := balanceStyle.Render(fmt.Sprintf("💰 $%.2f USDC", portfolio.TotalBalance))
		posStr := fmt.Sprintf("📊 %d positions ($%.2f)", portfolio.ActivePositions, portfolio.TotalValue)
		statusParts = append(statusParts, balanceStr, posStr)
	}

	fmt.Println(statusStyle.Render("   " + strings.Join(statusParts, "  |  ")))
	fmt.Println()

	if len(pairs) == 0 {
		fmt.Println(dimStyle.Render("   No opportunities matching your filters."))
		fmt.Println()
		return
	}

	headers := []string{"#", "Event", "Buy YES @", "Buy NO @", "Cost", "Profit", "APR", "Depth", "Days"}

	var rows [][]string
	for i, p := range pairs {
		if i >= 25 {
			break
		}
		rows = append(rows, []string{
			fmt.Sprintf("%d", i+1),
			truncate(p.Platform1.Question, 42),
			fmt.Sprintf("%.3f (%s)", p.Arbitrage.BuyYesPrice, p.Arbitrage.BuyYesSource),
			fmt.Sprintf("%.3f (%s)", p.Arbitrage.BuyNoPrice, p.Arbitrage.BuyNoSource),
			fmt.Sprintf("%.3f", p.Arbitrage.TotalCost),
			colorProfit(p.Arbitrage.ProfitPct),
			colorAPR(p.Arbitrage.APR),
			colorDepth(p.Depth.MaxUSD),
			fmt.Sprintf("%d", p.Arbitrage.DaysUntil),
		})
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))).
		Headers(headers...).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return lipgloss.NewStyle().Padding(0, 1)
		})

	fmt.Println(t)
	fmt.Println()

	// Platform legend
	platforms := make(map[string]bool)
	for _, p := range pairs {
		platforms[p.Platform1.Source] = true
		platforms[p.Platform2.Source] = true
	}
	var names []string
	for name := range platforms {
		names = append(names, name)
	}
	fmt.Println(dimStyle.Render(fmt.Sprintf("   Platforms: %s", strings.Join(names, ", "))))
	fmt.Println()
}
