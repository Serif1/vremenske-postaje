# Vremenske postaje

To je 5. domača naloga pri predmetu Porazdeljeni sistemi.

Program je nadgradnja 1. domače naloge in je napisan v jeziku Go. Razdeljen je na paket vreme in glavni program.

Program omogoča dodajanje vremenskih postaj, dodajanje temperatur in izpis meritev. Za vsako postajo izračuna najnižjo, najvišjo in povprečno temperaturo, če ima dovolj meritev.

Program zaženemo z go run ./cmd/postaje.

Uporabljena je knjižnica urfave/cli/v3, s katero lahko spreminjamo najnižjo in najvišjo dovoljeno temperaturo ter najmanjše število meritev za izračun statistike.

Privzete vrednosti so -50 za najnižjo temperaturo, 50 za najvišjo temperaturo in 3 za število meritev.

V mapi vreme je koda za delo s postajami, v mapi cmd/postaje pa glavni program.