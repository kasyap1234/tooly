package notes

import (
	"encoding/json"
	"net/http"
	

	"github.com/go-chi/chi"
"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	ID      uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateNotes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var notes Notes

		if err := json.NewDecoder(r.Body).Decode(&notes); err!= nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		notes.ID= uuid.New();
		
		if err := db.Create(&notes).Error; err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(notes)
	}
}
func GetNotes(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var notes []Notes
		if err := db.Find(&notes).Error; err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(notes)
	}
}
func GetNote(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "id")
		
		var note Notes
		if err := db.First(&note,"id=?", noteID).Error; err!= nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Record not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(note)
	}
}
func UpdateNote(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        noteID := chi.URLParam(r, "id")
        var note Notes
        if err := db.First(&note, "id = ?", noteID).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                http.Error(w, "Record not found", http.StatusNotFound)
            } else {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }
        if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }    
        if err := db.Save(&note).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(note)
    }
}
func DeleteNote(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID:= chi.URLParam(r, "id")
		
		var note Notes
		if err := db.First(&note, "id=?",noteID).Error; err!= nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Record not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if err := db.Delete(&note).Error; err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}



