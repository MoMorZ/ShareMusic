<template>
  <div class="Play">
    <div class="header">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/' }">{{back}}</el-breadcrumb-item>
        <el-breadcrumb-item>歌曲列表</el-breadcrumb-item>
        <el-breadcrumb-item>房间号: {{number}}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="middle">
      <Musics></Musics>
    </div>
    <div class="bottom">
      <Player></Player>
    </div>
  </div>
</template>

<script>
import Musics from "../components/Musics.vue";
import Player from "../components/Player.vue";
export default {
  components: {
    Musics,
    Player
  },
  data() {
    return {
      number: 123,
      back: "<-返回"
    };
  },
  methods: {},
  created: function() {
    if (this.$store.state.mstate == false) {
      this.$router.replace("/");
    }
    this.number = this.$store.state.number;
  },
  destroyed: function() {
    this.$store.commit("setMstate", false);
    this.$store.commit("setPlaystate", false);
    this.$store.commit("setNowmusic", -1);
    this.$axios({
      method: "get",
      url:
        "http://192.168.1.105:3000/sharemusic/left?id=" +
        this.$store.state.number
    }).then(res => {
      console.log(res);
    });
  }
};
</script>

<style>
.Play {
  padding: 0px;
}
.header {
  background-color: white;
  margin-top: 15px;
  height: 17px;
}
.bottom {
  position: static;
  bottom: 0px;
}
</style>
