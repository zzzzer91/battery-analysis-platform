<template>
  <div>
    <div class="container">
      <div class="mgb20">
        <el-button type="primary" icon="el-icon-plus" @click="handleNew">新建用户</el-button>
      </div>
      <el-table :data="tableData" border ref="multipleTable">
        <el-table-column prop="userName" label="用户名"></el-table-column>
        <el-table-column prop="createTime" label="创建时间" sortable></el-table-column>
        <el-table-column prop="lastLoginTime" label="最后登录" sortable></el-table-column>
        <el-table-column prop="loginCount" sortable label="登录次数"></el-table-column>
        <el-table-column prop="userStatus" sortable label="允许登录" :formatter="formatter"></el-table-column>
        <el-table-column prop="comment" label="备注" show-overflow-tooltip></el-table-column>
        <el-table-column label="操作">
          <template v-slot:default="scope">
            <el-button
              type="text"
              icon="el-icon-edit"
              @click="handleEdit(scope.$index, scope.row)"
            >编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页 -->
      <!-- <div class="pagination">
        <el-pagination
          background
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          layout="prev, pager, next"
          :total="tableDataTotal"
        ></el-pagination>
      </div>-->
    </div>

    <!-- 新建弹出框 -->
    <el-dialog title="新建用户" :visible.sync="newVisible" :close-on-click-modal="false" width="30%">
      <el-form ref="form" :model="newForm">
        <el-form-item label="用户名">
          <el-input maxlength="14" v-model="newForm.userName"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input maxlength="14" type="password" v-model="newForm.password"></el-input>
        </el-form-item>
        <el-form-item label="确认密码">
          <el-input maxlength="14" type="password" v-model="newForm.confirmPassword"></el-input>
        </el-form-item>
        <el-form-item label="备注">
          <el-input maxlength="64" v-model="newForm.comment"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer">
        <el-button @click="newVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveNew">确 定</el-button>
      </span>
    </el-dialog>

    <!-- 编辑弹出框 -->
    <el-dialog title="修改" :visible.sync="editVisible" :close-on-click-modal="false" width="30%">
      <el-form ref="form" :model="editForm">
        <el-form-item label="用户名">
          <el-input v-model="editForm.userName" maxlength="14" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="备注">
          <el-input maxlength="64" v-model="editForm.comment"></el-input>
        </el-form-item>
        <el-form-item label="状态">
          <el-radio v-model="editForm.userStatus" :label="1">允许登录</el-radio>
          <el-radio v-model="editForm.userStatus" :label="0">禁止登录</el-radio>
        </el-form-item>
      </el-form>
      <span slot="footer">
        <el-button @click="editVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveEdit">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import moment from 'moment'
import globals from '@/globals'

export default {
  name: 'UserManager',
  data() {
    return {
      currentPage: 1,
      tableDataTotal: 0,
      tableData: [],
      newVisible: false,
      editVisible: false,
      newForm: {
        userName: '',
        password: '',
        confirmPassword: '',
        comment: '',
        loginCount: 0,
        userStatus: globals.USER_STATUS_NORMAL
      },
      editForm: {
        userName: '',
        comment: '',
        userStatus: globals.USER_STATUS_FORBIDDEN_LOGIN
      },
      editFormIdx: -1
    }
  },
  computed: {
    tableDataLength() {
      return this.tableData.length
    }
  },
  methods: {
    // 切换分页
    handleCurrentChange: function(currentPage) {
      this.currentPage = currentPage
    },
    // 格式化用户状态，table 不能直接显示 bool 类型
    formatter(row, column) {
      return row.userStatus === globals.USER_STATUS_NORMAL ? '是' : '否'
    },
    handleNew() {
      this.newForm = {
        userName: '',
        password: '',
        confirmPassword: '',
        comment: '',
        userStatus: globals.USER_STATUS_NORMAL,
        loginCount: 0,
        createTime: null
      }
      this.newVisible = true
    },
    // 保存新建
    saveNew() {
      const err_msg = this.checkForm(this.newForm)
      if (err_msg === null) {
        this.$axios
          .post(globals.URL_API_USERS, this.newForm)
          .then(response => response.data)
          .then(jd => {
            if (jd.code !== globals.SUCCESS) {
              throw new Error(jd.msg)
            }
            this.newForm.createTime = moment().format('YYYY-MM-DD HH:mm:ss')
            this.tableData.push(this.newForm)
            this.newVisible = false
            this.$message.success(jd.msg)
          })
          .catch(error => {
            this.$message.error(error.message)
          })
      } else {
        this.$message.error(err_msg)
      }
    },
    handleEdit(index, row) {
      this.editFormIdx = index
      this.editForm = {
        userName: row.userName,
        comment: row.comment,
        userStatus: row.userStatus
      }
      this.editVisible = true
    },
    // 保存编辑
    saveEdit() {
      this.$axios
        .put(
          `${globals.URL_API_USERS}/${this.editForm.userName}`,
          this.editForm
        )
        .then(response => response.data)
        .then(jd => {
          if (jd.code !== globals.SUCCESS) {
            throw new Error(jd.msg)
          }

          this.tableData[this.editFormIdx].comment = this.editForm.comment
          this.tableData[this.editFormIdx].userStatus = this.editForm.userStatus
          this.editVisible = false
          this.$message.success(jd.msg)
        })
        .catch(error => {
          this.$message.error(error.message)
        })
    },
    handleDelete(index, row) {
      this.$confirm('此操作将永久删除该用户, 是否继续?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      })
        .then(() => {
          this.tableData.splice(index, 1)
          this.$message({
            type: 'success',
            message: `删除用户 ${row.userName} 成功!`
          })
        })
        .catch(() => {})
    },
    checkForm(form) {
      let err_msg = null
      if (form.userName === '') {
        err_msg = '用户名不能为空！'
      } else if (form.password === '' || form.confirmPassword === '') {
        err_msg = '密码不能为空！'
      } else if (!globals.RE_SIX_CHARACTER_CHECKER.test(form.userName)) {
        err_msg = '用户名须 5～14 位，并且只能是字母和数字！'
      } else if (!globals.RE_SIX_CHARACTER_CHECKER.test(form.password)) {
        err_msg = '密码须 5～14 位，并且只能是字母和数字！'
      } else if (form.password !== form.confirmPassword) {
        err_msg = '两次密码不一致！'
      } else if (form.comment.length > 64) {
        err_msg = '备注过长！'
      }
      return err_msg
    }
  },
  beforeCreate() {
    this.$axios
      .get(globals.URL_API_USERS)
      .then(response => response.data)
      .then(jd => {
        if (jd.code !== globals.SUCCESS) {
          throw new Error(jd.msg)
        }

        this.tableData = jd.data
      })
      .catch(() => {})
  }
}
</script>

<style scoped>
.pagination {
  margin: 20px 0;
  text-align: right;
}
.red {
  color: #ff0000;
}
</style>
