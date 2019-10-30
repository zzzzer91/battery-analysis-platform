<template>
  <div>
    <div class="container">
      <div class="mgb20">
        <el-button type="primary" icon="el-icon-plus" @click="handleNewForm">新建训练</el-button>
      </div>
      <el-table :data="tableData" border height="500" ref="multipleTable">
        <el-table-column prop="dataset" label="数据集" width="340" sortable></el-table-column>
        <el-table-column prop="hyperParameter" label="超参" width="100" show-overflow-tooltip>
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <textarea
                v-html="JSON.stringify(scope.row.hyperParameter,null,4)"
                readonly
                rows="16"
                cols="46"
              ></textarea>
              <div slot="reference">
                <el-tag>查看</el-tag>
              </div>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180"></el-table-column>
        <el-table-column prop="taskStatus" label="状态" width="100" sortable>
          <template scope="scope">
            <el-tag v-if="scope.row.taskStatus === 6" type="success">完成</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === 0" type="warning">创建中</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === 1" type="warning">执行中</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === 7" type="danger">失败</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="comment" label="备注" show-overflow-tooltip></el-table-column>
        <el-table-column label="操作" width="300">
          <template v-slot:default="scope">
            <el-button
              type="text"
              icon="el-icon-s-promotion"
              @click="plotTrainingHistory(scope.$index, scope.row)"
              v-show="scope.row.taskStatus===6"
            >训练过程</el-button>
            <el-button
              type="text"
              icon="el-icon-pie-chart"
              @click="plotEvalResult(scope.$index, scope.row)"
              v-show="scope.row.taskStatus===6"
            >评估结果</el-button>
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
      title="新建训练"
      :visible.sync="newTaskDialogVisible"
      :close-on-click-modal="false"
      width="25%"
    >
      <el-form ref="newForm" :model="newForm" label-width="100px">
        <el-form-item label="数据集" size="small">
          <el-cascader
            props.expand-trigger="hover"
            :options="formOptions.dataset"
            v-model="newForm.dataset"
          ></el-cascader>
        </el-form-item>
        <el-form-item label="输出层激活" size="small">
          <el-select v-model="newForm.hyperParameter.outputLayerActivation" placeholder="请选择">
            <el-option
              v-for="item in formOptions.outputLayerActivation"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="损失函数" size="small">
          <el-select v-model="newForm.hyperParameter.loss" placeholder="请选择">
            <el-option
              v-for="item in formOptions.loss"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="随机数种子" size="small">
          <el-input-number :step="1" v-model="newForm.hyperParameter.seed"></el-input-number>
        </el-form-item>
        <el-form-item label="Batch Size" size="small">
          <el-input-number
            :min="10"
            :max="10000"
            :step="10"
            v-model="newForm.hyperParameter.batchSize"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="Epochs" size="small">
          <el-input-number
            :min="10"
            :max="10000"
            :step="10"
            v-model="newForm.hyperParameter.epochs"
          ></el-input-number>
        </el-form-item>
        <el-form-item label="学习率" size="small">
          <el-input-number :min="0" :step="0.001" v-model="newForm.hyperParameter.learningRate"></el-input-number>
        </el-form-item>
        <el-form-item label="隐藏层数" size="small">
          <el-input-number
            :min="1"
            :max="10"
            :step="1"
            v-model="newForm.NnLayers"
            @change="changeLayer"
          ></el-input-number>
        </el-form-item>
        <el-form-item
          v-for="(layer, index) in newForm.hyperParameter.hiddenLayerStructure"
          :key="index"
          size="mini"
          style="margin-bottom:0px"
        >
          <el-row :gutter="10">
            <el-col :span="10">
              <el-input v-model.number="layer.neurons"></el-input>
            </el-col>
            <el-col :span="10">
              <el-select v-model="layer.activation" placeholder="请选择">
                <el-option
                  v-for="item in formOptions.activation"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-col>
          </el-row>
        </el-form-item>
        <!-- <el-form-item label="模型">
          <el-select v-model="newForm.nn" placeholder="请选择">
            <el-option
              v-for="item in formOptions.nn"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>-->
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
import globals from '@/globals'

