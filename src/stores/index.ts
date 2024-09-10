import {defineStore} from 'pinia'

const theme: boolean = true // true light   false  dark
const collapsed: boolean = false

export const useStore = defineStore('counter', {
    state() {
        return {
            theme: theme,
            collapsed: collapsed,
        }
    },
    actions: {
        setCollapsed(collapsed: boolean) {
            this.collapsed = collapsed
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
        }
    },
    getters: {
        themeString(): "light" | "dark" {
            return this.theme ? "light" : "dark"
        },
    }
})
