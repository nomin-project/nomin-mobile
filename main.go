package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/nomin-project/nomin-mobile/sender"
)

func main() {
	application := app.New()
	window := application.NewWindow("Nomin")

	from := widget.NewSelectEntry(curators)
	from.SetPlaceHolder("FROM:")
	to := widget.NewSelectEntry(galleries)
	to.SetPlaceHolder("TO:")
	subject := widget.NewEntry()
	subject.SetPlaceHolder("SUBJECT:")
	text := widget.NewMultiLineEntry()
	text.SetPlaceHolder("TEXT:")

	server := widget.NewEntry()
	server.SetPlaceHolder("SMTP server address")
	port := widget.NewEntry()
	port.SetPlaceHolder("Port number")

	provider := widget.NewSelect(providers, func(option string) {
		setProvider(option, server, port)
	})
	provider.PlaceHolder = "Choose your internet provider"

	sendButton := widget.NewButton("Send Mail", func() {
		verifyAndSend(from, to, subject, text, server, port, window)
	})
	sendButton.Importance = widget.HighImportance

	aboutURL, _ := url.Parse("https://gajdosik.org/nomin")
	aboutNomin := widget.NewHyperlink("About Nomin app.", aboutURL)

	mailAbout := widget.NewRichTextWithText("Get recommended by a famous curator. Write an email recommending your artistic persona, select the sending curator and receiving gallery or other institution and just get recommended without any need for boring and slow networking!")
	mailAbout.Wrapping = fyne.TextWrapWord

	smtpAbout := widget.NewRichTextWithText("To send the email Nomin needs the ADDRESS and the PORT of open SMTP mail server. You can insert one manually or choose SMTP server of your internet provider.")
	smtpAbout.Wrapping = fyne.TextWrapWord

	sendAbout := widget.NewRichTextWithText("Check the texts, addresses and configuration. If everything is allright, just send it out!")
	sendAbout.Wrapping = fyne.TextWrapWord

	smtpLayout := container.NewVBox(
		container.NewPadded(
			smtpAbout,
		),
		provider,
		server,
		port,
	)

	mailLayout := container.NewVBox(
		container.NewPadded(
			mailAbout,
		),
		from,
		to,
		subject,
		text,
	)

	sendLayout :=
		container.New(
			layout.NewVBoxLayout(),
			container.NewPadded(
				sendAbout,
			),
			sendButton,
			layout.NewSpacer(),
			container.NewMax(
				container.NewCenter(
					aboutNomin,
				),
			),
		)

	tabs := container.NewAppTabs(
		container.NewTabItem("1 Write", mailLayout),
		container.NewTabItem("2 Configure", smtpLayout),
		container.NewTabItem("3 Send", sendLayout),
	)

	window.SetContent(tabs)
	window.Show()
	application.Run()
}

func verifyAndSend(from, to *widget.SelectEntry, subject, text, address, port *widget.Entry, window fyne.Window) {
	err := sender.SendMail(
		from.Text,
		to.Text,
		subject.Text,
		text.Text,
		address.Text,
		port.Text,
	)

	if err != nil {
		w := dialog.NewInformation("Error!", "Message was not sent.\n\n1. Check your network connection.\n2. Check your SMTP configuration.\n\n"+err.Error(), window)
		w.Resize(window.Content().Size())
		w.Show()
	} else {
		w := dialog.NewInformation("Success!", "Email successfully sent!", window)
		w.Resize(window.Content().Size())
		w.Show()
	}
}

var providers = []string{
	"cz: UPC",
	"cz: Vodafone",
	"cz: T-Mobile",
	"custom",
}

func setProvider(provider string, server, port *widget.Entry) {
	var s, p string
	switch provider {
	case "cz: UPC":
		s = "mail.upcmail.cz"
		p = "25"
	case "cz: Vodafone":
		s = "smtp.vodafonemail.cz"
		p = "25"
	case "cz: T-Mobile":
		s = "smtp.t-email.cz"
		p = "25"
	default:
		s = ""
		p = "25"
	}

	server.SetText(s)
	port.SetText(p)

}

var curators = []string{
	"Naomi Beckwith <nbeckwith@guggenheim.org>",
	"Nicolas Bourriaud <NicolasBourriaud@moco.art>",
	"Cornelia Butler <cbutler@hammer.ucla.edu>",
	"RoseLee Goldberg <roselee.goldberg@nyu.edu>",
	"Thelma Golden <TGolden@studiomuseum.org>",
	"Hanru Hou <hanru.hou@fondazionemaxxi.it>",
	"Candice Hopkins <chopkins@torontobiennial.org>",
	"Bonaventure Ndikung <ndikung@gmx.de>",
	"Fatos Ustek <mail@fatosustek.com>",
	"KM Temporaer <info@kmtemporaer.de>",
	"Milda Batakyte <milda@rupert.lt>",
	"Hanne Mugaas <hanne@kunsthallstavanger.no>",
	"Samuel Leuenberger <samuel@salts.ch>",
	"Ovul Durmusoglu <durmusoglu@adbk-nuernberg.de>",
	"Krist Gruijthuijsen <hinnik50@hotmail.com>",
	"Melissa Chiu <melissa.chiu1@gmail.com>",
	"Julija Reklaite <julija@rupert.lt>",
}

var galleries = []string{
	"exhibitions@whitecube.com",
	"info@agora-gallery.com",
	"berlin@blainsouthern.com",
	"info@blainsouthern.com",
	"info@bombonprojects.com",
	"info@ermes-ermes.com",
	"info@marfaprojects.com",
	"info@bws.mx",
	"office@berlinbiennale.de",
	"info@cjch.cz",
}
