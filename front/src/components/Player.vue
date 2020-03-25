<template>
  <div class="Player">
    <div class="Components">
      <el-image class="Img" :src="picchange"></el-image>
      <el-slider
        class="Slider"
        v-model="slidertime"
        :format-tooltip="FormatProcessToolTip"
        @change="ChangeTime"
      ></el-slider>
      <div class="Text">{{timechange}}</div>
      <div class="Button">
        <el-button type="primary" icon="el-icon-arrow-left" circle @click="Before"></el-button>
        <el-button v-if="playchange" type="primary" circle @click="ChangeState">| |</el-button>
        <el-button v-else type="primary" icon="el-icon-caret-right" circle @click="ChangeState"></el-button>
        <el-button type="primary" icon="el-icon-arrow-right" circle @click="Next"></el-button>
      </div>
      <div class="Name">{{namechange}}</div>
      <audio
        ref="audio"
        @pause="onPause"
        @play="onPlay"
        @timeupdate="onTimeupdate"
        @loadedmetadata="onLoadedmetadata"
        preload="auto"
        autoplay
        :src="musicchange"
      ></audio>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      img_url: require("../img/enter_pm.jpg"),
      state: false,
      Name: "Test",
      music_url: "",
      nownum: -1,
      slidertime: 0,
      audio: {
        // 该字段是音频是否处于播放状态的属性
        playing: 0,
        currentTime: 0,
        maxtime: 0
      }
    };
  },
  computed: {
    playchange() {
      // console.log('playing change');
      var state = this.$store.state.playstate;
      if (this.audio.playing != state) {
        this.ChangeState();
      }
      return state;
    },
    picchange() {
      var num = this.$store.state.nowmusic;
      if (num != -1) return this.$store.state.namelist[num].al.picUrl;
      else return require("../img/enter_pm.jpg");
    },
    musicchange() {
      var num = this.$store.state.nowmusic;
      console.log("change");
      this.heart(true);
      if (num != -1) {
        return this.$store.state.playlist[num].url;
      } else return "";
    },
    namechange() {
      var num = this.$store.state.nowmusic;
      // console.log(this.$store.state);
      if (num == -1) return "";
      else return this.$store.state.namelist[num].name;
    },
    timechange() {
      // console.log('changetime')
      var time = this.audio.currentTime;
      return this.realFormatSecond(time);
    }
  },
  created: function() {
    console.log("create ", this.$store);
    this.heart(false);
  },
  methods: {
    onTimeupdate(res, num = -1) {
      if (num != -1) {
        res.target.currentTime = num;
      }
      // console.log(this.audio.currentTime);
      this.audio.currentTime = parseInt(res.target.currentTime);
      if (
        this.audio.currentTime == this.audio.maxtime &&
        this.audio.currentTime != 0
      ) {
        this.Next();
      }
      this.slidertime = parseInt(
        (this.audio.currentTime / this.audio.maxtime) * 100
      );
    },
    onLoadedmetadata(res) {
      // console.log(res);
      this.audio.maxtime = parseInt(res.target.duration);
    },
    Next() {
      if (
        this.$store.state.nowmusic == this.$store.state.playlist.length - 1 ||
        this.$store.state.nowmusic == -1
      ) {
        this.$store.commit("setNowmusic", 0);
      } else {
        this.$store.commit("setNowmusic", this.$store.state.nowmusic + 1);
      }
      this.$store.commit("setPlaystate", true);
      this.heart(true);
    },
    Before() {
      if (this.$store.state.nowmusic == 0 || this.$store.state.nowmusic == -1) {
        this.$store.commit(
          "setNowmusic",
          this.$store.state.playlist.length - 1
        );
      } else {
        this.$store.commit("setNowmusic", this.$store.state.nowmusic - 1);
      }
      this.$store.commit("setPlaystate", true);
      this.heart(true);
    },
    ChangeState() {
      console.log("playstate");
      if (this.audio.playing == 1) {
        this.audio.playing = 0;
        this.$refs.audio.pause();
      } else {
        if (this.$store.state.nowmusic == -1) {
          this.Next();
        }
        this.audio.playing = 1;
        this.$refs.audio.play();
      }
      this.$store.commit("setPlaystate", this.audio.playing);
      this.heart(true);
    },
    ChangeTime(index) {
      this.$refs.audio.currentTime = (index / 100) * this.audio.maxtime;
    },
    FormatProcessToolTip(index = 0) {
      index = parseInt((this.audio.maxtime / 100) * index);
      return "时间: " + this.realFormatSecond(index);
    },
    realFormatSecond(second) {
      var secondType = typeof second;

      if (secondType === "number" || secondType === "string") {
        second = parseInt(second);

        var hours = Math.floor(second / 3600);
        second = second - hours * 3600;
        var mimute = Math.floor(second / 60);
        second = second - mimute * 60;

        return (
          hours +
          ":" +
          ("0" + mimute).slice(-2) +
          ":" +
          ("0" + second).slice(-2)
        );
      } else {
        return "0:00:00";
      }
    },
    // 播放音频
    play() {
      this.$refs.audio.play();
    },
    // 暂停音频
    pause() {
      this.$refs.audio.pause();
    },
    // 当音频播放
    onPlay() {
      this.audio.playing = 1;
    },
    // 当音频暂停
    onPause() {
      this.audio.playing = 0;
    },
    heart(mode = true) {
      if (this.$store.state.mstate == false) {
        return;
      }
      // var random = require('../random')
      // var time = require('../time');
      // var nowtime = time()
      console.log("heart");
      let _this = this;
      var turl =
        "http://192.168.1.105:3000/sharemusic/heart?id=" +
        this.$store.state.number +
        "&time=0&mstate=" +
        this.audio.playing +
        "&nowmusic=" +
        this.$store.state.nowmusic;
      if (mode == true) {
        turl =
          "http://192.168.1.105:3000/sharemusic/heart?id=" +
          this.$store.state.number +
          "&time=" +
          (parseInt(this.$store.state.time) + 1).toString() +
          "&mstate=" +
          this.audio.playing +
          "&nowmusic=" +
          this.$store.state.nowmusic;
      }
      console.log(turl);
      this.$axios({
        method: "get",
        url: turl
      })
        .then(res => {
          console.log(res);
          // _this.$store.commit('setMstate', res.data.state);
          if (res.data.state == "500") {
            _this.$router.replace("/");
          }

          if (res.data.playstate != -1) {
            _this.$store.commit("setNowmusic", res.data.nowmusic);
            _this.$store.commit("setPlaystate", res.data.playstate);
          }
          _this.$store.commit("setTime", res.data.time);
          if (mode == false) {
            setTimeout(function() {
              _this.heart(false);
            }, 2000);
          }
        })
        .catch(res => {
          console.log(res);
        });
    }
  }
};
</script>

<style >
.Player {
  background-color: rgb(56, 6, 80);
  background-size: 100% 100%;
  height: 9%;
  position: absolute;
  width: 100%;
  margin-left: -8px;
  opacity: 1;
  content: "";
  z-index: -1;
  bottom: 0px;
}
.Img {
  position: absolute;
  left: 10%;
  width: 70px;
  height: 70px;
  bottom: 0px;
}
.Slider {
  position: absolute;
  width: 600px;
  left: 240px;
  bottom: 25px;
}
.Text {
  position: absolute;
  color: azure;
  left: 810px;
  bottom: 5px;
}
.Button {
  position: absolute;
  left: 860px;
  bottom: 20px;
}
.Name {
  color: aliceblue;
  position: absolute;
  left: 1150px;
  bottom: 30px;
}
</style>