/** * Created by lei on 2022/11/08 */
<template>
  <svg :class="svgClass" aria-hidden="true">
    <use :xlink:href="iconClassName" :fill="color"></use>
  </svg>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  iconName: {
    type: String,
    required: true
  },
  className: {
    type: String
    // required: true
  },
  color: {
    type: String,
    default: '#73767a'
  }
})

const iconClassName = computed(() => {
  return `#${props.iconName}`
})

const svgClass = computed(() => {
  if (props.className) {
    return `${props.className}`
  }
  return 'svg-icon'
})
</script>

<style scoped lang="less">
.svg-icon {
  position: relative;
  width: 1em;
  height: 1em;
  // fill: currentColor;
}

.icon-1-4em {
  position: relative;
  left: 30%;
  width: 1.4em;
  height: 1.4em;
}

.icon-2-4em {
  position: relative;
  width: 2.4em;
  height: 2.4em;
}
</style>
