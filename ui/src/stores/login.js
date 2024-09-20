
import {useStorage} from "@vueuse/core"

export default useStorage("loginState", {
    username:"",
    isLogin: false,
    token: "",
}, localStorage, {mergeDefaults: true})