package main

import (
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type binanceResp struct {
	Price float64 `json:"price,string"`
	Code  int64   `json:"code"`
}

type wallet map[string]float64

var db = map[int64]wallet{}

func main() {

	bot, err := tgbotapi.NewBotAPI("2114656862:AAEdmSJIVnJL9ANNpDqSEPH18nTahE4Feg41")
	if err != nil {
		log.Panic(err)
	}

	//	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msgArr := strings.Split(update.Message.Text, " ")

		if len(msgArr) > 3 {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, Слишком много аргументов"))
			continue
		}

		switch msgArr[0] {
		case "ADD":
			summ, err := strconv.ParseFloat(msgArr[2], 64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, невозможно добавить валюту"))
				continue
			}

			if _, ok := db[update.Message.Chat.ID]; !ok {
				db[update.Message.Chat.ID] = wallet{}
			}

			_, err = getPrice(msgArr[1])
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Нет Такой валюты"))
				continue
			}

			db[update.Message.Chat.ID][msgArr[1]] += summ
			msg := fmt.Sprintf("Баланс: %s %f", msgArr[1], db[update.Message.Chat.ID][msgArr[1]])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))

		case "SUB":
			summ, err := strconv.ParseFloat(msgArr[2], 64)
			//проверка на возможность и не отрицательность операции
			if err != nil || summ > db[update.Message.Chat.ID][msgArr[1]] {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, невозможно выполнить операцию"))
				continue
			}
			if _, ok := db[update.Message.Chat.ID]; !ok {
				db[update.Message.Chat.ID] = wallet{}
			}
			db[update.Message.Chat.ID][msgArr[1]] -= summ
			msg := fmt.Sprintf("Баланс: %s %f", msgArr[1], db[update.Message.Chat.ID][msgArr[1]])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
		case "DEL":
			delete(db[update.Message.Chat.ID], msgArr[1])
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Валюта удалена"))
		case "SHOW":
			msg := "Баланс:\n"
			var usdSumm, rubSumm, usd, rub float64

			cointPriceRub, err := getCurRub()
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			}

			for key, value := range db[update.Message.Chat.ID] {
				cointPrice, err := getPrice(key)
				if err != nil {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				}
				usd = value * cointPrice
				rub = value * cointPrice * cointPriceRub
				usdSumm += usd
				rubSumm += rub
				msg += fmt.Sprintf("%s: %f [%.2f - USD] [%.2f - RUB]\n", key, value, usd, value*cointPrice*cointPriceRub)
			}
			msg += fmt.Sprintf("Сумма USD: %.2f\n", usdSumm)
			msg += fmt.Sprintf("Сумма RUB: %.2f\n", rubSumm)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))

		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, неизвестная команда"))
		}

	}

}

func getPrice(coin string) (price float64, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", coin))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var jsonResp binanceResp
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return
	}

	if jsonResp.Code != 0 {
		err = errors.New("Некорректная валюта")
		return
	}

	price = jsonResp.Price

	return
}

func getCurRub() (PriceRub float64, err error) {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	avg := gojsonq.New().JSONString(string(body)).Find("Valute.USD.Value")
	PriceRub = avg.(float64)

	return
}
