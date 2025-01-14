package types

import "github.com/sashabaranov/go-openai"

type AIService interface {
	InjectRegionService(service RegionsService)
	InjectCarpoolsService(service CarpoolsService)
	InjectMatchingService(service MatchingService)
	DefineFunctions() []openai.FunctionDefinition
	RecommendCourses(req *RecommendCourseReq) ([]*CourseRecommendationAIRes, error)
	GetTourRecommendations(region string, interests []string) (string, error)
	RecommendMatchingPost(page int, pageSize int, location string, interests []string) ([]*MatchingDetailForAI, error)
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatFormat struct {
	Messages []*ChatMessage `json:"messages"`
}

const TrainingDataPath = "/app/internal/data/output.jsonl"

type AnswerResponse struct {
	Answer string `json:"content"`
}
