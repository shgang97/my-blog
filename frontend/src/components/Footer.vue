<template>
  <div class="foot-container">
    <div>
      <span id="timeDate">本站已安全运行&nbsp{{dateTime.days}}&nbsp天&nbsp{{dateTime.hours}}&nbsp小时&nbsp{{dateTime.minutes}}&nbsp分&nbsp{{dateTime.seconds}}&nbsp秒</span>
    </div>
    <div>
      <a href="https://beian.miit.gov.cn/" target="_blank">
        <el-link :underline="true">苏ICP备2022042734号-1</el-link>
      </a>
    </div>
    <div>
      <a href="https://www.beian.gov.cn/portal/registerSystemInfo?recordcode=32011502011719" target="_blank">
        <img src="../assets/images/gongan.png" width="20" style="position: relative;top: 4px;">
        <el-link :underline="true">苏公网安备 32011502011719号</el-link>
      </a>
    </div>
    <div>
      Copyright © 2022 shgang
    </div>
  </div>
</template>

<script>
export default {
  name: 'Footer',
  data() {
    return {
      dateTime: {
        days: '',
        hours: '',
        minutes: '',
        seconds: ''
      }
    };
  },
  methods: {
    runTime() {
      const start = new Date('2022/05/06 00:00:00');
      let now = new Date();
      const _this = this;
      now.setTime(now.getTime() + 250);
      let leftTime = now - start;
      let toDay = 1000 * 60 * 60 * 24;
      _this.dateTime.days = Math.floor(leftTime / toDay);
      leftTime %= toDay;
      let toHour = toDay / 24;
      _this.dateTime.hours = Math.floor(leftTime / toHour);
      leftTime %= toHour;
      let toMinute = toHour / 60;
      _this.dateTime.minutes = Math.floor(leftTime / toMinute);
      leftTime %= toMinute;
      let toSecond = toMinute / 60;
      _this.dateTime.seconds = Math.floor(leftTime / toSecond);
    }
  },
  mounted() {
    let _this = this;
    this.timer = setInterval(() => {
      _this.runTime();
    }, 1000);
  },
  beforeDestroy() {
    clearInterval(this.timer);
  }
};
</script>

<style scoped>
/*.run-time-container {*/
/*  display: flex;*/
/*  justify-content: center;*/
/*  padding: 5px;*/
/*  !*background-color: #45a1ff;*!*/
/*  !*color: white;*!*/
/*}*/
.foot-container {
  display: flex;
  justify-content: center;
  padding: 10px;
  color: white;
  width: 100%;
  height: 40px;
  background-color: #292c2f;
  position: absolute;
  bottom: 0;
  font-size: 16px;
}

a {
  color: white;
  /*color: inherit;//继承父级元素的颜色*/
  /*text-decoration: none;//去除下划线*/
}
</style>