<script setup lang="ts">
import { ref, onMounted } from "vue";

const props = defineProps<{ hash: string }>();
const clickCount = ref(0);

onMounted(async () => {
    try {
        const res = await fetch(`/gst/${props.hash}`);
        const data = await res.json();
        clickCount.value = data.click_count;
    } catch (e) {
        console.error(e);
    }
});
</script>

<template>
    <div class="stats-card">
        <h2 class="stats-count">{{ clickCount }} People opened this link</h2>
    </div>
</template>

<style scoped>
.stats-card {
    width: 100%;
    height: 60px;
    background-color: #212121;
    border: 1px solid #474745;
    border-radius: 20px;
    padding: 0 16px;
    display: flex;
    align-items: center;
    flex-shrink: 0;
}

.stats-count {
    color: #e0e0dd;
    font-size: clamp(12px, 2.5vw, 17px);
    font-weight: 600;
    margin: 0;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
</style>
