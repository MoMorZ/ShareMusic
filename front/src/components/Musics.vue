<template>
    <div class="Musics">
        <el-container class="Container">
            <el-aside width="65%" class="List">
                <Listm></Listm>
            </el-aside>
            <el-main>
                <Lyrics></Lyrics>  
            </el-main>
        </el-container>
    </div>
</template>

<script>
import Listm from './List.vue'
import Lyrics from './Lyrics'
export default {
    components:{
        Listm,
        Lyrics
    }
}
</script>

<style>
.List {
    margin-top: 16px;
    overflow: auto;
}
.Musics:before {
    background: url('~@/img/enter_night.jpg');
    background-size: 100% 100%;
    height: 85%;
    position: absolute;
    width: 100%;
    margin-left: -8px;
    opacity: 0.7;
    content: "";
    z-index: -1;
    margin-top: 1.1%;
}
</style>