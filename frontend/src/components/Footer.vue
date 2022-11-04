<template>
  <footer>
    <div >
      <span id="timeDate">本站已安全运行&nbsp{{dateTime.days}}&nbsp天&nbsp{{dateTime.hours}}&nbsp小时&nbsp{{dateTime.minutes}}&nbsp分&nbsp{{dateTime.seconds}}&nbsp秒</span>
    </div>
    <div>
      <a href="https://beian.miit.gov.cn/" target="_blank">
        苏ICP备2022042734号-1</a>
      <span class="mx-2">|</span>
      <a href="https://www.beian.gov.cn/portal/registerSystemInfo?recordcode=32011502011719" target="_blank">
        <img src="/images/gongan.png" width="20" style="position: relative;top: 4px;">
        苏公网安备 32011502011719号
      </a>
      <span class="mx-2">|</span>
      Copyright © 2022 shgang
    </div>
  </footer>
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
div {
  display: flex;
  justify-content: center;
  padding: 5px;
  /*background-color: #45a1ff;*/
  color: #000000;
}

/*#timeDate {*/
/*  display: flex;*/
/*  justify-content: center;*/
/*  align-items: center;*/
/*}*/

.mx-2 {
  margin-left: 0.5rem !important;
}


</style>