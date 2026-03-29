# 🎬 MyCineDen

> Your cozy corner for every film you've ever loved — log what you've watched, rate it your way, and keep track of what's next.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go&logoColor=white)
![Svelte](https://img.shields.io/badge/Svelte-5.x-FF3E00?style=flat&logo=svelte&logoColor=white)
![TMDB](https://img.shields.io/badge/TMDB-API-01B4E4?style=flat&logo=themoviedatabase&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green?style=flat)

---

## 🍿 What is MyCineDen?

MyCineDen is a personal movie tracking web app — a fun side project built with **Go** on the backend and **Svelte** on the frontend. It's your own cozy den where every film you've ever watched has a place.

No social feeds. No critic scores. No noise.
Just **your movies, your ratings, your den.**

---

## ✨ Features

### 🎥 Watched Movies
- Log every movie you've watched
- Give it your own personal rating (1–10)
- Add personal notes or a short review
- See your full watched history at a glance
- Movie posters and details auto-fetched via TMDB

### 📋 Watchlist
- Add movies you're planning to watch
- Keep your "to-watch" queue organized
- Move a movie from Watchlist → Watched once you've seen it

### 🔍 Search & Discover
- Search movies in **any language**
- Powered by the TMDB API
- Auto-complete with posters, year, and language info

### 🎨 Clean, Fast UI
- Built with Svelte — snappy, no bloat
- Minimal and focused interface
- Responsive design for desktop and mobile

---

## 🛠️ Tech Stack

| Layer | Technology |
|-------|-----------|
| **Backend** | Go (Golang) |
| **Frontend** | Svelte + SvelteKit |
| **Database** | SQLite |
| **Movie Data** | TMDB API |
| **API** | REST / JSON |

---

## 🌐 Movie Data — TMDB API

MyCineDen uses the [TMDB (The Movie Database)](https://www.themoviedb.org/documentation/api) API to fetch movie details, posters, and metadata.

- 🆓 Free for non-commercial use
- 📸 Provides posters, backdrops, cast, genres, release dates

To use it, grab a free API key at [themoviedb.org](https://www.themoviedb.org/signup) and add it to your `.env`:

```env
TMDB_API_KEY=your_api_key_here
```

---

## 🗺️ Roadmap

- [x] Project overview
- [ ] Watched movies list with personal rating
- [ ] Add / edit / delete watched movies
- [ ] Watchlist management
- [ ] TMDB search integration (multilingual)
- [ ] Movie posters & details auto-fetch
- [ ] Search and filter watched list
- [ ] Dark / light theme toggle

---

## 📄 License

MIT — do whatever you want with it.

---

*Built with ☕ and way too many late-night movies.*
