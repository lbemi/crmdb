/** * Created by lei on 2022/11/08 */
<template>
  <div>
    <svg :class="svgClass" aria-hidden="true">
      <use :xlink:href="iconClassName" :fill="color"></use>
    </svg>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive } from 'vue';


const props = defineProps({
  iconName: {
    type: String,
    required: true
  },
  className: {
    type: String,
    // required: true
  },
  color: {
    type: String,
    default: '#73767a'
  }
})

const iconClassName = computed(()=>{
  return `#${props.iconName}`
})

const svgClass  = computed(()=>{
  if (props.className) {
    return `svg-icon ${props.className}`
  }
  return 'svg-icon'
})

</script>

<style scoped lang="less">
.svg-icon {
  width: 1em;
  height: 1em;
  position: relative;
  // fill: currentColor;
}
</style>
