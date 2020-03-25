<template>
    <div class="Init">
        <div class="Title">
            ShareMusic
        </div>
        <div class="table">
            <Table></Table>
        </div>
    </div>
</template>

<script>
import Table from '../components/Table.vue'
export default {
    components: {
        Table
    }
}
</script>

<style>
.Init {
    background: url('~@/img/enter_pm.jpg');
    background-size: 100% 100%;
    height: 100%;
    position: fixed;
    width: 100%;
    margin-left: -8px;
    margin-top: -10px;
}
.Title {
    text-align: center;
    margin-top: 9%;
    font-size: 300%;
    font-family: 'Comic Sans MS', cursive, sans-serif;
}
.table {
    margin: 0 auto;
    height: 10%;
    width: 30%;
    margin-top: 3%;
}
</style>