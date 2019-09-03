const AWS = require('aws-sdk')
const { aws, cognito } = require('./config/config')

AWS.config.update({ ...aws })

const authConfig = {
    AuthFlow: 'ADMIN_NO_SRP_AUTH',
    ClientId: cognito.clientId,
    UserPoolId: cognito.userPoolId
}

exports.handler = (event, context, callback) => {
    context.callbackWaitsForEmptyEventLoop = false
    if (!event.hasOwnProperty('username') || !event.hasOwnProperty('password')) {
        return callback(new Error('Invalid parameters'))
    }

    const authParams = {
        ...authConfig, AuthParameters: {
            USERNAME: event.username,
            PASSWORD: event.password
        }
    }

    const cognito = new AWS.CognitoIdentityServiceProvider()

    cognito.adminInitiateAuth(authParams, (err, data) => {
        if (err) return callback(err)
        callback(null, data)
    })
}