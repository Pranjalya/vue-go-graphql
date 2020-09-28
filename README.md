# Vue-Go-GraphQL

A full stack web application demonstrating a ToDo list maker application using Golang as the backend server (specifically gqlgen) and Vue.js as frontend.


## Project setup
```
cd client
yarn install
```

### Compiles and hot-reloads for development
1. Start the GraphQL server
```
go run server.go
```
2. Start the frontend of the web application
```
yarn serve
```
3. Open `localhost:8080` for the frontend and/or `localhost:7979` to play with the queries

### Compiles and minifies for production
```
yarn build
```

### Run your unit tests
```
yarn test:unit
```

### Run your end-to-end tests
```
yarn test:e2e
```

### Lints and fixes files
```
yarn lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
