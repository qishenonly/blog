import request from '@/utils/request'

export function fetchList(params) {
    return request({
        url: '/post/list',
        method: 'get',
        params: params
    })
}

export function fetchFocus() {
    return request({
        url: '/focus/list',
        method: 'get',
        params: {}
    })
}

export function fetchCategory() {
    return request({
        url: '/category',
        method: 'get',
        params: {}
    })
}

export function fetchFriend() {
    return request({
        url: '/friend',
        method: 'get',
        params: {}
    })
}

export function fetchSocial() {
    return request({
        url: '/social',
        method: 'get',
        params: {}
    });
}

export function fetchSiteInfo() {
    return request({
        url: '/site',
        method: 'get',
        params: {}
    })
}

export function fetchComment() {
    return request({
        url: '/comment',
        method: 'get',
        params: {}
    })
}

// 注册
export function fetchRegister(data) {
    return request({
        url: '/auth/register',
        method: 'post',
        data: data
    })
}

// 注册验证码
export function fetchRegisterCode() {
    return request({
        url: '/auth/register/code',
        method: 'get',
    })
}

// 登录
export function fetchLogin(data) {
    return request({
        url: '/auth/login',
        method: 'post',
        data: data
    })
}

// 登录验证码
export function fetchLoginCode() {
    return request({
        url: '/auth/login/code',
        method: 'get',
    })
}

