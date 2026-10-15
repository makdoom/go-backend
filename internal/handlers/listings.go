package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/makdoom/go-backend/internal/middlewares"
)

type listing struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListingHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandler {
	return &ListingHandler{db: db, logger: logger}
}

func (lh *ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rows, err := lh.db.QueryContext(ctx, `
			SELECT id, title, description, price, city, created_at AS createdAt, updated_at AS updatedAt 
			FROM listings 
			ORDER BY created_at DESC 
			LIMIT 5
		`)
	if err != nil {
		lh.logger.Error("Handler.ListDB query failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	listings := []listing{}
	for rows.Next() {
		var l listing

		if err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt, &l.UpdatedAt); err != nil {
			lh.logger.Error("Handler.List: DB scan failed", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		listings = append(listings, l)
	}

	if err := rows.Err(); err != nil {
		lh.logger.Error("Handler.List: DB iteration failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := rows.Err(); err != nil {
		lh.logger.Error("Handler.List: DB iteration failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listings)

}

func (lh *ListingHandler) DeleteListing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, ok := middlewares.RequestIdFromContext(ctx)
	if !ok {
		requestId = "unknown"
	}

	id := r.PathValue("id")

	_, err := lh.db.ExecContext(ctx, `DELETE FROM listing WHERE id = $1`, id)
	if err != nil {
		lh.logger.Error("Handler.Delete: DB Delete failed", "listing_id", id, "requestId", requestId, "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
