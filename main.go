package main

import (
	"encoding/json"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)
/*
type cbrf struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       struct {
		Aud struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"AUD"`
		Azn struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"AZN"`
		Gbp struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"GBP"`
		Amd struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"AMD"`
		Byn struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"BYN"`
		Bgn struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"BGN"`
		Brl struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"BRL"`
		Huf struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"HUF"`
		Hkd struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"HKD"`
		Dkk struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"DKK"`
		Usd struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"USD"`
		Eur struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"EUR"`
		Inr struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"INR"`
		Kzt struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"KZT"`
		Cad struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"CAD"`
		Kgs struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"KGS"`
		Cny struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"CNY"`
		Mdl struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"MDL"`
		Nok struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"NOK"`
		Pln struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"PLN"`
		Ron struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"RON"`
		Xdr struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"XDR"`
		Sgd struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"SGD"`
		Tjs struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"TJS"`
		Try struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"TRY"`
		Tmt struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"TMT"`
		Uzs struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"UZS"`
		Uah struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"UAH"`
		Czk struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"CZK"`
		Sek struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"SEK"`
		Chf struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"CHF"`
		Zar struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"ZAR"`
		Krw struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"KRW"`
		Jpy struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"JPY"`
	} `json:"Valute"`
}


*/



type binanceResp struct {
	Price float64 `json:"price,string"`
	Code int64 `json:"code"`
}



type wallet map[string]float64

var db = map[int64]wallet{}



func main() {


	bot, err := tgbotapi.NewBotAPI("2114656862:AAEdmSJIVnJL9ANNpDqSEPH18nTahE4Feg4")
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


//		if len(msgArr[3]) > 0 {
//			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, Слишком много аргументов"))
//			continue
//		}



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
				db[update.Message.Chat.ID][msgArr[1]] += summ
				msg := fmt.Sprintf("Баланс: %s %f", msgArr[1], db[update.Message.Chat.ID][msgArr[1]])
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))

			case "SUB":
				summ, err := strconv.ParseFloat(msgArr[2], 64)
				//проверка на возможность и не отрицательность операции
				if err != nil || summ > db[update.Message.Chat.ID][msgArr[1]]{
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
				var usdSumm,rubSumm,usd,rub float64

				cointPriceRub, err := getCurRub()
				if err != nil {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				}

				for key, value :=range db[update.Message.Chat.ID]{
					cointPrice, err:= getPrice(key)
					if err != nil {
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
					}
				usd = value * cointPrice
				rub = value * cointPrice * cointPriceRub
				usdSumm += usd
				rubSumm += rub
				msg +=	fmt.Sprintf("%s: %f [%.2f - USD] [%.2f - RUB]\n", key, value, usd , value * cointPrice * cointPriceRub)
				}
				msg += fmt.Sprintf("Сумма USD: %.2f\n", usdSumm)
				msg += fmt.Sprintf("Сумма RUB: %.2f\n", rubSumm)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))

			default:
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка, неизвестная команда"))
			}

	}

}

func getPrice(coin string) (price float64, err error){
resp, err :=	http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", coin))
if err != nil {return}
defer resp.Body.Close()

 var jsonResp binanceResp
err = json.NewDecoder(resp.Body).Decode(&jsonResp)
if err != nil { return }

if jsonResp.Code !=0 {
	err = errors.New ("Некорректная валюта")
	return
}

price = jsonResp.Price

	return
}

func getCurRub()(PriceRub float64, err error)  {
resp, err :=	http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
if err != nil {return}
defer resp.Body.Close()


body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}


	avg := gojsonq.New().JSONString(string(body)).Find("Valute.USD.Value")
	PriceRub = avg.(float64)

return
}