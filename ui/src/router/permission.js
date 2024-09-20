
import loginState from "@/stores/login.js"

// 定义导航守卫
export async function beforeEachHandler(to, from, next) {
    // 只判断到后台页面
    if (to.fullPath.startsWith("/backend")) {
        // 如果未登陆 重定向到登录页面, 并且把目标页面作为重定向参数传递下去
        if (!loginState.value.isLogin) {
            console.log("not login");
            next({
                path: "/login",
                query: {
                    redirect: to.name,
                    ...to.query,
                },
            });
        } else {
            // 已经登录的用户直接放行
            next();
        }
    } else {
        // 不属于/backend的页面 直接放开
        next();
    }
}
