
# Forum

Rest API for Forum. Golang without dependencies (almost :D)


## Run Locally

Clone the project

```bash
  git clone https://github.com/Levap123/forum
```

Go to the project directory

```bash
  cd forum
```

Start the server

```bash
  make build && make run
```

## For Developers

**Preparing**
- [ ] Configs
- [x] Create and Connect to DB
- [x] Dockerfile

**Logic**
- [x] sign-up
- [x] sign-in
- [x] middleware (session tracker)
- [x] get user by user id
- [x] create post
- [ ] update post
- [x] get posts by user id
- [x] like or dislike psot
- [ ] create comment 
- [ ] delete comment
- [ ] like or dislike comment
- [ ] categories for post
