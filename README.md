[![Gober build status](https://github.com/akhmadreiza/gober/actions/workflows/go.yml/badge.svg)](https://github.com/akhmadreiza/gober/actions/workflows/go.yml)

# **Gober (Go Berita)**  
**Gober**, stands for `Go Berita`, is a monorepo project for aggregating news/article from various websites.  
- **Backend**: Written in Go, provides REST APIs for scraping and fetching news articles.  
- **Frontend**: Built with Vue.js, displays the news in an elegant, user-friendly web interface.

---

## **Features**  
- Scrape popular articles from multiple websites (e.g., detik.com, kompas.com).  
- Search articles by keyword. (WIP)  
- Fetch article details with **ad-free content** — scripts, iframes, and ad elements are stripped server-side before serving.  
- "Baca Juga" (related article) links inside articles are rewritten to stay within Gober instead of redirecting to the original site.  
- Editorial reading experience with clean typography (Playfair Display + Lora).  
- Responsive design for desktop and mobile.  

---

## **Available Sites**
| Sites | Status | Query Param | Origin |
| ----- | ------ | ----------- | ------- |
| detik.com | :white_check_mark: | `?source=detik` | :indonesia: |
| kompas.com | :white_check_mark: | `?source=kompas` | :indonesia: |
| tribunnews.com | :soon: | `?source=tribun` | :indonesia: |
| cnnindonesia.com | :soon: | `?source=ccnid` | :indonesia: |

#### Legend:
- :white_check_mark:: Up
- :x:: Need Fix
- :soon:: Coming Soon

---

## **Tech Stack**  
### Backend  
- Go (Golang)  
- Gin Web Framework  
- goquery (for web scraping)  
- Mockable HTTP client for testing  

### Frontend  
- Vue.js  
- Axios (for API calls)  
- CSS Grid and Flexbox for responsive layouts  

---

## **Getting Started**  

### Prerequisites  
- **Backend**:  
  - [Go 1.20+](https://golang.org/doc/install)  
  - [Git](https://git-scm.com/downloads)  

- **Frontend**:  
  - [Node.js 18+](https://nodejs.org/)  
  - [npm](https://www.npmjs.com/get-npm)  

---

## **Setup Instructions (Single server — production-like)**

### 1. **Clone the Repository**  
```bash
git clone https://github.com/your-username/gober.git
cd gober
```

### 2. **Build frontend and run**
```bash
make build-fe
go run main.go
```

Open `http://localhost:8080` in your browser. Both the frontend and API are served by the Go server on the same port.

---

## **Setup Instructions (Separate BE and FE)**

### 1. **Clone the Repository**  
```bash
git clone https://github.com/your-username/gober.git
cd gober
```

### 2. **Backend (Go)**  
#### Step-by-step:  
1. Navigate to the root directory (if not already there):  
   ```bash
   cd gober
   ```
2. Install dependencies:  
   ```bash
   go mod tidy
   ```
3. Run the backend server:  
   ```bash
   go run main.go
   ```
4. **API Endpoints**:  
   The server will run at `http://localhost:8080`. You can access the following endpoints:  
   - **Get popular articles**: `/articles/popular?source=detik`  
   - **Search articles**: `/articles?source=detik&q=keyword`  
   - **Get article details**: `/article?source=detik&detailUrl=encoded_url`
   
   See [Available Sites](#available-sites) for `source`.

---

### 3. **Frontend (Vue.js)**  
#### Step-by-step:  
1. Navigate to the `web` directory:  
   ```bash
   cd web
   ```
2. Install dependencies:  
   ```bash
   npm install
   ```
3. Run the development server:  
   ```bash
   npm run serve
   ```
4. Open your browser and visit `http://localhost:8081` to view the web application.

---

## **Project Structure**
```
gober/
├── parsers/                # Parsers for different news websites
├── models/                 # Data models
├── utils/                  # Utilities (e.g., HTTP client, helper functions)
├── main.go                 # Entry point for the backend server
├── web/                    # Frontend codebase (Vue.js)
│   ├── src/                # Source code
│   │   ├── components/     # Vue components
│   │   ├── views/          # Application views
│   │   └── assets/         # Static assets (CSS, images, etc.)
├── go.mod                  # Backend dependencies file
├── go.sum                  # Backend dependency checksums
├── README.md               # Project documentation
└── .gitignore              # Ignored files for Git
```

---

## **Testing**
### Backend Testing  
1. Navigate to the root directory:  
   ```bash
   cd gober
   ```
2. Run tests:  
   ```bash
   go test ./...
   ```

---

## **Contributing**  
We welcome contributions!  
- Fork the repository  
- Create a new branch (`feature/my-feature`)  
- Submit a pull request  

---

## **License**  
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
