<script setup lang="ts">
import AppLogo from "@/components/AppLogo.vue";

</script>

<template>
  <AppLogo/>
  <div class="not_found_div">
    <h1 class="not_found_text">Page not found</h1>
  </div>
</template>

<style scoped>
@import "@/assets/main.css";
.not_found_div{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 50vh;
}
.not_found_text{
  color: white;
  font-size: clamp(35px, 2.5vw, 70px);
}
</style>