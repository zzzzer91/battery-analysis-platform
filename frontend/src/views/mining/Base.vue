<template>
  <div>
    <div class="container">
      <el-row>
        <el-col :span="8">
          <el-form ref="form" :model="queryForm" label-width="80px">
            <el-form-item label="数据来源">
              <el-cascader
                :props="{ expandTrigger: 'hover' }"
                :options="formOptions.dataComeFrom"
                v-model="queryForm.dataComeFrom"
                @change="dataComeFromChange"
              ></el-cascader>
            </el-form-item>
            <el-form-item label="查询参数">
              <el-select v-model="queryForm.needParams" multiple collapse-tags placeholder="请选择">
                <el-option-group
                  v-for="group in formOptions.needParams[queryForm.dataComeFrom[0]]"
                  :key="group.label"
                  :label="group.label"
                >
                  <el-option
                    v-for="item in group.options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-option-group>
              </el-select>
            </el-form-item>
            <el-form-item label="起始日期">
              <el-date-picker v-model="queryForm.startDate" type="datetime" placeholder="选择起始日期时间"></el-date-picker>
            </el-form-item>
            <el-form-item label="数据限制">
              <el-input-number :min="10" :max="10000" :step="10" v-model="queryForm.dataLimit"></el-input-number>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="getChartData" :loading="buttonLoading">查询</el-button>
              <el-button type="primary" @click="showChartOptionDialog" v-if="showPlotButton">绘制</el-button>
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :span="16">
          <el-input
            type="textarea"
            readonly
            resize="none"
            rows="21"
            wrap="off"
            v-model="dataTextArea"
          ></el-input>
        </el-col>
      </el-row>
    </div>
    <!-- 创建任务弹出框 -->
    <el-dialog
      title="绘制选项"
      :visible.sync="chartOptionDialogVisible"
      :close-on-click-modal="false"
      width="50%"
    >
      <el-row>
        <el-col :span="15">
          <el-tabs v-model="plotOption.plotTabActiveName" tab-position="left">
            <el-tab-pane label="折线图" name="plotLineTab">
              <!-- <div class="mgb20">
                <span>X 轴：</span>
                <el-select v-model="plotOption.xAxisParam" placeholder="请选择">
                  <el-option
                    v-for="item in plotOption.xAxisParamOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </div>-->
              <!-- <div class="mgb20">
                  <span>X 轴类型：</span>
                  <el-radio v-model="plotOption.xAxisType" label="category">类目</el-radio>
                  <el-radio v-model="plotOption.xAxisType" label="value">数值</el-radio>
              </div>-->
              <div class="mgb20">
                <span>Y轴：</span>
                <el-select
                  v-model="plotOption.yAxisParams"
                  :multiple-limit="plotOption.yAxisDataLimit"
                  multiple
                  collapse-tags
                  placeholder="请选择"
                >
                  <el-option
                    v-for="item in plotOption.yAxisParamOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </div>
              <div class="mgb20">
                <span>线段：</span>
                <el-radio-group v-model="plotOption.lineType" size="small">
                  <el-radio-button label="直线"></el-radio-button>
                  <el-radio-button label="阶梯线"></el-radio-button>
                  <el-radio-button label="曲线"></el-radio-button>
                </el-radio-group>
              </div>
              <div class="mgb20">
                <span>选项：</span>
                <el-checkbox v-model="plotOption.doubleYAxis" @change="changeYAxisDataLimit">双Y轴</el-checkbox>
                <el-checkbox v-model="plotOption.yAxisAutoChange">Y轴自动调整</el-checkbox>
              </div>
            </el-tab-pane>
            <el-tab-pane label="3D散点图" name="plot3dScatterTab">
              <div class="mgb20">
                <span>X轴：</span>
                <el-select v-model="plotOption.xAxis3dParam" placeholder="请选择">
                  <el-option
                    v-for="item in plotOption.yAxisParamOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </div>
              <div class="mgb20">
                <span>Y轴：</span>
                <el-select v-model="plotOption.yAxis3dParam" placeholder="请选择">
                  <el-option
                    v-for="item in plotOption.yAxisParamOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                  ></el-option>
                </el-select>
              </div>
              <div class="mgb20">
                <span>选项：</span>
                <el-checkbox v-model="plotOption.xAxis3dAutoChange">X轴自动调整</el-checkbox>
                <el-checkbox v-model="plotOption.yAxis3dAutoChange">Y轴自动调整</el-checkbox>
              </div>
            </el-tab-pane>
          </el-tabs>
        </el-col>
        <el-col :span="9">
          <div class="mgb20">
            <span>数据开始索引：</span>
            <el-input-number
              v-model="plotOption.dataIndexStart"
              controls-position="right"
              :min="0"
              :max="chartDataCount"
              size="small"
            ></el-input-number>
          </div>
          <div class="mgb20">
            <span>数据结束索引：</span>
            <el-input-number
              v-model="plotOption.dataIndexEnd"
              controls-position="right"
              :min="0"
              :max="chartDataCount"
              size="small"
            ></el-input-number>
          </div>
          <div class="mgb20" style="margin-left: 120px">
            <el-checkbox v-model="plotOption.simplePlot">简易模式</el-checkbox>
          </div>
        </el-col>
      </el-row>
      <span slot="footer">
        <el-button @click="chartOptionDialogVisible=false">取 消</el-button>
        <el-button type="primary" @click="doPlot">确 定</el-button>
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
  name: 'MiningBase',
  data() {
    return {
      chartOptionDialogVisible: false,
      chartDialogVisible: false,
      showPlotButton: false,
      formOptions: {
        dataComeFrom: [
          {
            value: '宇通',
            label: '宇通',
            children: [
              {
                value: '4F37195C1A908CFBE0532932A8C0EECB',
                label: '4F37195C1A908CFBE0532932A8C0EECB'
              }
            ]
          },
          {
            value: '北汽',
            label: '北汽',
            children: [
              {
                value: 'LNBSCU3HXJR884327',
                label: 'LNBSCU3HXJR884327'
              }
            ]
          },
        ],
        needParams: {
          宇通: [
            // label 长度不要超过 7 个汉字，否则有样式问题
            {
              label: '总体',
              options: [
                {
                  value: '总电压',
                  label: '总电压'
                },
                {
                  value: '总电流',
                  label: '总电流'
                },
                {
                  value: '车速',
                  label: '车速'
                },
                {
                  value: '正向累计电量',
                  label: '正向累计电量'
                },
                {
                  value: '反向累计电量',
                  label: '反向累计电量'
                },
                {
                  value: '总里程',
                  label: '总里程'
                },
                {
                  value: 'SOC',
                  label: 'SOC'
                },
                {
                  value: '状态号',
                  label: '状态号'
                }
              ]
            },
            {
              label: '单体',
              options: [
                {
                  value: '单体最高温度',
                  label: '单体最高温度'
                },
                {
                  value: '单体最低温度',
                  label: '单体最低温度'
                },
                {
                  value: '单体最高电压',
                  label: '单体最高电压'
                },
                {
                  value: '单体最低电压',
                  label: '单体最低电压'
                },
                {
                  value: '最高温度电池号',
                  label: '最高温度电池号'
                },
                {
                  value: '最低温度电池号',
                  label: '最低温度电池号'
                },
                {
                  value: '最高电压电池号',
                  label: '最高电压电池号'
                },
                {
                  value: '最低电压电池号',
                  label: '最低电压电池号'
                }
              ]
            }
          ],
          北汽: [
            {
              options: [
                {
                  value: '动力电池内部总电压V1',
                  label: '动力电池内部总电压V1'
                },
                {
                  value: '动力电池充/放电电流',
                  label: '总动力电池充/放电电流'
                },
                {
                  value: '动力电池可用能量',
                  label: '动力电池可用容量'
                },
                {
                  value: '动力电池剩余电量SOC',
                  label: '动力电池剩余电量SOC'
                },
                {
                  value: '动力电池充放电状态',
                  label: '动力电池充放电状态'
                },
                {
                  value: 'MSODO总里程',
                  label: 'MSODO总里程'
                }
              ]
            }
          ]
        }
      },
      // 请求的参数
      queryForm: {
        dataComeFrom: ['宇通', '4F37195C1A908CFBE0532932A8C0EECB'],
        needParams: [],
        startDate: new Date(2018, 11, 1, 0, 0), // 貌似 js 中月份起始是 0？
        dataLimit: 500
      },
      buttonLoading: false, // 查询按钮的载入效果
      dataTextArea: '',
      // 绘制选项
      plotOption: {
        xAxisParamOptions: [],
        yAxisParamOptions: [],
        dataIndexStart: 0, // 数据起索引
        dataIndexEnd: 0, // 数据结束索引
        simplePlot: false, // 简洁绘图模式，不支持缩放
        doubleYAxis: false, // 双 Y 轴
        yAxisAutoChange: false, // Y 轴根据数据上下限自动调整
        yAxisDataLimit: 0, // 如果设置了双 Y 轴，则限制只能选取两个数据
        lineType: '直线', // 线段类型，有直线，阶梯线，曲线
        xAxisType: 'category',
        // 绘图前需重新设置
        plotTabActiveName: 'plotLineTab',
        xAxisParam: '时间', // 单选，值是个字符串
        yAxisParams: [], // 多选，值是个列表
        xAxis3dParam: null, // 3d
        yAxis3dParam: null,
        zAxis3dParam: '时间',
        xAxis3dAutoChange: false, // X 轴根据数据上下限自动调整
        yAxis3dAutoChange: false, // Y 轴根据数据上下限自动调整
      },
      // 图表数据
      chartData: [],
      chartDataCount: 0, // 数据长度
      chartOption: {}
    }
  },
  methods: {
    // 更换数据来源后，清空选择的参数项
    dataComeFromChange() {
      this.queryForm.needParams = []
    },
    getChartData() {
      if (this.queryForm.needParams.length === 0) {
        this.$message.error('查询参数不能为空！')
        return false
      }

      if (this.queryForm.startDate === null) {
        this.$message.error('起始日期不能为空！')
        return false
      }

      this.buttonLoading = true

      let params = {
        dataComeFrom: this.queryForm.dataComeFrom.join('_'),
        startDate: moment(this.queryForm.startDate).format(
          'YYYY-MM-DD HH:mm:ss'
        ),
        dataLimit: this.queryForm.dataLimit,
        needParams: this.queryForm.needParams.join(',')
      }

      return (
        this.$axios
          .get(globals.URL_API_MINING_BASE, {
            params: params
          })
          // response 有多种属性
          .then(response => response.data)
          .then(jd => {
            if (jd.code !== globals.SUCCESS) {
              throw new Error(jd.msg)
            }

            const colNames = ['时间'].concat(this.queryForm.needParams)
            this.chartData = jd.data
            this.chartDataCount = this.chartData.length

            this.dataTextArea = this.buildDataTextArea(colNames, this.chartData)

            let xAxisParamOptions = []
            let yAxisParamOptions = []
            for (let name of colNames) {
              // 填充 x 轴数据选项
              xAxisParamOptions.push({
                label: name,
                value: name
              })
              if (name !== '时间') {
                // 填充 y 轴数据选项
                yAxisParamOptions.push({
                  label: name,
                  value: name
                })
              }
            }

            this.plotOption.xAxisParamOptions = xAxisParamOptions
            this.plotOption.yAxisParamOptions = yAxisParamOptions
            this.plotOption.dataIndexEnd = this.chartDataCount

            this.buttonLoading = false
            this.showPlotButton = true

            this.$message.success(jd.msg)
          })
          .catch(error => {
            this.buttonLoading = false
            this.$message.error(error.message)
          })
      )
    },
    // 字符串化返回的数据
    buildDataTextArea(colNames, data) {
      let temp = []
      for (let colName of colNames) {
        temp.push(colName)
      }
      let s = [
        `                  ${temp[0]}             ${temp.slice(1).join('    ')}`
      ]
      let i = 0
      for (let row of data) {
        let temp = `${globals.leftFillSpace(i.toString(), 6)}    `
        for (let colName of colNames) {
          let col = row[colName]
          if (col !== null) {
            // 注意 0 也会被判 false
            temp += `${globals.leftFillSpace(col.toString(), 10)}`
          } else {
            temp += `${globals.leftFillSpace('null', 10)}`
          }
        }
        s.push(temp)
        i++
      }
      return s.join('\n')
    },
    showChartOptionDialog() {
      this.plotOption.plotTabActiveName = 'plotLineTab'
      this.plotOption.dataIndexStart = 0
      this.plotOption.xAxisParam = '时间'
      this.plotOption.yAxisParams = []
      this.plotOption.xAxis3dParam = null
      this.plotOption.yAxis3dParam = null
      this.plotOption.zAxis3dParam = '时间'
      this.chartOptionDialogVisible = true
    },
    changeYAxisDataLimit() {
      this.plotOption.yAxisDataLimit =
        this.plotOption.yAxisDataLimit === 0 ? 2 : 0
      if (this.plotOption.yAxisParams.length > 2) {
        this.plotOption.yAxisParams = []
      }
    },
    beforeChartDialogClose(done) {
      // Dialog 设置了 destroy-on-close="true"，
      // 会在关闭时销毁其中的元素，但这在与 echarts 配合时，会有一些问题，
      // 销毁元素后，echarts 会侦测到 this.chartOption 有数据，它会在销毁元素后马上进行绘制，
      // 这里防止这个问题
      this.chartOption = {}
      done()
    },
    _checkLinePlotOption(plotOption) {
      if (plotOption.dataIndexStart >= plotOption.dataIndexEnd) {
        return '请重新指定数据起止范围！'
      }

      if (plotOption.yAxisParams.length === 0) {
        return 'Y 轴不能为空！'
      }
      if (plotOption.yAxisParams.indexOf(plotOption.xAxisParam) !== -1) {
        return '字段不能同时指定为 X 轴和 Y 轴！'
      }
      if (
        plotOption.dataIndexEnd - plotOption.dataIndexStart > 2000 &&
        !plotOption.simplePlot
      ) {
        return '绘制折线图时，数据量不能大于 2000，否则会遭遇性能问题，请开启简易模式。'
      }

      return null
    },
    _check3dScatterPlotOption(plotOption) {
      if (plotOption.dataIndexStart >= plotOption.dataIndexEnd) {
        return '请重新指定数据起止范围！'
      }

      if (plotOption.xAxis3dParam === null) {
        return 'X 轴不能为空！'
      }
      if (plotOption.yAxis3dParam === null) {
        return 'Y 轴不能为空！'
      }
      if (plotOption.zAxis3dParam === null) {
        return 'Z 轴不能为空！'
      }
      return null
    },
    _buildLineChartOption(plotOption, chartData) {
      let chartOption = {
        toolbox: {
          feature: {
            saveAsImage: {}
          }
        },
        dataset: {
          dimensions: [plotOption.xAxisParam].concat(this.queryForm.needParams),
          source: chartData.slice(
            plotOption.dataIndexStart,
            plotOption.dataIndexEnd
          )
        },
        xAxis: {
          show: true,
          type: plotOption.xAxisType
        },
        legend: {},
        series: []
      }

      // 线段类型
      let step = false
      let smooth = false
      if (plotOption.lineType === '阶梯线') {
        step = true
      } else if (plotOption.lineType === '曲线') {
        smooth = true
      }

      // 创建 series 对象
      for (let k of plotOption.yAxisParams) {
        let name = k
        // 创建 series 对象
        chartOption.series.push({
          smooth,
          step,
          name,
          type: 'line'
        })
        // dataset 情况下不需要设置 legend
        //chartOption.legend.data.push(name)
      }

      // 是否根据数据上下限调整 Y 轴
      let min = null
      let max = null
      if (plotOption.yAxisAutoChange) {
        min = 'dataMin'
        max = 'dataMax'
      }

      // 双 Y 轴
      if (plotOption.doubleYAxis && chartOption.series.length === 2) {
        chartOption.yAxis = [
          {
            name: chartOption.series[0].name,
            min,
            max
          },
          {
            name: chartOption.series[1].name,
            min,
            max
          }
        ]
        chartOption.series[1].yAxisIndex = 1
      } else {
        chartOption.yAxis = {
          min,
          max
        }
      }

      // 简易模式，砍掉动态缩放等
      if (!plotOption.simplePlot) {
        chartOption.toolbox.feature.restore = {}

        chartOption.dataZoom = [
          {
            show: true,
            realtime: true
          },
          {
            type: 'inside',
            realtime: true
          }
        ]

        chartOption.tooltip = {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            label: {
              backgroundColor: '#6a7985'
            }
          }
        }
      }

      // vue 不会监控某个对象的成员，要改变整个对象才能响应式
      return chartOption
    },
    _build3dScatterChartOption(plotOption, chartData) {
      const xName = plotOption.xAxis3dParam
      const yName = plotOption.yAxis3dParam
      const zName = plotOption.zAxis3dParam
      let data = []
      for (
        let i = plotOption.dataIndexStart;
        i < plotOption.dataIndexEnd;
        i++
      ) {
        data.push([chartData[i][xName], chartData[i][yName], i])
      }

      // 是否根据数据上下限调整 X 轴 和 Y 轴
      let xMin = null
      let xMax = null
      if (plotOption.xAxis3dAutoChange) {
        xMin = 'dataMin'
        xMax = 'dataMax'
      }
      let yMin = null
      let yMax = null
      if (plotOption.yAxis3dAutoChange) {
        yMin = 'dataMin'
        yMax = 'dataMax'
      }

      return {
        grid3D: {},
        xAxis3D: {
          type: 'value',
          name: xName,
          min: xMin,
          max: xMax,
        },
        yAxis3D: {
          type: 'value',
          name: yName,
          min: yMin,
          max: yMax,
        },
        zAxis3D: {
          type: 'value',
          name: zName
        },
        series: [
          {
            type: 'scatter3D',
            symbolSize: 2.5,
            data: data
          }
        ]
      }
    },
    doPlot() {
      let chartOption = null
      if (this.plotOption.plotTabActiveName === 'plotLineTab') {
        const ret = this._checkLinePlotOption(this.plotOption)
        if (ret !== null) {
          this.$message.error(ret)
          return
        }
        chartOption = this._buildLineChartOption(
          this.plotOption,
          this.chartData
        )
      } else if (this.plotOption.plotTabActiveName === 'plot3dScatterTab') {
        const ret = this._check3dScatterPlotOption(this.plotOption)
        if (ret !== null) {
          this.$message.error(ret)
          return
        }
        chartOption = this._build3dScatterChartOption(
          this.plotOption,
          this.chartData
        )
      }
      this.chartDialogVisible = true
      this.chartOption = chartOption
    }
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
  height: 500px;
}
</style>