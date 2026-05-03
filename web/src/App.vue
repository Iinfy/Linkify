<template>
  <div class="page-wrapper">
    <div class="main__section">
      <AppLogo />
      <LinkInput />
      <RecentLinks />
    </div>
  </div>
</template>

<script setup lang="ts">
import AppLogo from './components/AppLogo.vue'
import LinkInput from './components/LinkInput.vue'
import RecentLinks from './components/RecentLinks.vue'

</script>

<style>
@import './assets/main.css';
</style>