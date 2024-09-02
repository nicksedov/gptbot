package localai

import (
	"context"
	"fmt"
	"gptbot/pkg/settings"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	openai "github.com/gopenai/openai-client"
)

// Warning: this is rather long lasting test, use DEBUG mode when launching from VSCode IDE
func TestSendRequest(t *testing.T) {
	
	client, err := openai.NewClient("http://localhost:5555/")
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
	prompt := "Расскажи команде, что завтра внедрение в пром релиза 20240831."

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
			{
				Role:    openai.ChatCompletionRequestMessageRoleAssistant,
				Content: "...",
			},
		},
	}
	startTime := time.Now()
	response, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	duration := time.Since(startTime)

	// Print the completed text
	fmt.Printf("Время работы: %s\nРезультат: %s\n", duration, response.Choices[0].Message.Value.Content)
}

func TestSendRequest2(t *testing.T) {
	settings := settings.GetSettings()
	testChatId := settings.Telegram.ServiceChat
	settings.GigaChat.Completions.Context = "Ты - специалист по истории Франции."
	assert.NotEqual(t, 0, testChatId)
	req1 := "Расскажи, кто возглавлял Францию в 1812 году"
	req2 := "В каких битвах он принимал участие в этот год?"
	res1 := SendRequest(testChatId, req1)
	res2 := SendRequest(testChatId, req2)
	fmt.Printf("Вопрос:\n%s\nОтвет:\n%s\nВопрос:\n%s\nОтвет:\n%s\n", 
	req1, res1.Choices[0].Message.Value.Content, 
	req2, res2.Choices[0].Message.Value.Content)
}
