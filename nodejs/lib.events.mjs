import { EventEmitter } from 'events';

const emitter = new EventEmitter();
emitter.on('hello', (name) => {
    console.info(`Hello ${name}`);
});

emitter.addListener('hello', (name) => {
    console.info(`Hai ${name}`);
})

emitter.emit('hello', 'Andre!');
