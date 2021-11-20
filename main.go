package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/nomin-project/nomin-mobile/sender"
)

func main() {
	application := app.New()
	window := application.NewWindow("Nomin")

	curators := []string{
		"Michal Novotný <michal.novotny@futura.cz>",
		"Michal Novotný 2 <michal.novotny2@futura.cz>",
	}

	galleries := []string{
		"Michal Novotný <michal.novotny@futura.cz>",
		"Michal Novotný 2 <michal.novotny2@futura.cz>",
	}

	sender := widget.NewSelectEntry(curators)
	recipient := widget.NewSelectEntry(galleries)

	hello := widget.NewLabel("Hello Fyne!")

	window.SetContent(container.NewVBox(
		hello,
		sender,
		recipient,
		widget.NewButton("Send Mail", func() {

			sendMail()
		}),
	))

	window.ShowAndRun()
}

func sendMail() {
	fmt.Println("blabla")
	//(from string, to string, subject string, message string, serverAddress string, serverPort string)
	sender.SendMail("janzalesak@ffa.vutbr.cz", "andreas.gajdosik@gmail.com", "Hello!", "Hello Andreas, it is nice to chat again!", "smtp.vodafonemail.cz", "25")
}
