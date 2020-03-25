<template>
    <div class="Lyrics">
        <!-- <el-card class="Box"> -->
            <div v-for="o in 4" :key="o" class="text">
                    {{'列表内容 ' + o }}
            </div>
        <!-- </el-card> -->
    </div>
</template>

<style scoped>
.Lyrics {
    background-color:rgb(255, 255, 255,0.5);
    background-size: 100% 100%;
    height: 85%;
    position: absolute;
    width: 34%;
    margin-left: -8px;
    opacity: 1;
    content: "";
    z-index: -1;
    
}
.text {
    text-align: center;
}
/* .Box {
    height: 100%;
} */
</style>