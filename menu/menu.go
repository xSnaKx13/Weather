package menu

import (
	"fmt"
	promptdata "weather/promptData"
)

func ShowMenu() string {
	prompt := promptdata.PromptData(
		"\n   -- Прогноз погоды --\n",
		"1. Погода в своем городе\n",
		"2. Погода в другом городе\n",
		"3. Выход\n",
	)

	var userCity string
	var isWorking bool = true
	for isWorking {
		switch prompt {
		case "1":
			userCity = ""
			isWorking = false
		case "2":
			userCity = promptdata.PromptData("Введите название города:")
			isWorking = false
		case "3":
			isWorking = false
		default:
			fmt.Println("\nНекорректный выбор.\nПожалуйста, выберите 1, 2 или 3.")
		}
	}
	return userCity
}
