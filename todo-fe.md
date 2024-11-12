# Gober FE System Design

A good alternative to React for a beginner-friendly frontend framework is **Vue.js**. Vue is known for its simplicity, ease of learning, and clean syntax, making it suitable for your project. Here's a high-level system design using **Vue.js**:

### 1. **Frontend Tech Stack**
   - **Framework**: **Vue.js** (beginner-friendly, approachable, and has a similar component-based structure to React).
   - **Styling**: **Vue’s Scoped CSS** (built-in option for CSS styling scoped to each component) or **Tailwind CSS** for a utility-based styling approach.
   - **HTTP Requests**: **Axios** for API requests.

### 2. **Frontend Structure**
   - **Home Page**: Lists popular articles by calling `/articles/popular?source=website_name`.
   - **Article Detail Page**: Displays detailed information for an article, calling `/article?detailUrl=encoded_url`.
   - **Search Functionality**: A search bar on the Home Page where users can search articles using keywords. The search results populate in the article list using `/articles?source=website_name&q=search_keyword`.

### 3. **Vue Component Breakdown**
   - **App Component**: Main entry point where routing is set up.
   - **ArticleList Component**: Shows a list of articles, fetching either popular articles or search results depending on the input.
   - **ArticleDetail Component**: Shows the details of a selected article.
   - **SearchBar Component**: Provides a search bar for users to type keywords.

### 4. **Routing and State Management**
   - **Routing**: Use **Vue Router** to navigate between the Home Page and the Article Detail Page.
   - **State Management**: Manage component states with **Vue’s reactivity system** and **Vue’s Composition API** for handling data within each component.

### 5. **Design Flow**
   - **Home Page**: Loads popular articles by default.
      - **Search Bar**: When a search query is submitted, `/articles?source=website_name&q=search_keyword` is called, and the list updates with search results.
   - **Article Click**: Clicking an article navigates to its detail page.
   - **Article Detail Page**: Calls `/article?detailUrl=encoded_url` to fetch the article's full details.

### Example Vue Component Structure
```plaintext
App Component
│
├── HomePage Component
│   ├── SearchBar Component
│   └── ArticleList Component
│
└── ArticleDetailPage Component
    └── ArticleDetail Component
```

This setup with Vue.js remains simple and beginner-friendly, giving you a strong base for future functionality. And, as a reminder, don't forget about **point 9: Store Results in a Database** for backend data persistence when you’re ready.