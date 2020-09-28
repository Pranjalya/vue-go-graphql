import http from '../http-path'

class TodoActions {
    addTodoItem(text: string, id: string) {
        console.log("Adding text : ", text);
        return http.post('/query', {
            query: `
            mutation createTodo {
                createTodo(input:{text:"${text}", userId:"${id}"}) {
                  user {
                    id
                  }
                  text
                  done
                }
              }
            `
        })
    }
    queryTodoItems() {
        return http.post('/query', {
            query: `
            query findTodos {
                todos {
                  text
                  done
                  user {
                    name
                  }
                }
            }
            `
        })
    }
}

export default new TodoActions;
