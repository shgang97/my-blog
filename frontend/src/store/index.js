import {createStore} from 'vuex'

export default createStore({
    state() {
        return {
            token: '',
            userInfo: {}
        }
    },
    mutations: {
        SET_TOKEN: (state, token) => {
            state.token = token
            localStorage.setItem('token', token)
        },
        SET_USERINFO: (state, userInfo) => {
            state.userInfo = userInfo
            sessionStorage.setItem('userInfo', JSON.stringify(userInfo))
        },
        REMOVE_INFO: (state) => {
            state.token = ''
            state.userInfo = {}
            localStorage.removeItem("token")
            sessionStorage.removeItem("userInfo")
            // localStorage.setItem('token', '')
            // sessionStorage.setItem('userInfo', JSON.stringify(''))
        }
    },
    getters: {
        getUser: state => state.userInfo
    }
})