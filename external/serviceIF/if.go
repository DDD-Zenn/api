package serviceIF

type (
	Gemini interface {
		GenerateResponse(prompt string) (string, error)
	}
)
