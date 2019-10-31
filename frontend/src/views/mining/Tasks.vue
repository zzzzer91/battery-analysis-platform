<template>
  <div>
    <div class="container">
      <div class="mgb20">
        <el-button type="primary" icon="el-icon-plus" @click="handleNewForm">新建任务</el-button>
      </div>
      <el-table :data="tableData" border height="500" ref="multipleTable">
        <el-table-column prop="taskName" label="任务名" width="100" sortable show-overflow-tooltip></el-table-column>
        <el-table-column prop="dataComeFrom" label="数据来源" width="340" sortable></el-table-column>
        <!-- <el-table-column prop="batteryStatus" label="电池状态" width="100" show-overflow-tooltip></el-table-column> -->
        <el-table-column prop="dateRange" label="时间范围" width="180" show-overflow-tooltip></el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180"></el-table-column>
        <el-table-column prop="taskStatus" label="状态" width="100" sortable>
          <template scope="scope">
            <el-tag v-if="scope.row.taskStatus === 0" type="warning">创建中</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === 1" type="warning">执行中</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === 6" type="success">完成</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === 7" type="danger">失败</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="comment" label="备注" show-overflow-tooltip></el-table-column>
        <el-table-column label="操作" width="160">
          <template v-slot:default="scope">
            <el-button
              type="text"
              icon="el-icon-pie-chart"
              @click="doPlot(scope.$index, scope.row)"
              v-show="scope.row.taskStatus===6"
            >绘制</el-button>
            <el-button
              type="text"
              icon="el-icon-delete"
              @click="deleteTask(scope.$index, scope.row)"
              class="red"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 创建任务弹出框 -->
    <el-dialog
      title="创建计算任务"
      :visible.sync="newTaskDialogVisible"
      :close-on-click-modal="false"
      width="40%"
    >
      <el-form ref="form" :model="newForm" label-width="80px">
        <el-form-item label="数据来源">
          <el-cascader
            props.expand-trigger="hover"
            :options="formOptions.dataComeFrom"
            v-model="newForm.dataComeFrom"
          ></el-cascader>
        </el-form-item>
        <el-form-item label="计算模型">
          <el-select v-model="newForm.taskName" placeholder="请选择">
            <el-option
              v-for="item in formOptions.taskName"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="电池状态：">
          <el-radio-group v-model="newForm.batteryStatus">
            <el-radio :label="0">全部</el-radio>
            <el-radio :label="1">充电</el-radio>
            <el-radio :label="2">放电</el-radio>
          </el-radio-group>
        </el-form-item>-->
        <el-form-item label="起止日期">
          <el-date-picker
            v-model="newForm.dateRange"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            :disabled="newForm.allData"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="所有数据">
          <el-checkbox v-model="newForm.allData"></el-checkbox>
        </el-form-item>
      </el-form>
      <span slot="footer">
        <el-button @click="newTaskDialogVisible=false">取 消</el-button>
        <el-button type="primary" @click="createTask">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 图表弹出框 -->
    <el-dialog
      :visible.sync="chartDialogVisible"
      :close-on-click-modal="false"
      :destroy-on-close="true"
      width="80%"
      :before-close="beforeChartDialogClose"
    >
      <v-chart :options="chartOption" />
    </el-dialog>
  </div>
</template>

<script>
import moment from 'moment'
import globals from '@/globals'

