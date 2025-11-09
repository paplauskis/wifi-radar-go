package wifi

import (
	"database/sql"
	"log"
)

func QueryWifiReview(wifiID int64, db *sql.DB) ([]WifiReviewDTO, error) {
	var wifiReviews []WifiReviewDTO

	rows, err := db.Query(`
        SELECT wr.id, wr.created_at, wr.wifi_id, wr.text, wr.rating, u.username
        FROM wifi_review wr
        JOIN "user" u ON u.id = wr.user_id
        WHERE wifi_id = $1
        ORDER BY created_at DESC
    `, wifiID)

	if err != nil {
		return wifiReviews, err
	}

	defer rows.Close()

	log.Println("rows:", rows)
	for rows.Next() {
		var review WifiReviewDTO
		if err := rows.Scan(
			&review.WifiReviewID,
			&review.CreatedAt,
			&review.WifiID,
			&review.Text,
			&review.Rating,
			&review.Username,
		); err != nil {
			return wifiReviews, err
		}

		wifiReviews = append(wifiReviews, review)
	}
	log.Printf("wifiReviews: %v", wifiReviews)
	return wifiReviews, nil
}
