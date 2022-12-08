<template>
  <div class="foot-container">
    <div>
      <span
          id="timeDate">本站已安全运行&nbsp{{ dateTime.days }}&nbsp天&nbsp{{ dateTime.hours }}&nbsp小时&nbsp{{ dateTime.minutes }}&nbsp分&nbsp{{ dateTime.seconds }}&nbsp秒</span>
    </div>
    <br/>
    <div>
      <a href="https://beian.miit.gov.cn/" target="_blank">
        &nbsp;&nbsp;苏ICP备2022042734号-1&nbsp;&nbsp;
      </a>
    </div>
    <div>
      <a href="https://www.beian.gov.cn/portal/registerSystemInfo?recordcode=32011502011719" target="_blank">
        <img src="../assets/images/gongan.png" width="20" style="position: relative;top: 4px;">
        苏公网安备 32011502011719号
      </a>
    </div>
    <div>
      &nbsp;&nbsp;Copyright © 2022 shgang
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
      now.setTime(now.getTime() + 250);
      let leftTime = now - start;
      let toDay = 1000 * 60 * 60 * 24;
      this.dateTime.days = Math.floor(leftTime / toDay);
      leftTime %= toDay;
      let toHour = toDay / 24;
      this.dateTime.hours = Math.floor(leftTime / toHour);
      leftTime %= toHour;
      let toMinute = toHour / 60;
      this.dateTime.minutes = Math.floor(leftTime / toMinute);
      leftTime %= toMinute;
      let toSecond = toMinute / 60;
      this.dateTime.seconds = Math.floor(leftTime / toSecond);
    }
  },
  mounted() {
    this.timer = setInterval(() => {
      this.runTime();
    }, 1000);
  },
  beforeDestroy() {
    clearInterval(this.timer);
  }
};
</script>

<style scoped>
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
  line-height: 35px;
}

a {
  color: inherit;
  text-decoration: none;
}
</style>