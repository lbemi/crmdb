<template>
  <el-space wrap :size="20">
    <el-card id="echart" style="width: 400px; height: 200px" />
    <el-card id="echart1" style="width: 400px; height: 200px" />
    <el-card id="echart2" style="width: 400px; height: 200px" />
  </el-space>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
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
  const charEle = document.getElementById('echart2') as HTMLElement
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
  const charEle = document.getElementById('echart1') as HTMLElement
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
</script>

<style scoped lang="less">
table,
table tr th,
table tr td {
  table-layout: fixed;
  border: 1px solid #dbdbdbe5;
  height: 45px;
  width: 100%;
  font-size: 12px;
  color: rgb(145, 143, 143);
  padding-left: 16px;
}
table {
  table-layout: fixed;
  width: 100%;
  min-height: 25px;
  line-height: 25px;
  text-align: left;
  border-collapse: collapse;
}
.title {
  background-color: #f5f7fa;
  color: #333333;
  font-size: 14px;
}
.el-tabs__content {
  padding: 0;
}
</style>
