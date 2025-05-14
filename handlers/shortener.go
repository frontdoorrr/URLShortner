// handlers/shortener.go
package handlers

import (
	"net/http"
	"net/url"
	"url-shortner/storage"
	"url-shortner/utils"

	"github.com/gin-gonic/gin"
)

type URLRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return false
	}
	return true
}

// ShortenURL은 URL 단축 요청을 처리합니다.
func ShortenURL(c *gin.Context) {
	var request URLRequest

	// JSON 바인딩: 요청의 Body를 구조체로 변환합니다.
	if err := c.ShouldBindJSON(&request); err != nil {
		// 오류가 발생하면 400 Bad Request를 반환합니다.
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if !isValidURL(request.OriginalURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format. URL must start with http:// or https://"})
		return
	}

	// 8자리 랜덤 문자열 생성
	shortURL := utils.RandomString(8)

	// In-Memory Storage에 저장
	storage.SaveURL(shortURL, request.OriginalURL)

	// 클라이언트에게 응답
	c.JSON(http.StatusOK, URLResponse{
		ShortURL: "http://localhost:8080/" + shortURL,
	})
}
