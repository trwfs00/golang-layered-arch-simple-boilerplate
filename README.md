# Go Fiber Layered Architecture Project
### Project Overview
This project is a boilerplate for creating a Go-based web service using the Go Fiber framework with a layered architecture, consisting of:

- Handlers: Handles HTTP requests and responses.
- Services: Contains the business logic.
- Repositories: Interacts with the database.
- Database: Uses GORM for ORM with PostgreSQL.
- Environment Configuration: Uses Viper for loading environment variables.
- Code Generation: Uses GORM Gen to generate models and queries.

### Project Structure
```bash
├── api/
│   ├── handler/
│   │   └── user/
│   │       ├── createUser.go
│   │       ├── userById.go
│   │       └── user.go
│   ├── repository/
│   │   └── users.go
│   ├── service/
│   │   └── user/
│   │       ├── command/
│   │       │   └── create.go
│   │       ├── query/
│   │       │   └── userById.go
│   ├── router.go
├── lib/
│   ├── database/
│   │   ├── dump/
│   │   │   └── dump-test_db-202410010149.dump
│   │   ├── entity/
│   │   │   └── users.gen.go
│   │   ├── db.go
│   │   ├── gen.tool.yaml
│   └── environment/
│       └── env.go
├── main.go
├── app.env
├── gen_models.sh
├── go.mod
└── README.md
```

### Key Components
- **Handlers** (api/handler): Define the HTTP handlers.
- **Services** (api/service): Contains the business logic and queries.
- **Repositories** (api/repository): Manages data access, performs CRUD operations.
- **Database** (lib/database): Manages database connection and migrations.
- **Environment** (lib/environment): Manages environment variables with Viper.
- **Migration** (lib/database/gen.tool.yaml): Contains auto-generated yaml config using GORM Gen.

### Prerequisites
- Go 1.18+
- PostgreSQL
- GORM
- Fiber
- Viper
- GORM Gen

### Setup Instructions
#### Step 1: Clone the Repository
Clone the repository to your local machine:
```git
$ git clone <repository_url>
$ cd <repository_name>
```
#### Step 2: Install Dependencies
Install the Go dependencies using:
```
$ go mod tidy
```
#### Step 3: Environment Configuration
Create an .env file at the root of your project to manage environment variables. Below is an example of the contents:
```env
ENV=development
SERVICE_PORT=9090
DB_DSN=host=localhost user=postgres password=your_db_password dbname=your_db_name port=5432 sslmode=disable
```
#### Step 4: Generate Models Using GORM Gen
To generate models and query helpers using gorm gen, run:
```
$ gentool -c "./lib/database/gen.tool.yaml"
```
This script will load your .env file, update the configuration, and generate models files under `lib/database/entity`.
#### Step 5: Run the Application
Run the application using:
```
$ go run main.go
```
The server will start at the port specified in the .env file (default is 9090).

#### Step 6: Access the API
The following API endpoints are available:
- **Get User by ID**: `GET /api/v1/user/:id`
- **Create User**: `POST /api/v1/user`

#### Creating a User
To create a new user, you can send a POST request to the `createUser` endpoint:
```json
{
    "name": "John Doe",
    "phone": "000-000-0000"
}
```

### Project Highlights
- Fiber Framework: Lightweight, fast web framework inspired by Express.js.
- Layered Architecture: Separation of concerns for better maintainability and testability.
- Environment Management: Viper is used to handle configuration across environments.
- GORM ORM: Provides an easy way to work with databases in Go.
- GORM Gen Tool: Automates the generation of models and queries, speeding up development.

### Configuration
The configuration is managed through an environment file (app.env) loaded using Viper. You can add or modify environment variables as needed:
- `ENV`: Environment name (`development`, `production`, etc.).
- `SERVICE_PORT`: Port on which the server will run.
- `DB_DSN`: Database connection string (DSN).

### Generating Models with GORM Gen
We use `gorm.io/gen` to generate models from the database. The configuration is handled through a YAML file:

1. **Edit** `lib/database/gen.tool.yaml` to update table information.
2. **Run** `$ gentool -c "./lib/database/gen.tool.yaml"` to generate models automatically.

### Dependencies
- [Fiber](https://gofiber.io/) - HTTP Web Framework.
- [GORM](https://gorm.io/) - ORM Library for Go.
- [Viper](https://github.com/spf13/viper) - Go configuration management.
- [GORM Gen](https://github.com/go-gorm/gen) - Tool to generate models and query builders for GORM.

### Running PostgreSQL with Docker
To run PostgreSQL in a Docker container, follow these steps:

#### Prerequisites
1. Docker: Ensure you have Docker installed on your machine. You can download and install Docker from Docker's official website.

#### Step 1: Pull the PostgreSQL Image
You can pull the official PostgreSQL Docker image using the following command:
```
$ docker pull postgres
```
This command downloads the latest version of the PostgreSQL image from Docker Hub. If you want a specific version, you can specify the version tag. For example, to pull version 15.2, use:
```
$ docker pull postgres:15.2
```

#### Step 2: Run the PostgreSQL Container
Once the image is downloaded, you can run a PostgreSQL container using the following command:
```
$ docker run --name myPostgreSQL -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=your_db_password -e POSTGRES_DB=test_db -p 5432:5432 -d postgres
```
In this command:
- `--name myPostgreSQL`: Names the container `myPostgreSQL`.
- `-e POSTGRES_USER=postgres`: Sets the PostgreSQL user.
- `-e POSTGRES_PASSWORD=your_db_password`: Sets the password for the PostgreSQL user.
- `-e POSTGRES_DB=test_db`: Creates a default database named mydb.
- `-p 5432:5432`: Maps port 5432 on your host to port 5432 in the container (the default PostgreSQL port).
- `-d`: Runs the container in detached mode.

#### Step 3: Verify the PostgreSQL Container is Running
To check if your PostgreSQL container is running, use the following command:
```
$ docker ps
```
This command lists all running containers. You should see `myPostgreSQL` in the list.

#### Step 4: Connect to PostgreSQL
You can connect to the PostgreSQL database running in the Docker container using a PostgreSQL client. Here’s how to connect using the command line:
```
$ docker exec -it myPostgreSQL psql -U postgres -d test_db
```
In this command:
- `docker exec -it myPostgreSQL`: Executes a command in the running my_postgres container.
- `psql -U postgres -d test_db`: Connects to the mydb database as the myuser user.

#### Step 5: Update Your Application Configuration
Make sure to update your application’s `.env` file with the correct connection details. If you're running your application on the same machine as your Docker container, use the following connection string format:
```
DB_DSN=host=localhost user=postgres password=your_db_password dbname=test_db port=5432 sslmode=disable
```

#### Step 6: Stop and Remove the PostgreSQL Container
To stop the PostgreSQL container, use:
```
$ docker stop myPostgreSQL
```
To remove the container, run:
```
$ docker rm myPostgreSQL
```


### Contributing
Feel free to fork the repository, create a new branch, and submit a pull request. Contributions are always welcome.

### License
This project is licensed under the MIT License.