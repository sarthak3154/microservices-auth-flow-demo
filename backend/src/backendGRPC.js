exports.getMessage = (ctx) => {
    ctx.res = { message: 'public' }
}

exports.getSecuredMessage = (ctx) => {
    ctx.res = { message: 'secured' }
}

exports.createResource = (ctx) => {
    logger.info(ctx.req)
    ctx.res = { success: true }
}