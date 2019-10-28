
// base url
const URL_LOGIN = '/login'
const URL_LOGOUT = '/logout'

// 静态资源 URL，nginx 分发
const URL_MEDIA = '/media'
const URL_AVATAR = URL_MEDIA + '/avatars'

// api url
const URL_API_BASE = '/api/v1'
const URL_API_USERS = URL_API_BASE + '/users'
const URL_API_MINING_BASE = URL_API_BASE + '/mining/base'
const URL_API_MINING_TASKS = URL_API_BASE + '/mining/tasks'
const URL_API_DL_TASKS = URL_API_BASE + '/dl/tasks'

// websocket url
const URL_WS_BASE = '/websocket/v1'
const URL_WS_MINING_TASKS = URL_WS_BASE + '/mining/tasks'
const URL_WS_DL_TASKS = URL_WS_BASE + '/dl/tasks'

// 返回的 json 状态码
const SUCCESS = 20000
const ERROR = 40000

// 用户类型
const USER_TYPE_SUPER_USER = 64
const USER_TYPE_COMMON_USER = 0

// 用户状态
const USER_STATUS_NORMAL = 1
const USER_STATUS_FORBIDDEN_LOGIN = 0

// 输入过滤
export const RE_SIX_CHARACTER_CHECKER = /^[0-9a-zA-Z]{5,14}$/;

// 全局函数
// 若字符串长度小于 len，左边补空格
function leftFillSpace(s, len) {
  const remain = len - s.length
  let temp = ''
  for (let i = 0; i < remain; i++) {
    temp += ' '
  }
  return temp + s
}

export default {
  URL_LOGIN,
  URL_LOGOUT,
  // media
  URL_MEDIA,
  URL_AVATAR,
  // api
  URL_API_BASE,
  URL_API_USERS,
  URL_API_MINING_BASE,
  URL_API_MINING_TASKS,
  URL_API_DL_TASKS,
  // ws
  URL_WS_BASE,
  URL_WS_MINING_TASKS,
  URL_WS_DL_TASKS,

  SUCCESS,
  ERROR,

  USER_TYPE_SUPER_USER,
  USER_TYPE_COMMON_USER,

  USER_STATUS_NORMAL,
  USER_STATUS_FORBIDDEN_LOGIN,

  RE_SIX_CHARACTER_CHECKER,

  leftFillSpace
}