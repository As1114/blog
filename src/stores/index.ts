import {defineStore} from 'pinia'
import {userInfo} from "@/api/user";
import router from "@/router";

const theme: boolean = true // true light   false  dark
const collapsed: boolean = false
const templateCollapsed: boolean = false

interface useStoreInfo {
    id: number
    nick_name: string
    avatar: string
    role: number
    token: string
}

const userStoreInfo: useStoreInfo = {
    id: 0,
    nick_name: '',
    avatar: '',
    role: 0,
    token: ''
}

export const useStore = defineStore('counter', {
    state() {
        return {
            theme: theme,
            collapsed: collapsed,
            templateCollapsed: templateCollapsed,
            userStoreInfo: userStoreInfo,
        }
    },
    actions: {
        setCollapsed(collapsed: boolean) {
            this.collapsed = collapsed
        },
        setTemplateCollapsed(templateCollapsed: boolean) {
            this.templateCollapsed = templateCollapsed
        },
        setTheme(localTheme?: boolean) {
            if (localTheme != undefined) {
                this.theme = localTheme
            } else {
                this.theme = !this.theme
            }
            document.documentElement.style.colorScheme = this.themeString
            document.body.setAttribute('arco-theme', this.themeString)
            localStorage.setItem("theme", JSON.stringify(this.theme))
        },
        loadTheme() {
            let val = localStorage.getItem("theme");
            if (val === null) {
                return
            }
            try {
                this.theme = JSON.parse(val);
                document.documentElement.style.colorScheme = this.themeString;
                document.body.setAttribute('arco-theme', this.themeString);
            } catch (e) {
                return
            }
        },
        logout() {
        },
        clearUserInfo() {
            this.userStoreInfo = userStoreInfo;
            localStorage.removeItem("userStoreInfo");
        },
        async setToken(data: any) {
            let res = await userInfo()
            this.userStoreInfo = {
                id: res.data.id,
                nick_name: res.data.nick_name,
                avatar: res.data.avatar,
                role: res.data.role,
                token: data.token,
            };
            localStorage.setItem("userStoreInfo", JSON.stringify(this.userStoreInfo));
        },
        loadToken() {
            const storedInfo = localStorage.getItem("userStoreInfo");
            if (!storedInfo) {
                return;
            }
            try {
                this.userStoreInfo = JSON.parse(storedInfo);
            } catch {
                this.clearUserInfo();
                router.push({name: "web_home"});
                return;
            }
            // let exp = Number(this.userStoreInfo.token_expires_at) * 1000
            // let nowTime = new Date().getTime()
            // if (exp <= nowTime) {
            //     Message.warning("登录已过期")
            //     this.clearUserInfo()
            //     router.push({name: "web_home"});
            //     return;
            // }
        },
    },
    getters: {
        themeString(): "light" | "dark" {
            return this.theme ? "light" : "dark"
        },
        isGeneral(): boolean {
            return this.userStoreInfo.role == 2
        },
        isLogin(): boolean {
            return this.userStoreInfo.role == 1 || this.userStoreInfo.role == 2
        }
    }
})
