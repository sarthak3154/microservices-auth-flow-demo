const Mali = require('mali')
const path = require('path')
const pino = require('pino')

const { getMessage, getSecuredMessage } = require('./backendGRPC')

const PORT = '8080'

const MESSAGE_PROTO_PATH = path.resolve(__dirname, '..', 'protos/message.proto')

global.logger = pino({
    name: 'grpc-backend',
    messageKey: 'message',
    changeLevelName: 'severity',
    useLevelLabels: true,
    prettyPrint: true
})

exports.initialize = async () => {
    const app = new Mali()
    app.addService(MESSAGE_PROTO_PATH, ['PublicAccessService', 'SecuredAccessService'])
    app.use({
        'PublicAccessService': { getMessage },
        'SecuredAccessService': { getMessage: getSecuredMessage }
    })
    await app.start(`localhost:${PORT}`)
    logger.info(`GRPC Server listening on Port: ${PORT}`)
}