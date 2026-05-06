import { ref, type Ref } from 'vue'

interface ShowQRCodeReturn {
    visible: Ref<boolean>
    text: Ref<string>
    showQR: (_text: string) => void
    close: () => void
}

const visible: Ref<boolean> = ref(false)
const text: Ref<string> = ref('not_found')


export function showQRCode(): ShowQRCodeReturn{
    const showQR = (_text: string) =>{
        text.value = _text
        visible.value = true
    }

    const close = () => {
        visible.value = false
    }

    return {visible,text, showQR ,close}
}