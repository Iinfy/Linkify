// composables/useNotification.ts
import { ref, type Ref } from 'vue'

interface UseNotificationReturn {
    visible: Ref<boolean>
    title: Ref<string>
    subtitle: Ref<string>
    show: (_title: string,_subtitle: string) => void
}

const visible: Ref<boolean> = ref(false)
const title: Ref<string> = ref('')
const subtitle: Ref<string> = ref('')
let timer: ReturnType<typeof setTimeout> | null = null

export function useNotification(): UseNotificationReturn {
    const show = (_title: string,_subtitle: string, duration: number = 2000): void => {
        title.value = _title
        subtitle.value = _subtitle
        visible.value = true

        if (timer) clearTimeout(timer)
        timer = setTimeout(() => {
            visible.value = false
        }, duration)
    }

    return { visible, title, subtitle, show }
}