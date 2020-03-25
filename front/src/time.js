module.exports = function(){
  var date = new Date();
  var year = date.getFullYear(); //获取当前年份4
  year = year.toString();
  var mon = date.getMonth() + 1; //获取当前月份2
  mon = mon.toString();
  while (mon.length < 2) {
    mon = "0"+mon;
  }
  var da = date.getDate(); //获取当前日2
  da = da.toString();
  while (da.length < 2) {
    da = "0"+da;
  }
  var h = date.getHours(); //获取小时2
  h = h.toString();
  while (h.length < 2) {
    h = "0"+h;
  }
  var m = date.getMinutes(); //获取分钟2
  m = m.toString();
  while (m.length < 2) {
    m = "0"+m;
  }
  var s = date.getSeconds(); //获取秒2
  s = s.toString();
  while (s.length < 2) {
    s = "0"+s;
  }
  var ms = date.getMilliseconds();
  ms = ms.toString();
  while (ms.length < 3) {
    ms = "0"+ms;
  }
  return year+mon+da+ h+ m+ s+ ms;
}