package utils

import (
	"kos-management/repositories"
	"log"
	"time"
)

// BillingScheduler handles automated billing tasks
type BillingScheduler struct {
	TagihanRepo *repositories.TagihanRepository
}

// NewBillingScheduler creates a new billing scheduler
func NewBillingScheduler(tagihanRepo *repositories.TagihanRepository) *BillingScheduler {
	return &BillingScheduler{
		TagihanRepo: tagihanRepo,
	}
}

// StartMonthlyBillingCheck starts the monthly billing check process
func (s *BillingScheduler) StartMonthlyBillingCheck() {
	go func() {
		for {
			// Update overdue bills
			if err := s.TagihanRepo.UpdateOverdueBills(); err != nil {
				log.Printf("Error updating overdue bills: %v", err)
			} else {
				log.Println("Successfully updated overdue bills")
			}

			// Sleep until next check (daily at midnight)
			now := time.Now()
			nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			time.Sleep(nextMidnight.Sub(now))
		}
	}()
}
