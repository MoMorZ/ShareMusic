import Vuex from 'vuex'
import vue from 'vue'
vue.use(Vuex)

export default new Vuex.Store({
  state: {
    playlist: [{id: 'dsafa'}],
    mstate: false, //是否解析
    playstate: 0, //是否在播放
    number: "",
    nowmusic: -1, 
    time: 0,
    namelist: [{name: 'test'}]
  },
  getters: {
    getPlaylist(state){
      return state.playlist;
    },
    getMstate(state){
      return state.mstate;
    },
    getNumber(state){
      return state.number;
    },
    getNowmusic(state){
      return state.nowmusic;
    }
  },
  mutations: {
    setPlaylist(state, playlist){
      state.playlist = playlist;
    },
    setMstate(state, newstate) {
      state.mstate = newstate;
    },
    setNumber(state, number){
      state.number = number;
    },
    setNowmusic(state, music){
      state.nowmusic = music;
    },
    setNamelist(state, namelist){
      state.namelist = namelist;
    },
    setPlaystate(state, playstate){
      state.playstate = playstate;
    },
    setTime(state, time){
      state.time = time;
    },
    setTIme2(state, add){
      state.time = state.time + add;
    }
  }
})