package menu

import (
	promptdata "weather/promptData"
)

func ShowMenu() string {
	prompt := promptdata.PromptData(
		"\n   -- Прогноз погоды --\n",
		"1. Погода в своем городе\n",
		"2. Погода в другом городе\n",
		"3. Выход\n",
	)
	return prompt
}
