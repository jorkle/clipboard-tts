package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"golang.design/x/clipboard"
)

func readClipboard() string {
	clipboard.Init()
	return string(clipboard.Read(clipboard.FmtText))
}

func generateTextExplanation(text string, apiKey string) string {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You are a helpful research assistant that accepts arbitrary text from the user and explains the concepts and information within that text that may not be obvious or common knowledge."),
			openai.UserMessage(text),
		}),
		Model:      openai.F(openai.ChatModelGPT4o),
		Modalities: openai.F([]openai.ChatCompletionModality{openai.ChatCompletionModalityText}),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return (chatCompletion.Choices[0].Message.Content)
}

func generateAudio(text string, apiKey string) {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	model := openai.SpeechModelTTS1
	voice := openai.AudioSpeechNewParamsVoiceAlloy
	format := openai.AudioSpeechNewParamsResponseFormatMP3
	req := openai.AudioSpeechNewParams{
		Input:          openai.F(text),
		Model:          openai.F(openai.SpeechModel(model)),
		Voice:          openai.F(openai.AudioSpeechNewParamsVoice(voice)),
		ResponseFormat: openai.F(openai.AudioSpeechNewParamsResponseFormat(format)),
		Speed:          openai.F(1.0),
	}
	resp, err := client.Audio.Speech.New(context.Background(), req)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	tempFile := os.TempDir() + "/clipboard-tts.mp3"
	if _, err := os.Stat(tempFile); err == nil {
		os.Remove(tempFile)
	}
	temporaryFile, err := os.Create(tempFile)
	if err != nil {
		log.Fatal(err)
	}
	defer temporaryFile.Close()
	_, err = io.Copy(temporaryFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func playAudio() {
	audioFile := os.TempDir() + "/clipboard-tts.mp3"
	cmd := exec.Command("ffplay", audioFile, "-autoexit", "-nodisp")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(audioFile)
}

func main() {
	apiKey := flag.String("apikey", "", "OpenAI API Key")
	flag.Parse()
	if *apiKey == "" {
		log.Fatal("API key is required")
	}

	text := readClipboard()

	generateAudio(text, *apiKey)
	playAudio()
}
