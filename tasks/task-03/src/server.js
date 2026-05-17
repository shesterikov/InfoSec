import Fastify from 'fastify'
import { Sequelize, DataTypes } from 'sequelize';

const fastify = Fastify({
    logger: true
})

const sequelize = new Sequelize('postgres://student:HjjnGytfnhg@db:5432/db') // В случае запуска Node локально требуется заменить db хост на localhost

const User = sequelize.define("users", {
    username: DataTypes.STRING,
    password: DataTypes.STRING,
    salt: DataTypes.STRING,
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
    const { username, password } = request.body;
    const passwordHash = hash(request.body.password) // Требуется заменить на корректную реализацию хеширования пароля

    const user = new User({ username, password });

    const result = await user.save();

    return result;
})

fastify.post('/login', async function handler(request, reply) {
    const { username, password } = request.body;

    const passwordHash = hash(password) // Требуется заменить на корректную реализацию хеширования пароля

    const user = await User.findOne({ where: { username, password: passwordHash } })

    return user;
})

try {
    await sequelize.sync();
    await fastify.listen({ port: 3000, host: '0.0.0.0' })
} catch (err) {
    fastify.log.error(err)
    process.exit(1)
}