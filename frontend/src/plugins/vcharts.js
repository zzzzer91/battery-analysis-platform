import Vue from 'vue'
import ECharts from 'vue-echarts'
// import 'echarts'

// 只导入需要组件，减少文件体积
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/toolbox'
import 'echarts/lib/component/title'
import 'echarts/lib/component/dataZoom'
import 'echarts/lib/component/legendScroll'
import 'echarts/lib/component/visualMap'
import 'echarts/lib/chart/line'
import 'echarts/lib/chart/bar'
import 'echarts/lib/chart/heatmap'
import 'echarts/lib/chart/pie'
// import 'echarts/lib/chart/scatter'

// import 'echarts/theme/roma.js'  // 主题

// 3D
import 'echarts-gl'

Vue.component('v-chart', ECharts)