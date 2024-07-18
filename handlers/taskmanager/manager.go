package taskmanager 
import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

)
type Task struct {
	gorm.Model 
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`	
	DueDate string `json:"dueDate"`
    
	
}
func GetTasks(db *gorm.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var tasks [] Task ; 
		if err :=db.Find(&tasks).Error; err!=nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return 
		}
		json.NewEncoder(w).Encode(tasks); 

	}
}
func CreateTask(db *gorm.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		var task Task 
		if err :=json.NewDecoder(r.Body).Decode(&task); err!= nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}	
		if err :=db.Create(&task).Error; err!=nil {
			http.Error(w, err.Error(),http.StatusInternalServerError);
			return 
		}	
		w.WriteHeader(http.StatusCreated); 
		json.NewEncoder(w).Encode(task); 

	}
}
func GetTask(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		taskID,err := strconv.Atoi(chi.URLParam(r,"id"))
		if err !=nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return 
		}
		var task Task; 
		if err := db.First(&task, taskID).Error; err != nil {
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "Record not found", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
json.NewEncoder(w).Encode(task)

	}
}
func UpdateTask(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		taskID,err := strconv.Atoi(chi.URLParam(r,"id"))
		if err !=nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		var task Task;
		if err := db.First(&task, taskID).Error; err != nil {
	if err == gorm.ErrRecordNotFound {
		http.Error(w, "Record not found", http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
if err :=json.NewDecoder(r.Body).Decode(&task); err!= nil {
	http.Error(w,err.Error(),http.StatusBadRequest)
	return
}
if err :=db.Save(&task).Error; err!=nil {
	http.Error(w, err.Error(),http.StatusInternalServerError);
	return
}
json.NewEncoder(w).Encode(task); 
	}
}
func DeleteTask(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        var task Task
        if err := db.First(&task, taskID).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                http.Error(w, "Record not found", http.StatusNotFound)
            } else {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }
        if err := db.Delete(&task).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}
