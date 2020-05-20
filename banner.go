package procyon

import "log"

type Banner interface {
	printBanner()
}

type ApplicationBanner struct {
}

const procyonAppText = " :: Procyon :: "

var bannerText = []string{"",
	"   ___",
	"  / _ \\  _ __   ___    ___  _   _   ___   _ __",
	" / /_)/ | '__| / _ \\  / __|| | | | / _ \\ | '_ \\",
	"/ ___/  | |   | (_) || (__ | |_| || (_) || | | |",
	"\\/     |_|     \\___/  \\___| \\__, | \\___/ |_| |_|",
	"                            |___/",
}

func (banner ApplicationBanner) PrintBanner() {
	for _, line := range bannerText {
		log.Print(line)
	}
}