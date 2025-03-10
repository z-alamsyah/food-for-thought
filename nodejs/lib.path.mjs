import path from 'path';

const filePath = '/Users/UMB/NodeJS/Standard-LIbrary.html';
console.info(path.sep);
console.info(path.dirname(filePath));
console.info(path.basename(filePath));
console.info(path.extname(filePath));
console.info(path.parse(filePath));