package telegram

import (
	"elcorteingles/domain"
	"fmt"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct{
	bot *tgbotapi.BotAPI
}

// func notifyProduct(pr domain.Product) error{
// 	tgbotapi.NewMessageToChannel("@DGx24","testPostInChanel")
// 	return nil
// }

func NewTelegram() Telegram{
	
	bot, err := tgbotapi.NewBotAPI("5301405725:AAEfeUJtuaGgO9ESxuB-yy25m9yPMf2zrik")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return Telegram{bot}
}

func (t *Telegram) NotifyProduct(pr domain.Product) error{
	fmt.Println("Notificar producto telegram init")
	text := pr.Title+"\n"+"Precio original: "+fmt.Sprint(pr.OriginalPrice)+"\n"+"Precio final: "+fmt.Sprint(pr.FinalPrice)+"\nDescuento: -"+fmt.Sprint(pr.Discount)+"%"+"\n"+pr.Link
	msg := tgbotapi.NewMessage(-694410984, text)
	rand.Seed(time.Now().UnixNano())
    n := rand.Intn(20)*10 // n will be between 0 and 10
    time.Sleep(time.Duration(n)*time.Second)
	t.bot.Send(msg)
	return nil
}