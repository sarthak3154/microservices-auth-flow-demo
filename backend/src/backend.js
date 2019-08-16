const Mali = require('mali')
const pino = require('pino')

const PORT = '8080'

global.logger = pino({
    name: 'grpc-backend',
    messageKey: 'message',
    changeLevelName: 'severity',
    useLevelLabels: true,
    prettyPrint: true
})

exports.initialize = () => {
    const app = new Mali()
    app.start(`localhost:${PORT}`)
    logger.info(`GRPC Server listening on Port: ${PORT}`)
}