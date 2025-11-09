package wifi

import (
	"database/sql"
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

	return wifiReviews, nil
}

func QueryWifiPassword(wifiID int64, db *sql.DB) ([]WifiPasswordDTO, error) {
	var wifiReviews []WifiPasswordDTO

	rows, err := db.Query(`
        SELECT wp.id, wp.created_at, wp.wifi_id, wp.password
        FROM wifi_password wp
        WHERE wifi_id = $1
    `, wifiID)

	if err != nil {
		return wifiReviews, err
	}

	defer rows.Close()

	for rows.Next() {
		var review WifiPasswordDTO
		if err := rows.Scan(
			&review.WifiPasswordID,
			&review.CreatedAt,
			&review.WifiID,
			&review.Password,
		); err != nil {
			return wifiReviews, err
		}

		wifiReviews = append(wifiReviews, review)
	}

	return wifiReviews, nil
}
