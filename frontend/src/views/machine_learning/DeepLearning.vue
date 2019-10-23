<template>
  <div>
    <div class="container">
      <div class="mgb20">
        <el-button type="primary" icon="el-icon-plus" @click="handleNewForm">新建训练</el-button>
      </div>
      <el-table :data="tableData" border height="500" ref="multipleTable">
        <el-table-column prop="dataset" label="数据集" width="340" sortable></el-table-column>
        <el-table-column prop="nnName" label="模型" width="180" sortable show-overflow-tooltip></el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180"></el-table-column>
        <el-table-column prop="taskStatus" label="状态" width="100" sortable>
          <template scope="scope">
            <el-tag v-if="scope.row.taskStatus === '完成'" type="success">{{scope.row.taskStatus}}</el-tag>
            <el-tag
              v-else-if="scope.row.taskStatus === '执行中'"
              type="warning"
            >{{scope.row.taskStatus}}</el-tag>
            <el-tag v-else-if="scope.row.taskStatus === '失败'" type="danger">{{scope.row.taskStatus}}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="comment" label="备注"></el-table-column>
        <el-table-column label="操作" width="160">
          <template v-slot:default="scope">
            <el-button
              type="text"
              icon="el-icon-pie-chart"
              @click="doPlot(scope.$index, scope.row)"
              v-show="scope.row.taskStatus==='完成'"
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
      title="新建训练"
      :visible.sync="newTaskDialogVisible"
      :close-on-click-modal="false"
      width="35%"
    >
      <el-form ref="form" :model="newForm">
        <el-form-item label="数据集">
          <el-select v-model="newForm.dataset" placeholder="请选择">
            <el-option
              v-for="item in datasetOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="模型">
          <el-select v-model="newForm.nnName" placeholder="请选择">
            <el-option
              v-for="item in nnOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
          </el-select>
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
import globals from '@/globals'

export default {
  name: 'DeepLearning',
  data() {
    return {
      ws: null, // websocket 对象
      wsTimer: null,
      newTaskDialogVisible: false,
      chartDialogVisible: false,
      datasetOptions: [
        {
          value: 'Minst',
          label: 'Minst'
        },
        {
          value: '北汽LNBSCU3HXJR884327',
          label: '北汽LNBSCU3HXJR884327'
        },
      ],
      // label 长度不要超过 7 个汉字，否则有样式问题
      nnOptions: [
        {
          value: 'BP神经网络',
          label: 'BP神经网络'
        }
      ],
      // 请求的参数
      newForm: {
        dataset: 'Minst',
        nnName: 'BP神经网络',
      },
      tableData: [],
      chartOption: {}
    }
  },
  methods: {
    handleNewForm() {
      this.newForm = {
        dataset: 'Minst',
        nnName: 'BP神经网络'
      }
      this.newTaskDialogVisible = true
    },
    createTask() {
      let params = {
        dataset: this.newForm.dataset,
        nnName: this.newForm.nnName,
      }

      return (
        this.$axios
          .post(globals.URL_API_ML_TASKS, params)
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
    doPlot(index, row) {
      this.$axios
        .get(`${globals.URL_API_ML_TASKS}/${row.taskId}`)
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }
          this.chartDialogVisible = true
          this.chartOption = null
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
            .delete(`${globals.URL_API_ML_TASKS}/${row.taskId}`)
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
    this.ws = new WebSocket(
      'ws://' + document.location.host + globals.URL_WS_ML_TASKS
    )

    // 连接成功后的回调
    this.ws.onopen = () => {
      this.ws.send('ping')
      this.wsTimer = window.setInterval(() => this.ws.send('ping'), 3000)
    }

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
        window.clearInterval(this.wsTimer)
        this.wsTimer = null
      }
    }
  },
  beforeDestroy() {
    // 如果没有这一步判断，当 ws 关闭时，定时器可能还在运作，控制台会报个错
    if (this.wsTimer !== null) {
      window.clearInterval(this.wsTimer)
      this.wsTimer = null
    }
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