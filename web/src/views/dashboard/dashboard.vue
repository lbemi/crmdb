/** * Created by lei on 2022/09/25 */
<template>
  <el-space wrap :size="40">
    <el-card id="echart" style="width: 400px; height: 200px" />
    <el-card id="echart1" style="width: 400px; height: 200px" />
    <el-card id="echart2" style="width: 400px; height: 200px" />
  </el-space>
  <el-space wrap>
    <div ref="barChart" :style="{ width: '400px', height: '300px' }"></div>
    <div ref="pieChart" :style="{ width: '500px', height: '300px' }"></div>
  </el-space>
</template>
<script setup lang="ts">
import * as echarts from 'echarts'
import { onMounted, ref } from 'vue'
import { ECharts, EChartOption, init } from 'echarts'

onMounted(() => {
  app1()
  app2()
  app3()
})

const app1 = () => {
  const charEle = document.getElementById('echart') as HTMLDivElement
  const charEch: ECharts = init(charEle)
  const option: EChartOption = {
    xAxis: {
      type: 'category',
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: [150, 230, 224, 218, 135, 147, 260],
        type: 'line'
      }
    ]
  }
  charEch.setOption(option)
}
const app2 = () => {
  const charEle = document.getElementById('echart2') as HTMLDivElement
  const charEch: ECharts = init(charEle)
  const option: EChartOption = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      top: '5%',
      left: 'center'
    },
    series: [
      {
        name: 'Access From',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 40,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: [
          { value: 1048, name: 'Search Engine' },
          { value: 735, name: 'Direct' },
          { value: 580, name: 'Email' },
          { value: 484, name: 'Union Ads' },
          { value: 300, name: 'Video Ads' }
        ]
      }
    ]
  }
  charEch.setOption(option)
}
const app3 = () => {
  const charEle = document.getElementById('echart1') as HTMLDivElement
  const charEch: ECharts = init(charEle)
  const option: EChartOption = {
    xAxis: {
      type: 'category',
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: [150, 230, 224, 218, 135, 147, 260],
        type: 'line'
      }
    ]
  }
  charEch.setOption(option)
}
const barChart = ref<HTMLDivElement>()
const myChart1 = ref<any>()
// 绘制柱形图
function initBarEcharts() {
  myChart1.value = echarts.init(barChart.value!)
  myChart1.value.setOption({
    title: {
      text: '测试',
      x: 'center'
    },
    tooltip: {
      trigger: 'item'
    },
    legend: {
      data: ['输出量', '输入量'],
      orient: 'vertical',
      right: 60,
      top: 20
    },
    xAxis: {
      data: ['六月', '七月', '八月', '九月', '十月']
    },
    yAxis: {},
    color: ['#c38bef', '#6dbcf7'],
    series: [
      {
        name: '输出量',
        type: 'line',
        data: [2, 9, 6, 3, 1],
        smooth: true,
        label: {
          show: true,
          position: 'top'
        }
      },
      {
        name: '输入量',
        type: 'bar',
        data: [9, 18, 12, 9, 6],
        barWidth: '20',
        label: {
          show: true,
          position: 'top'
        }
      }
    ]
  })
}
//绘制饼图
const pieChart = ref<HTMLDivElement>()
const myChart2 = ref<any>()
function initPieEcharts() {
  myChart2.value = echarts.init(pieChart.value!)
  myChart2.value.setOption({
    title: {
      text: '主机分类',
      x: 'center'
    },
    tooltip: {
      trigger: 'item',
      formatter: '{b}:{c}({d}%)'
    },
    legend: {
      type: 'scroll',
      orient: 'vertical',
      left: 0,
      top: 30,
      height: 150,
      data: [
        'html',
        'css',
        'javascript',
        'typescript',
        'vue2',
        'vue3',
        'react',
        'angular',
        'uniapp',
        'taro',
        'vite',
        'webpack',
        'node',
        'others'
      ]
    },
    color: [
      '#61adf2',
      '#56dae8',
      '#efa49e',
      '#7ea1ed',
      '#5ae05a',
      '#f2d2a2',
      '#5881e8',
      '#63d6c0',
      '#edc29f',
      '#5b97d3',
      '#3eceb3',
      '#6a96ed',
      '#426ed1',
      '#65d18b'
    ],
    series: [
      {
        name: '文章分类数量',
        type: 'pie',
        radius: '30%',
        center: ['60%', '50%'],
        data: [
          { name: 'html', value: 10 },
          { name: 'css', value: 12 },
          { name: 'javascript', value: 20 },
          { name: 'typescript', value: 15 },
          { name: 'vue2', value: 13 },
          { name: 'vue3', value: 12 },
          { name: 'react', value: 3 },
          { name: 'angular', value: 2 },
          { name: 'uniapp', value: 12 },
          { name: 'taro', value: 5 },
          { name: 'vite', value: 2 },
          { name: 'webpack', value: 3 },
          { name: 'node', value: 9 },
          { name: 'others', value: 10 }
        ],
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0,0,0,0.5)'
          }
        },
        label: {
          show: true,
          position: 'outside',
          formatter: '{b}:{c}({d}%)'
        }
      }
    ]
  })
}
onMounted(() => {
  initBarEcharts()
  initPieEcharts()
})
</script>

<style scoped lang="less"></style>
