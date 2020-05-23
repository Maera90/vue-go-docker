<template>
    <div class="container" style="margin-top: 50px">
        <div class="row">
            <div class="col-12">
                <table class="table">
                    <thead>
                    <tr>
                        <th>#Id</th>
                        <th>Text</th>
                        <th>Text</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr v-for="item in items" :key="item._id">
                        <td>{{item._id}}</td>
                        <td>{{item.text}}</td>
                        <td>
                            <button type="button" @click="DeleteItem(item._id)" class="btn btn-danger">Del</button>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>

    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: "ToDoList",
        props:[
            'items'
        ],
        methods:{
            async DeleteItem(id){
                await axios.delete("http://localhost:8081/items/"+id);
                this.$emit("ItemDeleted");
            }
        }
    }
</script>

<style scoped>

</style>