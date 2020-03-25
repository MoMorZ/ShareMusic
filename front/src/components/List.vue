<template>
    <div class="List">
        <ul style="overflow:auto">
            <p class="Music" v-for="(music, index) in musics" :key="index" @click="choose(index)">{{ music.name }}</p>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return{
            musics: [{id: 1}]
        }
    },
    methods: {
        choose(index){
            // console.log(index);
            this.$store.commit('setNowmusic', index);
            this.$store.commit('setPlaystate', true);
        }
    },
    created: function() {
        if(this.$store.state.mstate == false){
            this.$router.replace('/');
        }
        // console.log(this.$store.state.namelist)
        this.$data.musics = this.$store.state.namelist
    }
}
</script>

<style scoped>
.Music {
    height: 40px;
    background-color: rgb(126, 187, 240,0.3);
    margin-left: -4%;
    color: aliceblue;
}
.List {
    height:  600px;
}
</style>
