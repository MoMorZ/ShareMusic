<template>
  <div class="Table">
    <el-card class="card">
      <el-collapse class="menu" v-model="activename" accordion>
        <el-collapse-item name="create" title="创建房间">
          <el-form>
            <el-form-item label="歌单链接">
              <el-input v-model="url"></el-input>
            </el-form-item>
          </el-form>
          <el-row>
            <el-col :span="4">
              <el-button @click="clear('0')">清空</el-button>
            </el-col>
            <el-col :span="4" :offset="15">
              <el-button @click="CreateSure()">确定</el-button>
            </el-col>
          </el-row>
        </el-collapse-item>
        <el-collapse-item name="enter" title="进入房间">
          <el-form>
            <el-form-item label="房间号">
              <el-input v-model="number"></el-input>
            </el-form-item>
          </el-form>
          <el-row>
            <el-col :span="4">
              <el-button @click="clear('1')">清空</el-button>
            </el-col>
            <el-col :span="4" :offset="15">
              <el-button @click="Numbersure()">确定</el-button>
            </el-col>
          </el-row>
        </el-collapse-item>
      </el-collapse>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activename: "create",
      url: "",
      number: ""
    };
  },
  methods: {
    getMore(res) {
      console.log("getMore1 ", res);
      let _this = this;
      _this.$store.commit("setNumber", res.data.number);
      _this.$store.commit("setMstate", true);
      var allUrl = "";
      for (let i = 0; i < res.data.list.length; i++) {
        allUrl = allUrl + res.data.list[i].id + ",";
      }
      allUrl = allUrl.substring(0, allUrl.length - 1);
      // 获取列表所有歌曲的播放链接
      _this
        .$axios({
          method: "get",
          url: "http://192.168.1.105:3000/song/url?id=" + allUrl
        })
        .then(function(res2) {
          var result = res2.data.data;
          console.log("getMore2 ", result);
          result.sort(function(a, b) {
            return a.id - b.id;
          });
          _this.$store.commit("setPlaylist", result);
        });
      // 获取所有歌曲的详细信息
      _this
        .$axios({
          method: "get",
          url: "http://192.168.1.105:3000/song/detail?ids=" + allUrl
        })
        .then(function(res3) {
          var result = res3.data.songs;
          console.log("getMore3 ", result);
          result.sort(function(a, b) {
            return a.id - b.id;
          });
          _this.$store.commit("setNamelist", result);
          // 成功消息展示
          _this.$message({
            showClose: true,
            message: "解析成功",
            type: "success"
          });
          _this.$router.push("/player");
        });
    },
    CreateSure() {
      var pos0 = this.url.indexOf("?id=") + 3;
      var pos1 = this.url.indexOf("&");
      if (pos1 == -1) {
        pos0 = this.url.indexOf("playlist/") + 8;
        pos1 = this.url.indexOf("/", pos0 + 1);
      }
      var listnum =
        "http://192.168.1.105:3000/sharemusic/list?id=" +
        this.url.substr(pos0 + 1, pos1 - pos0 - 1);
      // console.log(listnum);
      let _this = this;
      this.$message({
        showClose: true,
        message: "解析链接中..."
      });
      // 解析歌单链接
      this.$axios({
        method: "get",
        url: listnum
      })
        .then(function(res) {
          // 状态正常
          if (res.status == "200" && res.data.code == "200") {
            _this.getMore(res);
          } else {
            console.log("a");
            _this.$message({
              showClose: true,
              message: "链接错误或者网络出错",
              type: "warning"
            });
          }
        })
        .catch(function(res) {
          console.log(res);
          _this.$message({
            showClose: true,
            message: "链接错误或者网络出错",
            type: "warning"
          });
        });
    },
    clear(mode) {
      console.log("test");
      switch (mode) {
        case "0":
          this.url = "";
          break;
        case "1":
          this.number = "";
          break;
      }
    },
    Numbersure() {
      this.$message({
        showClose: true,
        message: "解析房号中..."
      });
      let _this = this;
      console.log("解析房号 ", this.number);
      this.$axios({
        method: "get",
        url: "http://192.168.1.105:3000/sharemusic/number?id=" + this.number
      })
        .then(function(res) {
          // console.log(res);
          if (res.status == "200" && res.data.code == "200") {
            _this.getMore(res);
            _this.getMore(res);
            _this.$store.commit("setNumber", res.data.number);
            _this.$store.commit("setMstate", res.data.state);
            _this.$store.commit("setNamelist", res.data.namelist);
            _this.$store.commit("setNowmusic", res.data.nownum);
          } else {
            _this.$message({
              showClose: true,
              message: "链接错误或者网络出错",
              type: "warning"
            });
          }
        })
        .catch(function(res) {
          console.log(res);
          _this.$message({
            showClose: true,
            message: "链接错误或者网络出错",
            type: "warning"
          });
        });
    }
  }
};
</script>

<style scoped>
.Table {
  margin: 0 auto;
}
</style>