<template>
  <apexchart type="radialBar" :options="options" :series="series"></apexchart>

<!--  <div class="container">-->
<!--    <h1>{{message}}</h1>-->
<!--&lt;!&ndash;    <a @click="getMessage">Press Me!</a>&ndash;&gt;-->
<!--  </div>-->
</template>

<script>
export default {
  data() {
    return {
      series: [0],
      options: {
        labels: ['CPU Usage']
      }
      // message: " "
    };
  },
  mounted: function () {
    window.wails.Events.On('cpu_usage', cpu_usage => {
      //this.message = cpu_usage.avg;
      if(cpu_usage.avg){
        // only one value
        this.series = [cpu_usage.avg];
      }

    });
  }

  /*
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.Stats.GetCPUUsage().then(cpu_usage => {
        self.message = cpu_usage.avg;
      });
    }
  }

   */
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
