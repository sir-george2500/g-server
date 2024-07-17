[Watch the demo video](https://www.youtube.com/watch?v=xxoy7Iq_UH8)


# Why I Built This  

I'm a reader who loves coding blogs, but there's one big problem: I hate bouncing between tabs just to catch the latest posts on my favorite blogs. Enter `g-server` - a full-blown RSS feed aggregator. What does that mean? Well, it means if I register my favorite blogs with this server, it goes out and fetches the latest 10 posts from each of them. That way, I can read all the blog posts in a single app without tab-hopping like a caffeine-fueled maniac.

# Technologies Used to Build This

This server is written purely in Golang and uses PostgreSQL as its database.

# Why Go

Initially, I decided to build this with **Rust**, but the problem was Rust's compile times were making my stomach churn. 

Secondly, **Golang** handles concurrency in a way that's easier and more efficient than Rust. Maybe I'll change my mind about this later, but for now, given my current skill level, I believe Go is the way to go. I needed a way to scrape 10 posts from different blogs at the same time, and Go did this well and fast.

# PostgreSQL

The fact that this project is based on relationships—linking users to feeds and then to their posts—shows why we went for PostgreSQL over some NoSQL database. 

for my database client `psql` commandline tools was what I use instead of the Gui like pgadmin.

# Sqlc and goose 

Goose was use to run the migration 

Sqlc was use to converte raw sql queries to type safe go code

So there you have it! No more tab-hopping, just pure, uninterrupted reading bliss.

To build the executable run 
```bash 
go build && ./g-server
```

To view the api Docs please go to 
`http://localhost:8080/swagger/index.html#/`

Remember to add your  environment variables
```bash
PORT = 
DBURL = 
```
---
A deploy version of this project is at [g-server]("https://g-server-7fg9.onrender.com/swagger/index.html#/")

License under [MIT](https://opensource.org/licenses/MIT). This code is free to use and modify. If you like, make money out of it, even rob it as grease on your skin. I give you such permission hahaha.
