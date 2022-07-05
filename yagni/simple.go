package yagni

import (
	"crypto/tls"
	"fmt"
	"strings"

	gomail "gopkg.in/mail.v2"
)

var receiver = "bestilling@bedrift.no"

type KundeDetaljer struct {
	Navn       string
	Epost      string
	Mobilnr    string
	Addresse   string
	Addresse2  string
	Postnummer string
	Poststed   string
}

func (kd *KundeDetaljer) LeveringsAdresse() string {
	sb := strings.Builder{}
	sb.WriteString(kd.Navn + "\n")
	sb.WriteString(kd.Addresse + "\n")
	if kd.Addresse2 != "" {
		sb.WriteString("  " + kd.Addresse2 + "\n")
	}
	sb.WriteString(kd.Postnummer + " " + kd.Poststed + "\n")
	return sb.String()
}

type OrdreLinje struct {
	Antall    int
	Produkt   string
	Størrelse string
	Pris      float32
}

func (ol *OrdreLinje) TotalPris() float32 {
	return float32(ol.Antall) * ol.Pris
}

func (ol *OrdreLinje) String() string {
	return fmt.Sprintf("%s %s\t%d\t%.2f\t%.2f\n", ol.Produkt, ol.Størrelse, ol.Antall, ol.Pris, ol.TotalPris())
}

func PlaceOrder(sendTil *KundeDetaljer, ordreLinjer []*OrdreLinje) error {
	mail := gomail.NewMessage()

	mail.SetHeaders(map[string][]string{
		"From":    {sendTil.Epost},
		"To":      {receiver},
		"Subject": {fmt.Sprintf("Bestilling: %s", sendTil.Navn)},
	})

	ordre := ""
	tot := float32(0)
	for _, ol := range ordreLinjer {
		ordre += ol.String()
		tot += ol.TotalPris()
	}
	ordre += fmt.Sprintf("Totalt: %.2f", tot)

	body := fmt.Sprintf("Ny Bestilling\n\nKunde:\n%s\n%s\n%s\n\nLeveringsadresse:\n%s\n\nFakturaadresse:\n%s\n\nOrdre:\n%s", sendTil.Navn, sendTil.Epost, sendTil.Mobilnr, sendTil.LeveringsAdresse(), sendTil.LeveringsAdresse(), ordre)

	mail.SetBody("text/plain", body)
	fmt.Println(body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "from@gmail.com", "<email_password>")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
