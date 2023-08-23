import request from '@/utils/request'

// 查询ModelCompileApply列表
export function listModelCompileApply(query) {
    return request({
        url: '/api/v1/model-compile-apply',
        method: 'get',
        params: query
    })
}

// 查询ModelCompileApply详细
export function getModelCompileApply (id) {
    return request({
        url: '/api/v1/model-compile-apply/' + id,
        method: 'get'
    })
}

export function getUserRole() {
    return request({
        url: '/api/v1/model-compile-apply/userrole',
        method: 'get'
    })
}

// 新增ModelCompileApply
export function addModelCompileApply(data) {
    return request({
        url: '/api/v1/model-compile-apply',
        method: 'post',
        data: data
    })
}

// 修改ModelCompileApply
export function updateModelCompileApply(data) {
    return request({
        url: '/api/v1/model-compile-apply/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除ModelCompileApply
export function delModelCompileApply(data) {
    return request({
        url: '/api/v1/model-compile-apply',
        method: 'delete',
        data: data
    })
}

export function uploadFile(formData) {
    return request({
        url: '/api/myv1_c/upload_single',
        method: 'post',
        data: formData
    })
}

export function uploadFile_out(formData) {
    return request({
        url: '/api/myv1_c/upload_single_definename',
        method: 'post',
        data: formData
    })
}

