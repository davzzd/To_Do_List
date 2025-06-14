# To-Do List Application
A simple To-Do List application built with Go and Templ.

# Prerequisites
- Go 1.24.4 or higher
- Git (for cloning the repository)

# Dependencies
The project uses the following main dependencies:
- [github.com/a-h/templ](https://github.com/a-h/templ) v0.3.898 - For templating and UI components

# Installation
1. Clone the repository:
```bash
git clone https://github.com/davzzd/To_Do_List.git
cd to-do-list
```

2. Install dependencies:
```bash
go mod download
```

3. Generate templ files (if you make changes to .templ files):
```bash
templ generate
```

# Running the Application
1. Start the server:
```bash
go run main.go
```

2. Open your browser and navigate to:
==>  http://localhost:8080


# Testing
Run the tests using:
```bash
go test -v
```

