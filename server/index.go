package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func groq() {
	godotenv.Load()
	var GROQ_API_KEY = os.Getenv("GROQ_API_KEY")
	url := "https://api.groq.com/openai/v1/chat/completions"
	contentType := "application/json"
	data := []byte(`{"messages": [{"role": "user", "content": "Explain the importance of fast language models"}], "model": "llama3-8b-8192", "stream": true}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+ GROQ_API_KEY)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	var respStr strings.Builder
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respStr.Write(content)
	fmt.Println(respStr.String())
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Wagwan",
		})
	})
	r.POST("/", func(ctx *gin.Context) {
		var res gin.H
		json.NewDecoder(ctx.Request.Body).Decode(&res)
		ctx.JSON(http.StatusOK, res)
	})
	r.Run(":3000")
}