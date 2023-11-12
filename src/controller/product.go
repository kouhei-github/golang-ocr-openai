package controller

import (
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"net-http/myapp/utils"
	"net-http/myapp/utils/google"
	"net-http/myapp/utils/google/documents"
	"net-http/myapp/utils/google/drives"
	"net/http"
	"os"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"message": "Method Not Allowed"})
		return
	}
	// フォームデータの取得（最大100MB）
	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// フォームから "uploadfile"と名付けられたキャプチャを取得
	file, _, err := r.FormFile("target_file")

	defer file.Close()

	// 構造体GoogleClient生成
	myClient, err := google.NewClient(google.DriveFileScope, google.DocumentReadonlyScope)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// 構造体DriveFacadeの生成
	driveFacade, err := drives.NewDriveFacade(myClient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// GoogleDriveアップロード
	fileId, err := driveFacade.Upload("1YjFwicSDEqJZBug-MNHonEZFlUfG5XYV", file)
	json.NewEncoder(w).Encode(map[string]string{"message": fileId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// 構造体DriveFacadeの生成
	docsFacade, err := documents.NewDocumentFacade(myClient, fileId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	// GoogleDocumentで読み込む
	text, err := docsFacade.Read()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	// Documentを削除
	if err = driveFacade.Delete(fileId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message2": err.Error()})
		return
	}

	// プロンプト
	prompt := r.FormValue("prompt")
	var messages []openai.ChatCompletionMessage
	chat := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: utils.Generate(prompt, text),
	}
	messages = append(messages, chat)

	answer, err := utils.SendToGPT(os.Getenv("OPENAI_API_KEY"), messages)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": answer})
}
