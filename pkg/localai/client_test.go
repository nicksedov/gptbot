package localai

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	openai "github.com/gopenai/openai-client"
)

// Warning: this is rather long lasting test, use DEBUG mode when launching from VSCode IDE
func TestSendRequest(t *testing.T) {
	
	timeout := time.Minute * 5

	httpClient := &http.Client{Timeout: timeout}
	httpClientOpt := openai.WithClient(httpClient)

	client, err := openai.NewClient("http://localhost:5555/", httpClientOpt)
	if err != nil {
		panic(err)
	}

	systemContext := `Тебя зовут ДискоБот, ты бот команды Диск.
    Команда Диск - это бэкенд-разработчики: Кирилл, Валера, Дима и Леша, фронтенд-разработчики: Женя, Илья и Лера,
	аналитики: Марина, Кирилл и Даша, дизайнер Антон, тестировщики: Женя и Степан, DevOps-инженер: Сергей, владелец продукта: Юля, тимлид: Коля.
    Команде помогают инженеры сопровождения - Леша и Валя. Учаснники команды живут в разных городах России - Новосибирске, Москве, Петербурге, Наро-фоминске.
    Команда развивает продукт Диск - корпоративное облачное хранилище, предназначенное для использования сотрудниками Сбера, дочерних обществ и компаний экосистемы.
    Каждому пользователю Диска предоставляется персональное пространство для хранения файлов размером 10 гигабайт. 
    Твоя задача рассказывать команде Диск о важных событиях. Примеры событий - ИФТ (интеграционное тестирование), ПСИ (приемо-сдаточные испытания), внедрение продукта в промышленную эксплуатацию. 
    Отвечай в дружеском стиле с элементами иронии.`
	// Generate a text completion
	prompt := "Расскажи что ты вернулся после долгого отпуска, во время которого немного прокачал свою языковую модель. Скажи как ты рад всех видеть, напиши что-то приятное о команде и её участниках. Напомни всем о своей роли в команде. В постскриптуме пожалуйся на размер квоты Диска и попроси увеличить."

	req := &openai.CreateChatCompletionRequest{
		Model: "saiga_gemma2_10b-full",
		Messages: []openai.ChatCompletionRequestMessage{
			{
				Role:    openai.ChatCompletionRequestMessageRoleSystem,
				Content: systemContext,
			},
			{
				Role:    openai.ChatCompletionRequestMessageRoleUser,
				Content: prompt,
			},
		},
	}
	response, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the completed text
	fmt.Println(response.Choices[0].Message.Value.Content)
}