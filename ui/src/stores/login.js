
import {useStorage} from "@vueuse/core"


export const loginState =  useStorage("loginState", {
    username:"",
    isLogin: false,
}, localStorage, {mergeDefaults: true})

