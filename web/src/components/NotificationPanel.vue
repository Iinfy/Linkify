<script setup lang="ts">
import { useNotification } from '@/composables/useNotification'
const { visible, title,subtitle, hasError } = useNotification()

</script>

<template>
  <transition name="slide">
  <div v-if="visible" class="notif-container" :class="{ 'border-red': hasError, 'border-white': !hasError }">
    <img v-if="hasError" src="@/assets/icons/warning.svg" class="check" alt="">
    <img v-else src="@/assets/icons/check.svg" class="check" alt="">
    <div class="text-group">
      <h2>{{title}}</h2>
      <p>{{subtitle}}</p>
    </div>
  </div>
  </transition>
</template>

<style scoped>

* {
  margin: 0;
  padding: 0;

}

.notif-container {
  display: flex;
  align-items: center;
  position: fixed;
  bottom: 0;
  right: 0;
  padding: 10px 15px;
  margin-right: 15px;
  margin-bottom: 15px;
  background: #212121;
  height: 70px;
  width: auto;
  min-width: 280px;
  max-width: 480px;
  color: white;
  border-radius: 20px;
}

.border-red  { border: 1px solid red; }
.border-white { border: 1px solid white; }

.text-group {
  display: flex;
  flex-direction: column;
  margin-left: 45px;
  font-size: clamp(11px, 2vw, 14px);
}

.notif-container h2 {
  font-size: 15px;
  font-weight: 500;
  line-height: 1.2;
}

.notif-container p {
  margin-top: 4px;
  font-size: 13px;
  color: #5F5F5F;
  line-height: 1.2;
}

.check {
  position: absolute;
  left: 20px;
  top: 50%;
  transform: translateY(-50%);
  width: 22px;
  height: 22px;
  object-fit: contain;
  filter: invert(1);
}

.slide-enter-active {
  transition: all 0.4s ease;
}

.slide-leave-active {
  transition: all 0.4s ease;
}

.slide-enter-from {
  transform: translateX(100%);
}

.slide-leave-to {
  transform: translateX(100%);
}

@media (min-width: 769px) {
  .notif-container {
    bottom: 0;
    right: 0;
    margin: 0 20px 20px 0;
    width: auto;
    min-width: 320px;
    max-width: 420px;
  }
}


@media (min-width: 481px) and (max-width: 768px) {
  .slide-enter-from {
    transform: translate(-50%, 100%);
    opacity: 0;
  }

  .slide-leave-to {
    transform: translate(-50%, 100%);
    opacity: 0;
  }





  .notif-container {
    bottom: 0;
    left: 50%;
    right: auto;
    transform: translateX(-50%);
    margin: 0 0 20px 0;
    width: calc(100% - 40px);
    min-width: auto;
    max-width: 500px;
  }
}


@media (max-width: 480px) {
  .notif-container {
    top: 0;
    left: 0;
    right: 0;
    bottom: auto;
    margin: 0;
    width: 100%;
    border-radius: 0 0 20px 20px;
    max-width: 100%;
    animation: slideDown 0.4s ease;
  }
}

</style>