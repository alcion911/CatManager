# Cat Manager

Cat Manager is an application that allows you to manage a list of animals and displays a carousel of cat images obtained from an external API. The application is divided into two parts: a backend built with Go and a frontend built with React.

## Project Structure
CatManager/ ├── animal-crud/ # Backend in Go │ ├── Dockerfile # Dockerfile for the backend │ ├── go.mod # Go modules file │ ├── go.sum # Go modules checksum file │ ├── main.go # Entry point of the application │ ├── routes/ # API routes │ │ └── animal_routes.go # Routes for managing animals └── animal-crud-frontend/ # Frontend in React ├── Dockerfile # Dockerfile for the frontend ├── package.json # Dependencies and scripts for React ├── src/ # Source code of the React application │ ├── components/ # React components │ │ ├── AnimalForm.js # Component for adding/editing animals │ │ ├── AnimalList.js # Component for listing animals │ │ └── CatCarousel.js # Component for displaying cat images │ ├── context/ # Context for state management │ │ └── AnimalContext.js # Context for animals │ ├── App.js # Main component of the application │ └── index.js # Entry point of React └── docker-compose.yml # File to orchestrate the containers


## Prerequisites

- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/) installed on your machine.
- Internet connection to access the external cat API.

## Installation and Execution

### Running the Application with Docker

1. **Clone the repository**:

   ```bash
   git clone https://github.com/alcion911/CatManager.git
   cd CatManager
Build and start the containers:

Make sure you are in the root of the project (where the docker-compose.yml file is located) and run:

bash
docker-compose up --build

This command will build the Docker images for the backend and frontend and start the containers.

Access the application:

The frontend will be available at http://localhost:3000.
The backend API will be available at http://localhost:8080.
Implementation Details
Backend (Go)
Technologies: Go, Gin (web framework), PostgreSQL (database).

Routes:

- GET /animals: Retrieves the list of animals.
- POST /animals: Adds a new animal.
- PUT /animals/:id: Updates an existing animal.
- DELETE /animals/:id: Deletes an animal.
- GET /cat-images: Retrieves cat images from the external API (The Cat API).
- CORS: CORS has been configured to allow requests from the frontend.

##Frontend (React)
Technologies: React, Axios (for HTTP requests), react-slick (for the carousel).

Components:

AnimalForm: Component for adding and editing animals.
AnimalList: Component for listing existing animals.
CatCarousel: Component that displays a carousel of cat images obtained from the external API.
Global State: The React Context API is used to manage the state of animals in the application.

Contributions
If you would like to contribute to this project, feel free to open an issue or submit a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for more details.


### Clarification of AWS Integration
AWS supports docker container integration, so the next step for integration with AWS is to upload the docker containers to AWS using Amazon EC2 to lift the containers in a linux environment, and the DB would be lifted using Amazon RDS.

