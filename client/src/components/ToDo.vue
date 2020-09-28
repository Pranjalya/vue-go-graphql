<template>
  <div class="main">
    <h1>ToDo Application</h1>
    <div class="form-container">
      <form>
        <input type="text" class="todo-input" v-model="text" />
        <button
          class="todo-btn"
          id="plus-btn"
          type="submit"
          @click.prevent.stop="addTodo"
        >
          <i class="fas fa-plus"></i>
        </button>
      </form>
      <p v-if="getWarning" class="alert-text">Todo note can not be empty!</p>
    </div>
    <div class="todo-container">
      <div v-for="todo in todos" v-bind:key="todo.id">
        <div :class="todo.done ? 'todo-item todo-item-checked' : 'todo-item'">
          <p>{{ todo.text }}</p>
          <button id="checked-btn" class="todo-btn">
            <i class="fas fa-clipboard-check"></i>
          </button>
          <button id="delete-btn" class="todo-btn">
            <i class="far fa-trash-alt"></i>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import TodoActions from "../services/TodoActions";

@Options({
  props: {
    username: String,
  },
})
export default class ToDo extends Vue {
  username!: string;
  text!: string;
  warning = false;
  todos = [];

  public fetchAllItems(): void {
    TodoActions.queryTodoItems()
      .then((result) => {
        console.log("All todos", result.data.data.todos);
        this.todos = result.data.data.todos;
      })
      .catch((err) => console.log);
  }

  public addTodo(): void {
    if (this.text == undefined) {
      this.warning = true;
    } else {
      this.warning = false;
      TodoActions.addTodoItem(this.text, this.username)
        .then((result) => {
          console.log("Data", result.data.data.createTodo);
          this.fetchAllItems();
        })
        .catch((err) => console.log);
    }
  }

  get getWarning(): boolean {
    return this.warning;
  }

  mounted() {
    this.fetchAllItems();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

.main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

.form-container {
  margin-top: 1rem;
  width: 100%;
  height: 5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
form {
  background: #fff;
  height: 3rem;
  width: 40%;
  transition: all ease-in-out 0.3s;
  position: relative;
}
.form-alert {
  border: red 3px dotted;
}
.alert-text {
  color: red;
  font-size: 1.2rem;
  padding: 0.3rem;
}

.todo-input {
  margin: 0 auto;
  width: 100%;
  height: 2.5rem;
  font-size: 1.5rem;
  border: none;
  outline: none;
}
.todo-btn {
  height: 2rem;
  width: 2rem;
  border: none;
  font-size: 1.5rem;
  background: #000082;
  color: #fff;
  cursor: pointer;
  transition: all ease-in-out 0.2s;
}
#plus-btn {
  position: absolute;
  right: 0.5rem;
  top: 0.5rem;
}
.todo-btn:hover {
  background: #ff6b24;
  transform: scale(1.1);
}
select {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  outline: none;
  border: none;
}
.select {
  position: absolute;
  right: 0.5rem;
  top: 0.5rem;
}
select {
  background: #1d4e89;
  width: 7rem;
  height: 2rem;
  text-align: center;
  font-size: 1rem;
  color: #fff;
  cursor: pointer;
  padding: 0 0 0 0.5rem;
}
.select::after {
  content: "\25BC";
  position: absolute;
  color: #fff;
  padding: 1rem 0.5rem 0 0;
  top: 0.5rem;
  right: 0;
}

.todo-container {
  margin-top: 1rem;
  width: 40%;
  transition: all ease-in-out 0.3s;
}

.todo-item {
  background: #1d4e89;
  display: flex;
  align-items: center;
  padding: 0.5rem;
  margin-bottom: 0.5rem;
  transition: all ease-in-out 0.5s;
}
.todo-item-checked {
  text-decoration: line-through;
  text-decoration-thickness: 2px;
  text-decoration-color: red;
  opacity: 0.7;
}
.todo-item p {
  font-size: 1.5rem;
  color: #fff;
  width: 90%;
}
.todo-item button {
  margin-left: 0.5rem;
}
#delete-btn:hover {
  background: #f00000;
}
#checked-btn:hover {
  background: #008000;
}
.fa-trash-alt,
.fa-clipboard-check {
  pointer-events: none;
}
.fall {
  opacity: 0;
  transform: scale(0.5);
}

@media screen and (max-width: 992px) {
  form,
  .todo-container {
    width: 60%;
  }
}
@media screen and (max-width: 600px) {
  form,
  .todo-container {
    width: 80%;
  }
}
</style>
