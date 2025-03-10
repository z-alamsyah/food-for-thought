import { readFile, writeFile } from 'fs/promises';

async function extractFile() {
    try {
        const data = await readFile('data.txt', 'utf8');
        console.log('Isi file:', data);
    } catch (err) {
        console.error('Error membaca file:', err);
    }
}

async function writeToFile() {
    try {
        await writeFile('output.txt', 'Hello, Node.js!');
        console.log('File berhasil ditulis.');
    } catch (err) {
        console.error('Error menulis file:', err);
    }
}

extractFile(); // call read File Function
writeToFile(); // call write File Function