# tooly
# tooly

Tooly is a versatile Go-based web application that provides various utilities including arithmetic operations, financial calculations, note-taking functionality, and image generation.

## Features

- Arithmetic operations (percentage calculations, averages)
- Financial tools (SIP calculator)
- Note-taking system with CRUD operations
- Image generation (AI-based)
- Logging middleware
- PostgreSQL database integration

## Project Structure

- `handlers/`: Contains request handlers for different functionalities
  - `arithmetic/`: Percentage and average calculations
  - `finance/`: SIP calculator
  - `notes/`: CRUD operations for notes
  - `image/`: AI-based image generation
   - `taskmanager/`: Task manager for managing several tasks . 
- `middlewares/`: Custom middleware (logging)
- `db/`: Database connection and configuration
- `Dockerfile`: PostgreSQL container configuration

## Setup and Installation

### Prerequisites

- Go 1.x
- PostgreSQL
- Docker 

### Database Setup

1. Use the provided Dockerfile to set up a PostgreSQL container:

   ```bash
   docker build -t tooly-postgres -f Dockerfile .
   docker run -d -p 5432:5432 --name tooly-db tooly-postgres



Configure the database connection in your application:

dsn := "host=localhost user=youruser password=yourpassword dbname=database port=5432 sslmode=disable"
database.ConnectDB(dsn, &notes.Notes{})



Usage
Arithmetic Operations
Percentage Calculations
POST /calculate-percentage
Body: {"A": 25, "B": 100}



Average Calculation
POST /average
Body: [1, 2, 3, 4, 5]



Financial Tools
SIP Calculator
POST /sip
Body: {
  "MonthlyInvestment": "1000",
  "Duration": "12",
  "Interest": "0.08",
  "Total": "0"
}



Note-taking System
POST /notes       // Create a new note
GET /notes        // Retrieve all notes
GET /notes/{id}   // Retrieve a specific note
PUT /notes/{id}   // Update a note
DELETE /notes/{id} // Delete a note



Image Generation
POST /generate-image
Body: "A futuristic cityscape"



Middleware
The application includes a logging middleware that logs request details and response time:

r.Use(middlewares.LoggerMiddleware)



Database Operations
The project uses GORM for database operations. Example of creating a database connection:

database.ConnectDB(dsn, &notes.Notes{})



Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

License
This project is licensed under the MIT License.