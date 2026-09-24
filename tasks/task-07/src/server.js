import Fastify from 'fastify';
import { fileURLToPath } from 'url';
import { dirname, join } from 'path';
import { Server } from 'socket.io';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const fastify = Fastify({
    logger: true
});

fastify.register(import('@fastify/static'), {
    root: join(__dirname, '../public'),
    prefix: '/'
});

fastify.get('/health', async function handler(request, reply) {
    return { status: 'ok' }
})

const server = fastify.server;

const io = new Server(server);

io.on('connection', (socket) => {
    console.log('New client connected');

    socket.on('new message', (msg) => {
        console.log('Message:', msg);
        io.emit('new message', msg);
    });

    socket.on('disconnect', () => {
        console.log('Client disconnected');
    });
});

try {
    await fastify.listen({ port: 3000, host: '0.0.0.0' })
} catch (err) {
    fastify.log.error(err)
    process.exit(1)
}