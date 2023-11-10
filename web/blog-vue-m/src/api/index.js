import request from '@/utils/request'

export function fetchList(data, token) {
    return request({
        url: '/article/list',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchUserInfo(token) {
    return request({
        url: '/user/info',
        method: 'get',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
    })
}

export function fetchIsLogin(token) {
    return request({
        url: '/auth/is_login',
        method: 'get',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
    })
}

export function fetchLogOut(token) {
    return request({
        url: '/auth/logout',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
    })
}

export function fetchUpdatePwd(data, token) {
    return request({
        url: '/user/update_password',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToUpdateMotto(data, token) {
    return request({
        url: '/user/update_motto',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToUpdateSocialAccountQQ(data, token) {
    return request({
        url: '/user/update_account_qq',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToUpdateSocialAccountGithub(data, token) {
    return request({
        url: '/user/update_account_github',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToUpdateSocialAccountGitee(data, token) {
    return request({
        url: '/user/update_account_gitee',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToUpdateSocialAccountCSDN(data, token) {
    return request({
        url: '/user/update_account_csdn',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToUpdateEmail(data, token) {
    return request({
        url: '/user/update_email',
        method: 'post',
        headers: {
            'Authorization': `${token}` // 将 Token 添加到 Authorization 头部
        },
        data: data
    })
}

export function fetchToGetCateGory() {
    return request({
        url: '/article/category',
        method: 'get',
    })
}

export function fetchUploadArticleImage(file, token) {
    const formData = new FormData();
    formData.append('file', file); // 将文件添加到 FormData 中

    return request({
        url: '/upload/article_image',
        method: 'post',
        headers: {
            'Authorization': token // 将 Token 添加到 Authorization 头部
        },
        data: formData, // 使用 FormData 作为请求的数据
    });
}

export function fetchUploadArticleCover(file, token) {
    const formData = new FormData();
    formData.append('file', file); // 将文件添加到 FormData 中

    return request({
        url: '/upload/article_cover',
        method: 'post',
        headers: {
            'Authorization': token // 将 Token 添加到 Authorization 头部
        },
        data: formData, // 使用 FormData 作为请求的数据
    });
}

export function fetchToCreateArticle(data, token) {
    return request({
        url: '/article/write',
        method: 'post',
        headers: {
            'Authorization': token // 将 Token 添加到 Authorization 头部
        },
        data: data,
    });
}

export function fetchToArticleUpvote(data, token) {
    return request({
        url: '/article/upvote',
        method: 'post',
        headers: {
            'Authorization': token // 将 Token 添加到 Authorization 头部
        },
        data: data,
    });
}

export function fetchArticle(id) {
    return request({
        url: `/article/detail/${id}`,
        method: 'get',
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

// 重置密码
export function fetchResetPwd(email,data) {
    return request({
        url: `/auth/reset_pwd/${email}`,
        method: 'post',
        data: data
    })
}

// 重置密码验证码
export function fetchResetPwdCode() {
    return request({
        url: '/auth/reset_pwd/code',
        method: 'get',
    })
}