export default {
  name: 'MiningTasks',
  data() {
    return {
      ws: null, // websocket 对象
      wsTimer: null,
      newTaskDialogVisible: false,
      chartDialogVisible: false,
      formOptions: {
        dataComeFrom: [
          // {
          //   value: '北汽',
          //   label: '北汽',
          //   children: [
          //     {
          //       value: '占位',
          //       label: '占位'
          //     }
          //   ]
          // },
          {
            value: '宇通',
            label: '宇通',
            children: [
              {
                value: '4F37195C1A908CFBE0532932A8C0EECB',
                label: '4F37195C1A908CFBE0532932A8C0EECB'
              }
            ]
          }
        ],
        // label 长度不要超过 7 个汉字，否则有样式问题
        taskName: [
          {
            value: '充电过程',
            label: '充电过程'
          },
          // {
          //   value: '工况',
          //   label: '工况'
          // },
          {
            value: '电池统计',
            label: '电池统计'
          },
          {
            value: 'pearson相关系数',
            label: 'pearson相关系数'
          }
        ]
      },
      // 请求的参数
      newForm: {
        dataComeFrom: ['宇通', '4F37195C1A908CFBE0532932A8C0EECB'],
        taskName: null,
        batteryStatus: 0,
        dateRange: [new Date(2019, 0, 1, 0, 0), new Date(2019, 0, 2, 0, 0)],
        allData: false
      },
      tableData: [],
      chartOption: {}
    }
  },
  methods: {
    handleNewForm() {
      this.newForm = {
        dataComeFrom: ['宇通', '4F37195C1A908CFBE0532932A8C0EECB'],
        taskName: null,
        batteryStatus: 0,
        dateRange: [new Date(2019, 0, 1, 0, 0), new Date(2019, 0, 2, 0, 0)],
        allData: false
      }
      this.newTaskDialogVisible = true
    },
    createTask() {
      if (this.newForm.taskName === null) {
        this.$message.error('计算模型不能为空！')
        return false
      }
      // 时间范围清空数据，会返回 null
      if (this.newForm.dateRange === null && !this.newForm.allData) {
        this.$message.error('时间范围不能为空！或请勾选所有数据')
        return false
      }

      const startDate = moment(this.newForm.dateRange[0]).format(
        'YYYY-MM-DD HH:mm:ss'
      )
      const endDate = moment(this.newForm.dateRange[1]).format(
        'YYYY-MM-DD HH:mm:ss'
      )
      let params = {
        dataComeFrom: this.newForm.dataComeFrom.join('_'),
        taskName: this.newForm.taskName,
        batteryStatus: this.newForm.batteryStatus,
        startDate: startDate,
        endDate: endDate,
        allData: this.newForm.allData
      }

      return (
        this.$axios
          .post(globals.URL_API_MINING_TASKS, params)
          // response 有多种属性
          .then(response => response.data)
          .then(jd => {
            if (jd.code !== globals.SUCCESS) {
              throw new Error(jd.msg)
            }

            this.tableData.unshift(jd.data)

            this.newTaskDialogVisible = false
            this.$message.success(jd.msg)
          })
          .catch(error => {
            this.$message.error(error.message)
          })
      )
    },
    beforeChartDialogClose(done) {
      // Dialog 设置了 destroy-on-close="true"，
      // 会在关闭时销毁其中的元素，但这在与 echarts 配合时，会有一些问题，
      // 销毁元素后，echarts 会侦测到 this.chartOption 有数据，它会在销毁元素后马上进行绘制，
      // 这里防止这个问题
      this.chartOption = {}
      done()
    },
    _transformJsonData(data) {
      // 将返回的json数据转换成 echarts 可接收的类型
      // 进行转换，而不用 dataset 的原因是，dataset 在查看数据视图时有 bug
      let obj = {}
      const keys = Object.keys(data[0])
      for (let k of keys) {
        obj[k] = []
      }

      for (let row of data) {
        for (let k of keys) {
          obj[k].push(row[k])
        }
      }

      return obj
    },
    _buildLineOrBar(taskName, data) {
      let mapping = null // 变量名到中文名的映射
      let dimensions = null // 用于保持数据顺序
      let xAxisName = null // X 轴名
      let boundaryGap = false // 坐标轴两边留不留白
      let seriesType = null // 图表类型
      let magicType = null // 切换图表样式时用，支持折线和柱状图
      if (taskName === '充电过程') {
        mapping = {
          init_soc: '初始SOC',
          last_soc: '终止SOC',
          last_vol: '终止电压',
          max_vol: '最大电压',
          sub_vol: '压差'
        }
        dimensions = [
          // 不要改，用于保持数据顺序
          'index',
          'init_soc',
          'last_soc',
          'last_vol',
          'sub_vol',
          'max_vol'
        ]
        xAxisName = '充电次数'
        seriesType = 'line'
        magicType = ['line', 'bar']
      } else if (taskName === '电池统计') {
        mapping = {
          max_t_count: '最大温度次数',
          min_t_count: '最小温度次数'
        }
        dimensions = ['number', 'max_t_count', 'min_t_count']
        xAxisName = '电池号'
        boundaryGap = true // 柱状图要设为 true，来留白，不然有显示问题
        seriesType = 'bar'
        magicType = ['stack', 'tiled']
      } else if (taskName === '工况') {
        mapping = {
          speed: '速度',
          cur: '电流'
        }
        dimensions = ['time', 'speed', 'cur']
        xAxisName = '时间'
        seriesType = 'line'
        magicType = ['line', 'bar']
      }

      let transformedData = this._transformJsonData(data)

      // 构建 X 轴
      let xAxis = {
        name: xAxisName,
        type: 'category',
        boundaryGap: boundaryGap,
        data: transformedData[dimensions[0]]
      }

      // 构建 series
      let series = []
      let legendData = []
      for (let k of dimensions.slice(1)) {
        let name = mapping[k]
        series.push({
          name: name,
          type: seriesType,
          data: transformedData[k]
        })
        legendData.push(name)
      }

      return {
        title: {
          text: taskName,
          left: 'top'
        },
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
            // 使用 dataset 时 dataView 会有显示问题，貌似是官方 bug
            dataView: { show: true, readOnly: true },
            // 可以切换图表类型
            magicType: { type: magicType },
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
        xAxis: xAxis,
        yAxis: {},
        // dataset: {
        //   dimensions: dimensions,
        //   source: data
        // },
        series: series,
        legend: {
          data: legendData
        }
      }
    },
    _buildHeatMap(taskName, data) {
      // 字段名不能太长，否则会显示不全
      const xy = [
        '总电压',
        '总电流',
        '车速',
        'SOC',
        '单体最高温度',
        '单体最低温度',
        '单体最高电压',
        '单体最低电压'
      ]
      return {
        title: {
          text: taskName,
          left: 'top'
        },
        tooltip: {
          position: 'top'
        },
        animation: false,
        grid: {
          height: '70%'
        },
        xAxis: {
          type: 'category',
          data: xy,
          splitArea: {
            show: true
          }
        },
        yAxis: {
          type: 'category',
          data: xy,
          splitArea: {
            show: true
          }
        },
        visualMap: {
          min: -1,
          max: 1,
          calculable: true,
          orient: 'horizontal',
          left: 'center',
          inRange: {
            color: [
              '#313695',
              '#4575b4',
              '#74add1',
              '#abd9e9',
              '#e0f3f8',
              '#ffffbf',
              '#fee090',
              '#fdae61',
              '#f46d43',
              '#d73027',
              '#a50026'
            ]
          }
        },
        series: [
          {
            type: 'heatmap',
            data: data,
            label: {
              normal: {
                show: true
              }
            },
            itemStyle: {
              emphasis: {
                shadowBlur: 10,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      }
    },
    doPlot(index, row) {
      this.$axios
        .get(`${globals.URL_API_MINING_TASKS}/${row.taskId}/data`)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }
          this.chartDialogVisible = true
          if (row.taskName === 'pearson相关系数') {
            this.chartOption = this._buildHeatMap(row.taskName, jd.data)
          } else {
            this.chartOption = this._buildLineOrBar(row.taskName, jd.data)
          }
        })
        .catch(error => {
          this.$message.error(error.message)
        })
    },
    deleteTask(index, row) {
      this.$confirm('确定删除？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.$axios
            .delete(`${globals.URL_API_MINING_TASKS}/${row.taskId}`)
            .then(response => response.data)
            .then(jd => {
              if (jd.code !== globals.SUCCESS) {
                throw new Error(jd.msg)
              }

              this.tableData.splice(index, 1)
              this.$message.success(jd.msg)
            })
            .catch(error => {
              this.$message.error(error.message)
            })
        })
        .catch(() => {})
    },
    getTaskList() {
      this.$axios
        .get(globals.URL_API_MINING_TASKS)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }
          return jd.data
        })
        .catch(error => {
          this.$message.error(error.message)
        })
    }
  },
  created() {
    this.ws = new WebSocket(
      'ws://' + document.location.host + globals.URL_WS_MINING_TASKS
    )

    // 收到服务器数据后的回调函数
    // 这里要用箭头函数，不然 this 的指向不对
    this.ws.onmessage = event => {
      let jd = JSON.parse(event.data)
      if (jd.code !== globals.SUCCESS) {
        return
      }
      this.tableData = jd.data
    }
  },
  beforeDestroy() {
    this.ws.close()
    this.ws = null
  }
}
</script>

<style scoped>
.echarts {
  width: 100%;
  height: 500px;
}
.red {
  color: #ff5722;
}
</style>