exports.getMessage = (ctx) => {
    ctx.res = { message: 'public' }
}

exports.getSecuredMessage = (ctx) => {
    ctx.res = { message: 'secured' }
}