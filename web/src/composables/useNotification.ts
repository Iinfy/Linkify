// composables/useNotification.ts
import { ref, type Ref } from 'vue'

interface UseNotificationReturn {
    visible: Ref<boolean>
    title: Ref<string>
    subtitle: Ref<string>
    hasError: Ref<boolean>
    show: (_title: string,_subtitle: string, hasError: boolean) => void
}

const visible: Ref<boolean> = ref(false)
const title: Ref<string> = ref('')
const subtitle: Ref<string> = ref('')
const hasError: Ref<boolean> = ref(false)
let timer: ReturnType<typeof setTimeout> | null = null

export function useNotification(): UseNotificationReturn {
    const show = (_title: string,_subtitle: string, _hasError: boolean, duration: number = 2000): void => {
        title.value = _title
        subtitle.value = _subtitle
        visible.value = true
        hasError.value = _hasError

        if (timer) clearTimeout(timer)
        timer = setTimeout(() => {
            visible.value = false
        }, duration)
    }

    return { visible, title, subtitle, hasError, show }
}