import Fastify from 'fastify'
import * as jwt from '@fastify/jwt'
import { Sequelize, DataTypes } from 'sequelize';

const fastify = Fastify({
    logger: true
})

fastify.register(jwt, {
    secret: 'qwerty'
})

fastify.decorate("authenticate", async function (request, reply) {
    try {
        await request.jwtVerify()
    } catch (err) {
        reply.send(err)
    }
})

const sequelize = new Sequelize('postgres://student:HjjnGytfnhg@db:5432/db') // В случае запуска Node локально требуется заменить db на localhost

const User = sequelize.define("users", {
    id: {
        type: DataTypes.INTEGER,
        autoIncrement: true,
        primaryKey: true,
    },
    username: DataTypes.STRING,
    password: DataTypes.STRING,
    role: {
        defaultValue: 'user',
        type: DataTypes.ENUM("user", "admin")
    }
});

/**
 * Функция - образец
 * @param {string} str 
 * @returns {string}
 */
function hash(str) {
    return '46827DF0837878';
}

fastify.get('/health', async function handler(request, reply) {
    return { status: 'ok' }
})

fastify.post('/register', async function handler(request, reply) {
    const passwordHash = hash(request.body.password)
    const newUser = new User({ username: request.body.username, password: request.body.password });

    const result = await newUser.save();

    return { message: 'success', data: result }
})

fastify.post('/login', async function handler(request, reply) {
    const user = await User.findByPk(1);

    const token = fastify.jwt.sign({
        id: user.id,
        role: user.role
    })

    return { token }
})

fastify.get('/all', async function handler(request, reply) {
    return { message: 'visible_all' }
})

fastify.get("/users", async function handler(request, reply) {
    return { message: 'super_secret_for_users' }
})

fastify.get('/admins', async function handler(request, reply) {
    return { message: 'super_secret_for_admins' }
})

try {
    await sequelize.sync();
    await fastify.listen({ port: 3000, host: '0.0.0.0' })
} catch (err) {
    fastify.log.error(err)
    process.exit(1)
}