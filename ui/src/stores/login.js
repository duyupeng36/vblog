
import {useStorage} from "@vueuse/core"

export default useStorage("loginState", {
    username:"",
    isLogin: false,
}, localStorage, {mergeDefaults: true})