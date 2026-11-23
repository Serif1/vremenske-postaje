
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Serif1/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "postaje",
		Usage: "Program za spremljanje temperaturnih meritev",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "Najnižja dovoljena temperatura",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "Najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Najmanjše število meritev za statistiko",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			if minTemp > maxTemp {
				return fmt.Errorf("minTemp ne sme biti večji od maxTemp")
			}
			if stMeritev < 1 {
				return fmt.Errorf("stMeritev mora biti vsaj 1")
			}

			postaje := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

			// Začetne meritve iz prve domače naloge.
			zacetneMeritve := map[string][]float64{
				"KP01": {24.5, 27, 26.1, 28.4},
				"KR01": {2, -1.5, 3.5},
				"LJ01": {18.2, 21.5, 19.8, 16.5},
				"MB01": {12.4},
			}

			for _, oznaka := range []string{"KP01", "KR01", "LJ01", "MB01"} {
				for _, temperatura := range zacetneMeritve[oznaka] {
					vreme.DodajMeritev(postaje, oznaka, temperatura, minTemp, maxTemp)
				}
			}

			vreme.IzpisMeritev(postaje)
			vreme.IzpisPorocila(postaje, stMeritev)

			fmt.Println("\nPREIZKUS DODAJANJA:")
			vreme.DodajMeritev(postaje, "MB01", 14.2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", 16.1, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 65, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "XX01", 20, minTemp, maxTemp)

			vreme.IzpisMeritev(postaje)
			vreme.IzpisPorocila(postaje, stMeritev)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
