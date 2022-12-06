<template>
  <div class="warp">
    <div ref="bar" class="bar"></div>
  </div>
</template>

<script lang="ts" setup>
import { ref, nextTick } from 'vue'
import '@/styles/loading.scss'

let bar = ref<HTMLElement>()
let speed = ref<number>(1)
let timer = ref<number>(0)

const startLoading = () => {
  speed.value = 1
  let dom = bar.value as HTMLElement
  timer.value = window.requestAnimationFrame(function fn() {
    if (speed.value < 100) {
      speed.value += 1
      dom.style.width = speed.value + '%'
      timer.value = window.requestAnimationFrame(fn)
    } else {
      speed.value = 1
      window.cancelAnimationFrame(timer.value)
    }
  })
}
const endLoading = () => {
  let dom = bar.value as HTMLElement
  setTimeout(() => {
    window.requestAnimationFrame(() => {
      speed.value = 100
      dom.style.width = speed.value + '%'
    })
  }, 0)
}

const show = () => {
  const bodys: Element = document.body
  const div = document.createElement('div')
  div.className = 'block-loading'
  div.innerHTML = `
            <div class="block-loading-box">
                <div class="block-loading-box-warp">
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                    <div class="block-loading-box-item"></div>
                </div>
            </div>
        `
  bodys.insertBefore(div, bodys.childNodes[0])
}
const hide = () => {
  nextTick(() => {
    setTimeout(() => {
      const el = document.querySelector('.block-loading')
      el && el.parentNode?.removeChild(el)
    }, 1000)
  })
}

defineExpose({
  startLoading,
  endLoading
})
</script>

<style scoped lang="less">
.warp {
  position: fixed;
  top: 0;
  width: 100%;
  height: 2px;

  .bar {
    width: 0;
    height: inherit;
    background: rgb(48 101 248);
  }
}
</style>
