<template>
  <div class="container">
    <v-chart :options="chartOption" />
  </div>
</template>

<script>
import globals from '@/globals'

export default {
  name: 'MiningWorkingCondition',
  data() {
    return {
      chartOption: null,
      // 注意 index 要排第一，作为 X 轴
      // 也可作为数据顺序
      needParams: ['time', 'speed', 'cur'],
      // 返回数据的 key，转换成中文名
      mapping: {
        time: '时间',
        speed: '速度',
        cur: '电流'
      }
    }
  },
  methods: {
    getChartData() {
      this.$axios
        .get(globals.URL_API_MINING + '/working-condition')
        // response 有多种属性
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }

          let data = jd.data

          let chartOption = {
            tooltip: {
              trigger: 'axis',
              axisPointer: {
                type: 'cross',
                label: {
                  backgroundColor: '#6a7985'
                }
              }
            },
            toolbox: {
              feature: {
                // dataView: {show: true, readOnly: true},
                restore: {},
                saveAsImage: {}
              }
            },
            dataZoom: [
              {
                show: true,
                realtime: true
              },
              {
                type: 'inside',
                realtime: true
              }
            ],
            xAxis: {
              name: '时间',
              type: 'category'
            },
            yAxis: {},
            dataset: {
              dimensions: this.needParams,
              source: data
            },
            series: [],
            legend: {
              data: []
            }
          }

          for (let k of this.needParams.slice(1)) {
            let name = this.mapping[k]
            chartOption.series.push({
              name: name,
              type: 'line'
            })
            chartOption.legend.data.push(name)
          }

          // vue 不会监控某个对象的成员，要改变整个对象才能响应式
          this.chartOption = chartOption
        })
        .catch(error => {
          this.$message.error(error.message)
        })
    }
  },
  created() {
    this.getChartData()
  }
}
</script>

<style scoped>
/**
 * 默认尺寸为 600px×400px，如果想让图表响应尺寸变化，可以像下面这样
 * 把尺寸设为百分比值（同时请记得为容器设置尺寸）。
 */
.echarts {
  width: 100%;
  height: 450px;
}
</style>