export default {
  name: 'DeepLearning',
  data() {
    return {
      ws: null, // websocket 对象
      wsTimer: null,
      newTaskDialogVisible: false,
      chartDialogVisible: false,
      formOptions: {
        dataset: [
          {
            value: '北汽',
            label: '北汽',
            children: [
              {
                value: 'LNBSCU3HXJR884327放电',
                label: 'LNBSCU3HXJR884327放电'
              }
            ]
          }
        ],
        // label 长度不要超过 7 个汉字，否则有样式问题
        nn: [
          {
            value: '普通神经网络',
            label: '普通神经网络'
          }
        ],
        outputLayerActivation: [
          {
            value: 'Softmax',
            label: 'Softmax'
          },
          {
            value: 'Sigmoid',
            label: 'Sigmoid'
          },
          {
            value: 'Linear',
            label: 'Linear'
          }
        ],
        loss: [
          {
            value: 'MSE',
            label: 'MSE'
          },
          {
            value: 'L1',
            label: 'L1'
          },
          {
            value: 'Sickle-L1',
            label: 'Sickle-L1'
          }
        ],
        activation: [
          {
            value: 'ReLu',
            label: 'ReLu'
          },
          {
            value: 'Softmax',
            label: 'Softmax'
          },
          {
            value: 'Linear',
            label: 'Linear'
          }
        ]
      },
      formStepActive: 1,
      // 请求的参数
      newForm: {
        dataset: ['北汽', 'LNBSCU3HXJR884327放电'],
        hyperParameter: {
          outputLayerActivation: 'Linear',
          loss: 'MSE',
          seed: 1,
          batchSize: 600,
          epochs: 100,
          learningRate: 0.001,
          hiddenLayerStructure: []
        },
        NnLayers: 1,
        nn: '普通神经网络'
      },
      tableData: [],
      chartOption: {}
    }
  },
  methods: {
    handleNewForm() {
      this.newForm = {
        dataset: ['北汽', 'LNBSCU3HXJR884327放电'],
        hyperParameter: {
          outputLayerActivation: 'Linear',
          loss: 'MSE',
          seed: 1,
          batchSize: 600,
          epochs: 100,
          learningRate: 0.001,
          hiddenLayerStructure: [
            {
              neurons: 64,
              activation: 'ReLu'
            }
          ]
        },
        NnLayers: 1,
        nn: '普通神经网络'
      }
      this.newTaskDialogVisible = true
    },
    changeLayer(count) {
      const len = this.newForm.hyperParameter.hiddenLayerStructure.length
      if (count > len) {
        for (let i = 0, temp = count - len; i < temp; i++) {
          this.newForm.hyperParameter.hiddenLayerStructure.push({
            neurons: 64,
            activation: 'ReLu'
          })
        }
      } else if (count < len) {
        for (let i = 0, temp = len - count; i < temp; i++) {
          this.newForm.hyperParameter.hiddenLayerStructure.pop()
        }
      }
    },
    _checkNewForm(form) {
      return null
    },
    createTask() {
      const ret = this._checkNewForm(this.newForm)
      if (ret !== null) {
        this.$message.error(ret)
        return
      }

      let params = {
        dataset: this.newForm.dataset.join('_'),
        hyperParameter: this.newForm.hyperParameter
      }

      return (
        this.$axios
          .post(globals.URL_API_DL_TASKS, params)
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
    _buildTrainingHistoryPlotOption(data) {
      return {
        title: {
          text: '训练过程',
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
          // name: 'Epochs',
          // nameLocation: 'start',
          type: 'category',
          data: globals.range(1, data['loss'].length + 1, 1)
        },
        yAxis: [
          {
            name: 'Loss',
            min: 0
          },
          {
            name: 'Accuracy',
            min: 0,
            max: 1,
            splitNumber: 10
          }
        ],
        series: [
          {
            name: 'Loss',
            type: 'line',
            data: data['loss'],
            yAxisIndex: 0
          },
          {
            name: 'Accuracy',
            type: 'line',
            data: data['accuracy'],
            yAxisIndex: 1
          }
        ],
        legend: {
          data: ['Loss', 'Accuracy']
        }
      }
    },
    plotTrainingHistory(index, row) {
      const status = row.status
      const url = `${globals.URL_API_DL_TASKS}/${row.taskId}/training-history`
      this.$axios
        .get(url)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }
          this.chartDialogVisible = true
          this.chartOption = this._buildTrainingHistoryPlotOption(jd.data)
        })
        .catch(error => {
          this.$message.error(error.message)
        })
    },
    _buildEvalResultOption(data) {
      return {
        title: {
          text: '模型评估误差分布',
          //subtext: '子标题',
          x: 'center'
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          orient: 'vertical',
          left: 'left',
          data: [
            '误差<1%',
            '误差1%～2%',
            '误差2%～3%',
            '误差3%～4%',
            '误差>=4%'
          ]
        },
        series: [
          {
            name: '访问来源',
            type: 'pie',
            //radius: '55%',
            //center: ['50%', '60%'],
            data: [
              { value: data['a1Count'], name: '误差<1%' },
              { value: data['a2Count'], name: '误差1%～2%' },
              { value: data['a3Count'], name: '误差2%～3%' },
              { value: data['a4Count'], name: '误差3%～4%' },
              { value: data['aOtherCount'], name: '误差>=4%' }
            ],
            itemStyle: {
              emphasis: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      }
    },
    plotEvalResult(index, row) {
      this.$axios
        .get(`${globals.URL_API_DL_TASKS}/${row.taskId}/eval-result`)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }
          this.chartDialogVisible = true
          this.chartOption = this._buildEvalResultOption(jd.data)
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
            .delete(`${globals.URL_API_DL_TASKS}/${row.taskId}`)
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
    }
  },
  created() {
    this.$axios
      .get(globals.URL_API_DL_TASKS)
      .then(response => response.data)
      .then(jd => {
        if (jd.code !== globals.SUCCESS) {
          throw new Error(jd.msg)
        }
        this.tableData = jd.data
      })
      .catch(error => {
        this.$message.error(error.message)
      })

    this.ws = new WebSocket(
      'ws://' + document.location.host + globals.URL_WS_DL_TASKS
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

    this.ws.onclose = () => {
      // 这里代码不能少，如果遇到异常关闭，会走 onclose
      if (this.wsTimer !== null) {
        this.wsTimer = null
      }
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