<template>
  <div id="app">
    <nav class="navbar navbar-light bg-light">
      <span class="navbar-brand mb-0 h1">ToDo List with Docker,Vue,Golang & MongoDB</span>
    </nav>
    <todo-creator @ItemCreated="RefreshItems"></todo-creator>
    <to-do-list :items="items" @ItemDeleted="RefreshItems"></to-do-list>
  </div>
</template>

<script>
import TodoCreator from "./components/TodoCreator";
import ToDoList from "./components/ToDoList";
import axios from "axios";
export default {
  name: 'App',
  components: {
    TodoCreator,
    ToDoList
  },
  data:function(){
    return {
      items:[]
    }
  },
  methods:{
    async RefreshItems(){
      var result = await axios.get("http://localhost:8081/items");
      this.items = result.data
    }
  },
  created() {
    this.RefreshItems();
  }
}
</script>

<style>

</style>
