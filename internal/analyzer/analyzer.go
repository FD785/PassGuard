package analyzer

import "passguard/internal/checks"

type Result struct {
	Score           int      `json:"score"`
	RiskLevel       string   `json:"risk_level"`
	Issues          []string `json:"issues"`
	Recommendations []string `json:"recommendations"`
}

func Analyze(password string) Result {
	issues := []string{}
	recs := []string{}
	score := 100

	if len(password) < 12 {
		issues = append(issues, "Password too short")
		recs = append(recs, "Use at least 12 characters")
		score -= 25
	}

	if !checks.HasUpper(password) {
		issues = append(issues, "Missing uppercase letters")
		recs = append(recs, "Add uppercase characters")
		score -= 10
	}

	if !checks.HasLower(password) {
		issues = append(issues, "Missing lowercase letters")
		recs = append(recs, "Add lowercase characters")
		score -= 10
	}

	if !checks.HasNumber(password) {
		issues = append(issues, "Missing numbers")
		recs = append(recs, "Add numeric characters")
		score -= 10
	}

	if !checks.HasSymbol(password) {
		issues = append(issues, "Missing symbols")
		recs = append(recs, "Add special characters")
		score -= 10
	}

	if checks.IsCommon(password) {
		issues = append(issues, "Common password pattern")
		recs = append(recs, "Avoid dictionary words")
		score -= 30
	}

	risk := "low"
	if score < 70 {
		risk = "medium"
	}
	if score < 40 {
		risk = "high"
	}

	if score < 0 {
		score = 0
	}

	return Result{
		Score:           score,
		RiskLevel:       risk,
		Issues:          issues,
		Recommendations: recs,
	}
}
