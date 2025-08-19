# Nazar

**Nazar** is a blogging platform built with Golang, MongoDB, and HTML templates.
It provides a clean and SEO-friendly way to manage and display articles with categories, slugs, and pagination support.

## Features

1. 📰 **Articles Management**
    - Add, view, and delete articles
    - Each article has a unique slug for SEO
    - Prevents duplicate article slugs

2. 🏷️ **Categories**
    - Articles can be grouped into categories
    - Category pages list all related articles

3. 🔎 **Search**
    - Full-text search support for articles by title

4. **📄 Pagination**

    - Articles and category pages come with page navigation

5. **🔒 Admin Dashboard**

    - Secure login
    - Manage articles and categories
    - Flash messages for success/error

6. **🎨 UI/UX**
    - Responsive blog cards
    - Dark theme with customizable colors via CSS variables

7. **🛠️ Tech Stack**
    - Backend: Golang (net/http, MongoDB driver)
    - Database: MongoDB
    - Frontend: HTML, CSS (Dark theme with cards layout)

## Template Engine: Go html/template

**⚙️ Installation**

Clone the repo:

git clone https://github.com/Sahil0420/Nazar.git
cd Nazar


Copy .env.example to .env and update with your settings:

DB_NAME=nazar
DB_URI=mongodb://localhost:27017
PORT=8000


Install dependencies (Go modules):

go mod tidy


Run the server:

go run main.go


Open in browser:

http://localhost:8000

📂 Project Structure
Nazar/
├── db/                # Database functions
├── handlers/          # HTTP Handlers
├── models/            # Data models
├── static/            # CSS, JS, Images
├── templates/         # HTML templates
├── main.go            # Entry point
├── go.mod
└── README.md

🔮 Roadmap

 User authentication (public users)

 Article comments

 Image uploads

 Docker support

 Deploy to cloud (Render/Heroku)

🤝 Contributing

Fork the project

Create a feature branch (git checkout -b feature/new-feature)

Commit changes (git commit -m "Add new feature")

Push to branch (git push origin feature/new-feature)

Open a Pull Request

📜 License

This project is licensed under the MIT License.