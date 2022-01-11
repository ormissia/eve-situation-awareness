package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"aeon/main/esa-zkill-history/model"
	"aeon/utils"
)

const dayLayout = "20060102"

var (
	dayChan = make(chan int)
	wg      sync.WaitGroup
)

func init() {
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(goroutineIndex int) {
			defer wg.Done()
			for {
				timeFlag := utils.TimeCost()
				day, ok := <-dayChan
				if !ok {
					break
				}
				currentDayUrl := "https://zkillboard.com/api/history/" + strconv.Itoa(day) + ".json"
				getClient := utils.GenerateGetRequest(currentDayUrl)
				respBody, err := getClient(nil)
				if err != nil {
					panic(err)
				}
				respJson := string(respBody)
				fmt.Println(respJson)
				fmt.Println(">>>>>\n>>>>>\n>>>>>")
				formatData := make(map[string]string)
				if err = json.Unmarshal([]byte(respJson), &formatData); err != nil {
					panic(err)
				}

				historyData := make([]model.KillmailHash, 0, len(formatData))
				for k, v := range formatData {
					killmailID, _ := strconv.Atoi(k)
					kmHash := model.KillmailHash{
						BelongDay:    day,
						KillmailId:   killmailID,
						KillmailHash: v,
						CreateTime:   time.Now(),
					}
					historyData = append(historyData, kmHash)
				}

				if err = new(model.KillmailHash).BatchSave(historyData); err != nil {
					panic(err)
				} else {
					fmt.Println(">>>>>>Insert successful!!!")
				}
				timeFlag(func(duration time.Duration) {
					fmt.Println(">>>>>>Worker goroutine index: " + strconv.Itoa(goroutineIndex) + " using time: " + duration.String())
				})
			}
		}(i)
	}
}

func HistoryDump() {
	defer utils.TimeCost()(func(duration time.Duration) {
		fmt.Println("using time: " + duration.String())
	})

	day := CurrentDay()
	today, _ := strconv.Atoi(time.Now().Format(dayLayout))
	for i := day(); i < today; {
		dayChan <- i
		fmt.Print(i)
		fmt.Print("\tto chan")

		// i++
		currentDay, _ := time.Parse(dayLayout, strconv.Itoa(i))
		nextDay := currentDay.Add(time.Hour * 24)
		nexDayStr := nextDay.Format(dayLayout)
		i, _ = strconv.Atoi(nexDayStr)
	}
	close(dayChan)
	wg.Wait()
}
